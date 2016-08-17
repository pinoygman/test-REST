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


#sed -i -e 's/{cfport}/'${PORT}'/g' ./raas-outbound/config/smtp.ini
haraka -c raas-outbound &
#./node_modules/Haraka/bin/haraka -c raas-outbound &
sleep 10;
./${ARTIFACT}_linux
#./pcs-v1_Darwin
