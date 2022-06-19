# Step CA installation

## Get the docker image
```sh
sudo docker pull smallstep/step-ca
```

## Configure CA online
`https://smallstep.com/app/jyotish-mundali/login` is the online step CA.
> `jyotish-mundali` is the name of the authority that I created.<br>
> Password hint: `old ambassador`

## Bootstrap step CA
```sh
$ export STEP_CA_BASE=`pwd` # skip if same as pwd
$ export STEP_CA_URL=<step-ca-url>
$ export STEP_CA_FINGER_PRINT=<step-ca-finger-print>
$ make
```

## Clean up step CA
```sh
$ export STEP_CA_BASE=`pwd` # skip if same as pwd
$ make clean
```

## Generate certificate
```sh
$ ./gen-cert.sh <hostname>
```
The default certificate generated expire in 30 days.