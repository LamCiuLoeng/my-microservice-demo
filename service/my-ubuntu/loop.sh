#!/usr/bin/env bash
trap "curl http://${DOCKER_HOST_IP}:2379/v2/keys/test -XDELETE ; exit 0" SIGTERM
#trap "curl http://localhost:2379/v2/keys/test -XDELETE ; exit 0" SIGTERM

echo $DOCKER_HOST_IP
curl http://${DOCKER_HOST_IP}:2379/v2/keys/test -XPUT -d value="start@$(date)"

while true;
	do :;
	done
