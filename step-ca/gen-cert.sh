#!/bin/sh

# usage: ./gen-cert.sh <host-name>

BASEDIR=${STEP_CA_BASE-${PWD}}

sudo docker run -it --rm --network host --volume ${BASEDIR}/data:/home/step smallstep/step-ca \
	step ca certificate $1 $1.crt $1.key --not-after 720h
sudo docker run -it --rm --network host --volume ${BASEDIR}/data:/home/step smallstep/step-ca \
	step certificate inspect --short $1.crt
