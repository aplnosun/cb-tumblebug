#!/bin/bash
source ../setup.env

#for NAME in "${CONNECT_NAMES[@]}"
#do
	curl -H "${AUTH}" -sX GET http://$TUMBLEBUG_IP:1323/tumblebug/ns/${NS_ID}/resources/publicIp | json_pp &
#done

