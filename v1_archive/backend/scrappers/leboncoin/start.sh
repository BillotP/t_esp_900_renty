#!/usr/bin/env bash

# Get env secrets
source backend/openfaas/.dev.env || source ../../openfaas/.dev.env
# Run scrappy locally
STAGE=dev \
 DEBUG=true \
 arango_host=$DB_HOST \
 arango_port=$DB_PORT \
 arango_scheme=$DB_SCHEME \
 arango_user=$DB_USER \
 arango_password=$DB_PASSWORD \
 arango_dbname=$DB_NAME \
 arango_tlsverify=$DB_TLSVERIFY \
 go run .
