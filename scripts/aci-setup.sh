#!/usr/bin/env bash

cd "$(dirname "$0")"

docker run --network="app" -p 27017:27017 --name some-mongo -d mongo:4.0-xenial
docker run --network="app" -d -p 5672:5672 --hostname my-rabbit --name some-rabbit rabbitmq:3

cd ../send/
docker build -t tripdubroot/send:aci .

cd ../receive/
docker build -t tripdubroot/receive:aci .

cd ../show/
docker build -t tripdubroot/show:aci .

cd ..
docker push tripdubroot/send:aci
docker push tripdubroot/receive:aci
docker push tripdubroot/show:aci
