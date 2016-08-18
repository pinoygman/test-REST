#!/bin/bash
#
#  Copyright (c) 2016 General Electric Company. All rights reserved.
#
#  The copyright to the computer software herein is the property of
#  General Electric Company. The software may be used and/or copied only
#  with the written permission of General Electric Company or in accordance
#  with the terms and conditions stipulated in the agreement/contract
#  under which the software has been supplied.
#
#  author: chia.chang@ge.com
#

#set -x
set -e

function cf_push () {

    export https_proxy=${PROXY}
    
    op=$(cf login -a ${CF_END} -u ${CF_USER} -p ${CF_PWD} -o ${CF_ORG} -s ${CF_SPC})

    echo Checking CF Envs..

    cf a &> /dev/null
    if [ $? -eq 0 ] ; then   
	echo Good you have logged in the CF;
	cp -r ./assets ./${DIST}
	cp -r ./email/* ./${DIST}
	cd ./${DIST}
	cf push
	# pcs-backend-${ENV} -c "./${ARTIFACT}_linux" -b https://github.com/cloudfoundry/binary-buildpack.git --no-start
	# cf start pcs-backend-${ENV}
    else 
	echo Please log in the Cloud Foundry org/space.; 
	exit 1; 
    fi;

}

function docker_run () {
    docker pull golang
    docker build -t ${ARTIFACT}_img .
    docker run -i --name ${ARTIFACT}_inst ${ARTIFACT}_img
    
    CID=$(docker ps -aqf "name=${ARTIFACT}_inst")
    IID=$(docker images -q ${ARTIFACT}_img)
    
    docker cp ${CID}:/${DIST} ./${DIST}

    docker rm ${CID}

    docker rmi ${IID}
}

function readinputs () {

    if [ $# -eq 0 ]
    then
	printf "   %-*s\n" 10 "-p    | proxy"
	printf "   %-*s\n" 10 "-r    | revision"
	printf "   %-*s\n" 10 "-b    | build version by Jenkins"
	
	printf "   %-*s\n" 10 "-e    | environemntal variables"
	printf "   %-*s\n" 10 "-user | cf username"
	printf "   %-*s\n" 10 "-pwd  | cf password"
	printf "   %-*s\n" 10 "-end  | cf endpoint"
	printf "   %-*s\n" 10 "-org  | cf org"
	printf "   %-*s\n" 10 "-spc  | cf space"
	printf "   %-*s\n" 10 "-sql  | postgresSql constr"
	printf "   %-*s\n" 10 "-husr | haraka user"
	printf "   %-*s\n" 10 "-hpwd | haraka pwd"
	printf "   %-*s\n" 10 "-hhst | haraka host"
	printf "   %-*s\n" 10 "-hsht | haraka smtp host"
    else
	for ((i = 1; i <=$#; i++));
	do
	    case ${@:i:1} in
		-p)
		    PROXY=${@:i+1:1}
		    ;;
		-r)
		    REV=${@:i+1:1}
		    ;;
		
		-e)
		    ENV=${@:i+1:1}
		    ;;
		-b)
		    BUILD_VER=${@:i+1:1}
		    ;;
		-user)
		    CF_USER=${@:i+1:1}
		    ;;
		-pwd)
		    CF_PWD=${@:i+1:1}		   
		    ;;
		-end)
		    CF_END=${@:i+1:1}		   
		    ;;
		-org)
		    CF_ORG=${@:i+1:1}		   
		    ;;
		-spc)
		    CF_SPC=${@:i+1:1}		   
		    ;;
		-sql)
		    SQLDSN=${@:i+1:1}
		    ;;
		-husr)
		    HUSER=${@:i+1:1}
		    ;;
		-hpwd)
		    HPWD=${@:i+1:1}
		    ;;
		-hhst)
		    HHOST=${@:i+1:1}
		    ;;
		-hsht)
		    HSHOST=${@:i+1:1}
		    ;;
		*)
		    #echo "Invalid option ${@:i:1}"
	            ;;
	    esac
	done
    fi

}

readinputs $@

DIST=dist
ARTIFACT=pcs_backend_${REV}_${ENV}
BUILD_TIME=`date +%FT%T%z`
LDFLAGS="main.REV=${REV}"
DHOME=github.build.ge.com/predixsolutions/catalog-onboarding-backend

#predix select
HOST=run.asv-pr.ice.predix.io

eval "sed -i -e 's#{DHOME}#${DHOME}#g' ./Dockerfile"
eval "sed -i -e 's#{ARTIFACT}#${ARTIFACT}#g' ./Dockerfile"
eval "sed -i -e 's#{REV}#${REV}#g' ./Dockerfile"
eval "sed -i -e 's#{PROXY}#${PROXY}#g' ./Dockerfile"
eval "sed -i -e 's#{BUILD_TIME}#${BUILD_TIME}#g' ./Dockerfile"
eval "sed -i -e 's#{BUILD_VER}#${BUILD_VER}#g' ./Dockerfile"
eval "sed -i -e 's#{LDFLAGS}#${LDFLAGS}#g' ./Dockerfile"
eval "sed -i -e 's#{DIST}#${DIST}#g' ./Dockerfile"
eval "sed -i -e 's#{SQLDSN}#${SQLDSN}#g' ./Dockerfile"

eval "sed -i -e 's#{HOST}#pcs-backend-${ENV}.${HOST}#g' ./assets/swagger.json"
eval "sed -i -e 's#{BASE}#/${REV}/api#g' ./assets/swagger.json"
eval "sed -i -e 's#{ENV}#${ENV}#g' ./assets/swagger.json"

eval "sed -i -e 's#{ENV}#${ENV}#g' ./email/manifest.yml"
eval "sed -i -e 's#{SQLDSN}#${SQLDSN}#g' ./email/manifest.yml"
eval "sed -i -e 's#{ARTIFACT}#${ARTIFACT}#g' ./email/manifest.yml"
eval "sed -i -e 's#{HUSER}#${HUSER}#g' ./email/manifest.yml"
eval "sed -i -e 's#{HPWD}#${HPWD}#g' ./email/manifest.yml"
eval "sed -i -e 's#{HHOST}#${HHOST}#g' ./email/manifest.yml"
eval "sed -i -e 's#{HSHOST}#${HSHOST}#g' ./email/manifest.yml"

docker_run
cf_push
