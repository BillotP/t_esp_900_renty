#!/usr/bin/env bash
get_latest_release() {
  curl -s "https://api.github.com/repos/$1/releases/latest" |       # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}
echo -e "[INFO] Getting latest release version from kubernetes/dashboard github"
tag=$(get_latest_release kubernetes/dashboard)
echo -e "[INFO] Get tag $tag"
echo -e "[INFO] Apply yaml crd to cluster"
k3s kubectl create -f https://raw.githubusercontent.com/kubernetes/dashboard/${tag}/aio/deploy/recommended.yaml
echo -e "[INFO] Creating admin user"
echo "
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
" | k3s kubectl apply -f -
echo -e "[INFO] Creating admin role"
echo "
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
" | k3s kubectl apply -f -
echo -e "[INFO] Getting Authorization Token (that you should copy)"
k3s kubectl -n kubernetes-dashboard describe secret admin-user-token | grep ^token
echo -e "[INFO] Starting kubectl proxy in background"
k3s kubectl proxy &
echo -e "[INFO] Done, now head to http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/ to access dashboard (with the previous token on clipboard)"
