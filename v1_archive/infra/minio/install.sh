#!/usr/bin/env bash
set -e
get_mc_cli() {
  curl -Ls "https://dl.min.io/client/mc/release/linux-amd64/mc" --output mc
  chmod +x mc
}
HOST=$(ifconfig wlp2s0 | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p')
HOSTURL=$(echo -n $HOST | sed 's/\./-/g')
DEFAULTBUCKETNAME=renty-assets-dev
DEFAULTBUCKETPOLICY=public
echo -e "[INFO] Adding minio helm repository"
helm repo add minio https://helm.min.io/
echo -e "[INFO] Updating helm repos"
helm repo update
echo -e "[INFO] Export local kubeconfig envvar"
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
[ -f .credential ] && source .credential || echo -e "[INFO] Generating strong random access and secret key pair"
[ -f .credential ] || export ACCESSKEY=$(head -c 12 /dev/urandom | shasum | cut -d' ' -f1)
[ -f .credential ] || export SECRETKEY=$(head -c 12 /dev/urandom | shasum | cut -d' ' -f1)
echo -e "[INFO] Installing minio to kube with previously generated secrets and default renty-assets-dev bucket"
helm upgrade minio --install minio/minio \
	--set accessKey=$ACCESSKEY,secretKey=$SECRETKEY \
	--set buckets[0].name=$DEFAULTBUCKETNAME,buckets[0].policy=$DEFAULTBUCKETPOLICY,buckets[0].purge=false
echo -e "[INFO] Saving previously generated credential in local .credential file"
echo \
"ACCESSKEY=$ACCESSKEY
SECRETKEY=$SECRETKEY
" > .credential
echo -e "[INFO] Adding an ingress rule through gloo gateway on http://bucket.$HOSTURL.sslip.io/minio"
k3s kubectl apply -f minio-proxy.yaml
echo -e "[INFO] Downloading mc cli if missing"
[ -f ./mc ] || get_mc_cli
echo -e "[INFO] Waiting setup and setting default local alias"
sleep 45
./mc alias set local "http://bucket.$HOSTURL.sslip.io" $ACCESSKEY $SECRETKEY
echo -e "\n\t Done , enjoy your MinIO buckets 🪣 !"
