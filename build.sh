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
	
	cd ./${DIST}
	cf push pcs-backend-${ENV} -c "./${ARTIFACT}_linux" -b https://github.com/cloudfoundry/binary-buildpack.git      
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
	printf "   %-*s\n" 10 "-p | proxy"
	printf "   %-*s\n" 10 "-r | revision"
	printf "   %-*s\n" 10 "-b | build version by Jenkins"
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
		*)
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

eval "sed -i -e 's#{DHOME}#${DHOME}#g' ./Dockerfile"
eval "sed -i -e 's#{ARTIFACT}#${ARTIFACT}#g' ./Dockerfile"
eval "sed -i -e 's#{REV}#${REV}#g' ./Dockerfile"
eval "sed -i -e 's#{PROXY}#${PROXY}#g' ./Dockerfile"
eval "sed -i -e 's#{BUILD_TIME}#${BUILD_TIME}#g' ./Dockerfile"
eval "sed -i -e 's#{BUILD_VER}#${BUILD_VER}#g' ./Dockerfile"
eval "sed -i -e 's#{LDFLAGS}#${LDFLAGS}#g' ./Dockerfile"
eval "sed -i -e 's#{DIST}#${DIST}#g' ./Dockerfile"

docker_run $@
cf_push
