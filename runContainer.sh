#!/bin/bash

help() {
    echo "Options: "
    echo " -i: set the docker image name. Default: smartcontract"
    echo " -c: set the codebase path to use with the container." 
    echo "         Default: '$PWD'" 
    echo " -b: force build the container"
    echo "Usage: $0 < -i dockerimage >|< -c codepath >|< -b >"
    exit;
}

# get arguments
# https://stackoverflow.com/questions/22383336/option-requires-an-argument-error
while getopts i:c:d:p:n:h:b option
do
case "${option}"
in
i) DOCKERIMAGE=${OPTARG};;
c) CODEPATH=${OPTARG};;
b) FORCEBUILD=${OPTARG};;
h) help;;
esac
done

# set defaults
if [[ ! ${DOCKERIMAGE} ]]; then DOCKERIMAGE="smartcontract"; fi;
if [[ ! ${CODEPATH} ]]; then CODEPATH="$PWD"; fi;



# check if the container exists if not build it.
if ! grep --quiet $DOCKERIMAGE <<< "$(docker images)"; then docker build --build-arg UID=$(id -u)  . -t $DOCKERIMAGE -- < Dockerfile 
fi;

if [[ ${FORCEBUILD} ]]; then docker build --build-arg UID=$(id -u)  . -t $DOCKERIMAGE -- < Dockerfile ; echo "Build finished." ; exit; fi; 


MOUNT_CODEPATH="-v ${CODEPATH}:/code"
MOUNT_SSHPATH="-v ${HOME}/.ssh:/home/dporto/.ssh"

echo "Usage options: $0 -h"
docker run --name smartcontractdev --rm  $MOUNT_SSHPATH $MOUNT_CODEPATH -it ${DOCKERIMAGE}
