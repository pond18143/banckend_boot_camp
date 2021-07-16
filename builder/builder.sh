#!/bin/bash

env=$1
id=PID
echo "checkout ..."
./checkout.sh
echo "$env Building..."
GOOS=linux
GOARCH=amd64
cd .. && git checkout $env && git pull origin $env && go build .
echo "Copy file to server..."
cd builder\ malar

#scp -r -i big.pem ../api-wecode-supplychain ubuntu@ec2-54-174-95-219.compute-1.amazonaws.com:/home/ubuntu/

echo "Finish"

