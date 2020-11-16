#!/usr/bin/env bash
set -e
source ../../backend/openfaas/.dev.env || source backend/openfaas/.dev.env

encoded=$(echo -n $SMTP_HOST | base64)
k3s kubectl create secret generic smtp-host \
  --from-literal smtp-host=$SMTP_HOST \
  -n openfaas-fn || kubectl patch secret smtp-host \
  -n openfaas-fn -p="{\"data\":{\"smtp-host\": \"$encoded\"}}" -v=1

encoded=$(echo -n $SMTP_PORT | base64)
k3s kubectl create secret generic smtp-port \
  --from-literal smtp-port=$SMTP_PORT \
  -n openfaas-fn || kubectl patch secret smtp-port \
  -n openfaas-fn -p="{\"data\":{\"smtp-port\": \"$encoded\"}}" -v=1

encoded=$(echo -n $SMTP_USER | base64)
k3s kubectl create secret generic smtp-user \
  --from-literal smtp-user=$SMTP_USER \
  -n openfaas-fn || kubectl patch secret smtp-user \
  -n openfaas-fn -p="{\"data\":{\"smtp-user\": \"$encoded\"}}" -v=1

encoded=$(echo -n $SMTP_PASSWORD | base64)
k3s kubectl create secret generic smtp-password \
  --from-literal smtp-password=$SMTP_PASSWORD \
  -n openfaas-fn || kubectl patch secret smtp-password \
  -n openfaas-fn -p="{\"data\":{\"smtp-password\": \"$encoded\"}}" -v=1

encoded=$(echo -n $SMTP_DOMAIN | base64)
k3s kubectl create secret generic smtp-domain \
  --from-literal smtp-domain=$SMTP_DOMAIN \
  -n openfaas-fn || kubectl patch secret smtp-domain \
  -n openfaas-fn -p="{\"data\":{\"smtp-domain\": \"$encoded\"}}" -v=1
