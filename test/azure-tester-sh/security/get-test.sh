#!/bin/bash
source ../setup.env

#for NAME in "${CONNECT_NAMES[@]}"
#do
#	ID=security01-powerkim
#	curl -sX GET http://$RESTSERVER:1024/securitygroup/${ID}?connection_name=${NAME} |json_pp &
#done

TB_SECURITYGROUP_IDS=`curl -sX GET http://$TUMBLEBUG_IP:1323/ns/$NS_ID/resources/securityGroup | json_pp |grep "\"id\"" |awk '{print $3}' |sed 's/"//g' |sed 's/,//g'`
#echo $TB_SECURITYGROUP_IDS | json_pp

if [ "$TB_SECURITYGROUP_IDS" != "" ]
then
        TB_SECURITYGROUP_IDS=`curl -sX GET http://$TUMBLEBUG_IP:1323/ns/$NS_ID/resources/securityGroup | json_pp |grep "\"id\"" |awk '{print $3}' |sed 's/"//g' |sed 's/,//g'`
        for TB_SECURITYGROUP_ID in ${TB_SECURITYGROUP_IDS}
        do
                echo ....Get ${TB_SECURITYGROUP_ID} ...
                curl -sX GET http://$TUMBLEBUG_IP:1323/ns/$NS_ID/resources/securityGroup/${TB_SECURITYGROUP_ID} | json_pp
        done
else
        echo ....no securityGroups found
fi