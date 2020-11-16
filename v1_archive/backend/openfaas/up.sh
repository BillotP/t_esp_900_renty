#!/usr/bin/env bash
# exit when any command fails
set -e
source ../../.credential && source .dev.env
args=""
if [ $# -eq 1 ]; then args="--filter $1" ; fi
faas-cli up \
    --build-arg CI_USER=$GITHUB_USER \
    --build-arg CI_TOKEN=$GITHUB_GOGET_TOKEN \
    --build-arg arango_host=$DB_HOST \
    --build-arg arango_port=$DB_PORT \
    --build-arg arango_scheme=$DB_SCHEME \
    --build-arg arango_user=$DB_USER \
    --build-arg arango_password=$DB_PASSWORD \
    --build-arg arango_dbname=$DB_NAME \
    --build-arg arango_tlsverify=$DB_TLSVERIFY \
    $args
