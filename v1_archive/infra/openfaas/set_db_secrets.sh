#!/usr/bin/env bash
source ../../backend/openfaas/.dev.env || source backend/openfaas/.dev.env

encoded=$(echo -n $DB_HOST | base64)
k3s kubectl create secret generic arango-host \
  --from-literal arango-host=$DB_HOST \
  -n openfaas-fn || kubectl patch secret arango-host \
  -n openfaas-fn -p="{\"data\":{\"arango-host\": \"$encoded\"}}" -v=1

encoded=$(echo -n $DB_PORT | base64)
k3s kubectl create secret generic arango-port \
  --from-literal arango-port=$DB_PORT \
  --namespace openfaas-fn || kubectl patch secret arango-port \
  -n openfaas-fn -p="{\"data\":{\"arango-port\": \"$encoded\"}}" -v=1

encoded=$(echo -n $DB_SCHEME | base64)
k3s kubectl create secret generic arango-scheme \
  --from-literal arango-scheme=$DB_SCHEME \
  --namespace openfaas-fn || kubectl patch secret arango-scheme \
  -n openfaas-fn -p="{\"data\":{\"arango-scheme\": \"$encoded\"}}" -v=1

encoded=$(echo -n $DB_USER | base64)
k3s kubectl create secret generic arango-user \
  --from-literal arango-user=$DB_USER \
  --namespace openfaas-fn || kubectl patch secret arango-user \
  -n openfaas-fn -p="{\"data\":{\"arango-user\": \"$encoded\"}}" -v=1

encoded=$(echo -n $DB_PASSWORD | base64)
k3s kubectl create secret generic arango-password \
  --from-literal arango-password=$DB_PASSWORD \
  --namespace openfaas-fn || kubectl patch secret arango-password \
  -n openfaas-fn -p="{\"data\":{\"arango-password\": \"$encoded\"}}" -v=1

encoded=$(echo -n $DB_NAME | base64)
k3s kubectl create secret generic arango-dbname \
  --from-literal arango-dbname=$DB_NAME \
  --namespace openfaas-fn || kubectl patch secret arango-dbname \
  -n openfaas-fn -p="{\"data\":{\"arango-dbname\": \"$encoded\"}}" -v=1

encoded=$(echo -n $DB_ASSETNAME | base64)
k3s kubectl create secret generic arango-assetdbname \
  --from-literal arango-assetdbname=$DB_ASSETNAME \
  --namespace openfaas-fn || kubectl patch secret arango-assetdbname \
  -n openfaas-fn -p="{\"data\":{\"arango-assetdbname\": \"$encoded\"}}" -v=1

encoded=$(echo -n $DB_TLSVERIFY | base64)
k3s kubectl create secret generic arango-tlsverify \
  --from-literal arango-tlsverify=$DB_TLSVERIFY \
  --namespace openfaas-fn ||  kubectl patch secret arango-tlsverify \
  -n openfaas-fn -p="{\"data\":{\"arango-tlsverify\": \"$encoded\"}}" -v=1
