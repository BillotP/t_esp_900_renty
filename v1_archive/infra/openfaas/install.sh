#!/usr/bin/env bash
set -e
function get_faas_cli() {
    echo -e "[INFO] Download and install faas-cli"
    curl -sSL https://cli.openfaas.com | sudo sh
}
function create_dot_env() {
 source infra/minio/.credential
 echo \
"DB_HOST=$HOST
DB_PORT=8529
DB_SCHEME=https
DB_USER=michel
DB_PASSWORD=D8bFJaufwGdUCSC
DB_NAME=renty-dev
DB_TLSVERIFY=false
JWT_ISSUER=https://api.$HOSTURL.sslip.io/
JWT_AUDIENCE=myapi
JWT_KEYID=web-key
JWT_EXPIRACY=15m
SMTP_HOST=[ask_for_it]
SMTP_PORT=587
SMTP_USER=[ask_for_it]
SMTP_PASSWORD=[ask_for_it]
SMTP_DOMAIN=therentyapp.com
BUCKET_ACCESSKEY=$ACCESSKEY
BUCKET_SECRETKEY=$SECRETKEY
" > backend/openfaas/.dev.env
}
HOST=$(ifconfig wlp2s0 | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p')
HOSTURL=$(echo -n $HOST | sed 's/\./-/g')
INGRESS="https://openfaas.$HOSTURL.sslip.io"
echo -e "[INFO] Creating openfaas required namespaces"
k3s kubectl apply -f https://raw.githubusercontent.com/openfaas/faas-netes/master/namespaces.yml
echo -e "[INFO] Using helm 3 to get openfaas charts"
helm repo add openfaas https://openfaas.github.io/faas-netes
echo -e "[INFO] Updating helm repos"
helm repo update
echo -e "[INFO] Setup kubeconfig envvar"
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
echo -e "[INFO] Generating strong random password"
export PASSWORD=$(head -c 12 /dev/urandom | shasum| cut -d' ' -f1)
echo -e "[INFO] Install OpenFAAS to local cluster\n"
helm upgrade openfaas --install openfaas/openfaas \
    --namespace openfaas  \
    --set functionNamespace=openfaas-fn \
    --set generateBasicAuth=false \
    --set operator.create=true \
    --set rgateway.readTimeout=4h \
    --set gateway.writeTimeout=4h \
    --set gateway.upstreamTimeout=3h50m
echo -e "[INFO] Setup previously generated password as admin BASIC Auth credential"
k3s kubectl -n openfaas delete secret basic-auth || true
k3s kubectl -n openfaas create secret generic basic-auth \
	--from-literal=basic-auth-user=admin \
	--from-literal=basic-auth-password="$PASSWORD"
echo -e "[INFO] Storing password in infra/openfaas/.credential file"
echo \
"OPENFAAS_URL=http://127.0.0.1:31112
OPENFAAS_USER=admin
OPENFAAS_PASSWORD=$PASSWORD" > infra/openfaas/.credential
###
if ! command -v faas &> /dev/null
then
    get_faas_cli
fi
echo "[INFO] Will start to watch for sucessfull deployments, look for READY 1/1 on all lines, exit when ready (might take a long time)" && sleep 3
watch -n 3 'k3s kubectl -n openfaas get deployments -l "release=openfaas, app=openfaas"'
echo -e "[INFO] Login with faas-cli"
export OPENFAAS_URL=http://127.0.0.1:31112 && faas login --password $PASSWORD
echo "[INFO] Setting up a proxy with gloo on $INGRESS host"
sed -i  "s/openfaas.*.sslip.io/openfaas.$HOSTURL.sslip.io/g" infra/openfaas/gateway-proxy.yaml
k3s kubectl apply -f infra/openfaas/gateway-proxy.yaml
echo "[INFO] Creating required secrets in .dev.env file if missing"
[ -f backend/openfaas/.dev.env ] || create_dot_env
##
echo -e "[INFO] Creating development RSA signing key for JWT authentication service if missing"
[ -f backend/openfaas/rsa-2048bit-key-pair.pem ] || openssl genrsa -out backend/openfaas/rsa-2048bit-key-pair.pem
echo -e "[INFO] Extracting public key if missing"
[ -f backend/openfaas/rsa-2048bit-key-pair.pub ] || openssl rsa -in backend/openfaas/rsa-2048bit-key-pair.pem -pubout > backend/openfaas/rsa-2048bit-key-pair.pub
##
echo -e "[INFO] Setting up registry secrets"
./infra/openfaas/set_registry_secrets.sh
echo "[INFO] Creating secrets in kube openfaas-fn namespace (to be updated with the real values when needed)"
./infra/openfaas/set_db_secrets.sh
./infra/openfaas/set_jwt_secrets.sh
./infra/openfaas/set_smtp_secrets.sh
echo -e "\n\t Done , enjoy your OpenFAAS Gateway 🛎️\n(deploy function with up.sh script)"
