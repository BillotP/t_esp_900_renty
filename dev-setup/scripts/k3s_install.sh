#!/usr/bin/env bash
set -e
function create_credential_file() {
    echo -e "What's your gitlab username ?"
    read gitlabuser
    echo -e "What's your github username ?"
    read githubuser
    echo -e "What's your treescale username ?\n(head to https://dash.treescale.com/sign-up to create one)"
    read treescaleuser
    echo -e "Enter your gitlab personal access token with read repository rights\n(if missing go to https://gitlab.com/profile/personal_access_tokens)"
    read gogetkey
    echo -e "Enter your gitlab personal access token with read and write container registry rights\n(if missing go to https://gitlab.com/profile/personal_access_tokens)"
    read gitlabtoken
    echo -e "Enter your github personal access token with read repository rights\n(if missing go to https://github.com/settings/tokens)"
    read githubgogettoken
    echo -e "Enter your github personal access token with read and write container registry rights\n(if missing go to https://github.com/settings/tokens)"
    read githubregistrytoken
    echo -e "Enter your treescale access token\n(available at https://dash.treescale.com/u/$treescaleuser/auth-tokens)\n"
    read treescaletoken
    echo -e "And finally enter your treescale repo name (create one on treescale page if not allready done)"
    read treescalereponame
    echo -e "[INFO] Creating .credential file on this repo's root"
    echo "GOGET_USER=$gitlabuser
GOGET_TOKEN=$gogetkey
GITLAB_USER=$gitlabuser
GITLAB_TOKEN=$gitlabtoken
GITLAB_REGISTRY_URL=registry.gitlab.com/ddng/draftlabs/kubebeber
GITHUB_USER=$githubuser
GITHUB_REGISTRY_TOKEN=$githubregistrytoken
GITHUB_GOGET_TOKEN=$githubgogettoken
GITHUB_REGISTRY=docker.pkg.github.com/billop/renty
TREESCALE_USER=$treescaleuser
TREESCALE_TOKEN=$treescaletoken
TREESCALE_REGISTRY_URL=repo.treescale.com/$treescaleuser/$treescalereponame
" > .credential
}
echo -e "[INFO] Installing k3s binary from get.k3s.io"
curl -sfL https://get.k3s.io | K3S_KUBECONFIG_MODE="644" INSTALL_K3S_EXEC="--disable=traefik" sh -
echo -e "[INFO] Disabling starting k3s at boot"
sudo systemctl disable k3s
echo -e "[INFO] Export KUBECONFIG var" && export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
echo -e "[INFO] Wait 30s for cluster setup" && sleep 30
echo -e "[INFO] Checking if new permissions are correct"
k3s kubectl get nodes
echo -e "[INFO] Creating .credential for container registry and go private package auth in this repo's root"
[ -f .credential ] || create_credential_file
source .credential
echo -e "[INFO] Setting Gitlab registry credential secret"
k3s kubectl create secret docker-registry gitlab-registry \
    --docker-username=$GITLAB_USER \
    --docker-password=$GITLAB_TOKEN \
    --docker-email=lopeurd@gmail.com \
    --docker-server=registry.gitlab.com || true
echo -e "[INFO] Setting TreeScale registry credential secret"
k3s kubectl create secret docker-registry treescale-registry \
    --docker-username=$TREESCALE_USER \
    --docker-password=$TREESCALE_TOKEN \
    --docker-email=registry@therentyapp.com \
    --docker-server=repo.treescale.com || true
echo -e "[INFO] Login in to gitlab registry to be able to push image"
echo $GITLAB_TOKEN | docker login registry.gitlab.com --username $GITLAB_USER --password-stdin
echo -e "[INFO] Login in to treescale registry"
echo $TREESCALE_TOKEN | docker login repo.treescale.com --username $TREESCALE_USER --password-stdin
echo -e "[INFO] Login in to treescale registry"
echo $GITHUB_REGISTRY_TOKEN | docker login docker.pkg.github.com --username $GITHUB_USER --password-stdin
echo -e "[INFO] Done, you could now use your images in kube default namespace"
echo -e "\n\t All Done, enjoy your kube! ⛵\n"
