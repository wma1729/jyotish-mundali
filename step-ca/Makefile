STEP_CA_BASE ?= $(shell pwd)

all: bootstrap

.PHONY: bootstrap clean

datadir:
	mkdir -p $(STEP_CA_BASE)/data

bootstrap: datadir
	sudo docker run -it --rm --network host --volume $(STEP_CA_BASE)/data:/home/step smallstep/step-ca \
		step ca bootstrap --ca-url $(STEP_CA_URL) --fingerprint $(STEP_CA_FINGER_PRINT)

clean:
	/bin/rm -rf $(STEP_CA_BASE)/data
