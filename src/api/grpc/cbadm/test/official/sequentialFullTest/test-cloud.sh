#!/bin/bash

function dozing()
{
	duration=$1
	printf "Dozing for %s : " $duration
	for (( i=1; i<=$duration; i++ ))
	do
		printf "%s " $i
		sleep 1
	done
	echo "(Back to work)"
}

#function test_cloud() {
    FILE=../conf.env
    if [ ! -f "$FILE" ]; then
        echo "$FILE does not exist."
        exit
    fi

	FILE=../credentials.conf
    if [ ! -f "$FILE" ]; then
        echo "$FILE does not exist."
        exit
    fi

	source ../conf.env
	AUTH="Authorization: Basic $(echo -n $ApiUsername:$ApiPassword | base64)"
	source ../credentials.conf

	echo "####################################################################"
	echo "## Create MCIS from Zero Base"
	echo "####################################################################"

	CSP=${1}
	REGION=${2:-1}
	POSTFIX=${3:-developer}
	if [ "${CSP}" == "aws" ]; then
		echo "[Test for AWS]"
		INDEX=1
	elif [ "${CSP}" == "azure" ]; then
		echo "[Test for Azure]"
		INDEX=2
	elif [ "${CSP}" == "gcp" ]; then
		echo "[Test for GCP]"
		INDEX=3
	elif [ "${CSP}" == "alibaba" ]; then
		echo "[Test for Alibaba]"
		INDEX=4
	else
		echo "[No acceptable argument was provided (aws, azure, gcp, alibaba, ...). Default: Test for AWS]"
		CSP="aws"
		INDEX=1
	fi

	../1.configureSpider/register-cloud.sh $CSP $REGION $POSTFIX


	_self="${0##*/}"

	echo ""
	echo "[Logging to notify latest command history]"
	echo "[CMD] ${_self} ${CSP} ${REGION} ${POSTFIX}" >> ./executionStatus
	echo ""
	echo "[Executed Command List]"
	cat  ./executionStatus
	echo ""
#}

#test_cloud