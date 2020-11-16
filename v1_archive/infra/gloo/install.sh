#!/usr/bin/env bash
set -e
get_latest_release() {
  curl -s "https://api.github.com/repos/$1/releases/latest" |       # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}

get_glooctl() {
  latest=$(get_latest_release solo-io/gloo)
  curl -Ls "https://github.com/solo-io/gloo/releases/download/$latest/glooctl-linux-amd64" --output glooctl
  chmod +x glooctl
}

echo -e "[INFO] Adding gloo repo to helm"
helm repo add gloo https://storage.googleapis.com/solo-public-helm
echo -e "[INFO] Update repo"
helm repo update
echo -e "[INFO] Creating dedicated gloo-system namespace"
k3s kubectl create namespace gloo-system || true
echo -e "[INFO] Installing gloo helm chart"
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
helm upgrade gloo --install gloo/gloo --namespace gloo-system
echo -e "[INFO] Downloading glooctl"
[ -f glooctl ] || get_glooctl
echo -e "[INFO] Watching gloo pods"
k3s kubectl get all -n gloo-system
echo -e "[INFO] Waiting 30s for setup and check if everythings all right with glooctl"
sleep 30
./glooctl get upstreams
