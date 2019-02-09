#!/usr/bin/env bash
set -e

rundb(){
  aws dynamodb --endpoint-url http://localhost:8000 $@
}

if [[ "$1" == "start" ]]; then
  docker run --detach -p 127.0.0.1:8000:8000 --name dynamodb amazon/dynamodb-local

elif [[ "$1" == "scan" ]]; then
  rundb scan --table-name keva

else
  rundb list-tables
  rundb create-table \
        --table-name keva \
        --attribute-definitions AttributeName=key,AttributeType=S \
        --key-schema AttributeName=key,KeyType=HASH \
        --provisioned-throughput ReadCapacityUnits=2,WriteCapacityUnits=2
fi
