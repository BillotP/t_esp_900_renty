#!/usr/bin/env bash
get_latest_release() {
  curl -s "https://api.github.com/repos/$1/releases/latest" |       # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}
echo -e "[INFO] Getting latest release tag from arangodb/kube-arangodb repo"
tag=$(get_latest_release arangodb/kube-arangodb)
echo -e "[INFO] Found tag $tag"
echo -e "[INFO] Apply arango-crd to local cluster"
kubectl apply -f https://raw.githubusercontent.com/arangodb/kube-arangodb/$tag/manifests/arango-crd.yaml
echo -e "[INFO] Apply arango-deployment"
kubectl apply -f https://raw.githubusercontent.com/arangodb/kube-arangodb/$tag/manifests/arango-deployment.yaml
echo -e "[INFO] Apply arango-storage service"
kubectl apply -f https://raw.githubusercontent.com/arangodb/kube-arangodb/$tag/manifests/arango-storage.yaml
echo -e "[INFO] Apply single server ArangoDeployment"
echo "
apiVersion: database.arangodb.com/v1alpha
kind: ArangoDeployment
metadata:
  name: single-server
spec:
  mode: Single
" | k3s kubectl apply -f -
###
###
echo -e "\n Done, enjoy ArangoDB 🥑 !\n"
