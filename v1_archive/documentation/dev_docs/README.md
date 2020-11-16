# Renty 🏠 - A cloud native usefull service
> Because finding a decent appartment for rent is a pain

## 🔋Requirements

- Better run on a **linux** environment (see top 10 linux distrib on [distrowatch](https://distrowatch.com/) if you don't have your favorite one yet)

- Having **bash** installed (most distro allready ship it)

- Install **Docker** or Podman to build OCI compactible image format

- [**Golang**](https://golang.org/) sdk installed if you're on the backend side

- Having generated 2 [**personal access token**](https://gitlab.com/profile/personal_access_tokens) one for **go proxy authentication** and the other for **gitlab container registry** access (save the first one on a [~/.netrc](https://medium.com/@jwenz723/fetching-private-go-modules-during-docker-build-5b76aa690280) file and consume the other with `docker login registry.gitlab.com`)

- A free account on [TreeScale](https://treescale.com/) for the fastest **container registry** available (with token and docker login like above)

- Of course a **reliable internet connection** 🤷‍♀️

## 🔎 What's included

- A dummy install script for [**k3s**](https://k3s.io/) lightweight Kubernetes implementation

- Another dummy install script for [Arangodb](https://www.arangodb.com/) multi model database

- Another one for [MinIO](https://min.io/) S3 compactible object storage program

- And finaly another install script for [Openfaas](https://www.openfaas.com/) serverless style functions gateway

- A [Graphql](https://gqlgen.com/) Gateway (Golang) in [backend/api](./backend/api) directory

- A [Leboncoin](https://www.leboncoin.fr/) scrapper (Golang) in [scrappers/leboncoin](./backend/scrappers/leboncoin)

- A [NameEntityRecognition](https://spacy.io/api/entityrecognizer) service (Python) in [ner/spacy](./backend/ner/spacy) dir

## 🏁 Getting started

- First install k3s on your local machine by running `make kubeinstall`

- If everythings allright, a `k3s kubectl get nodes` command should return your local master node, if not [ask for help](https://stackoverflow.com/help/how-to-ask) on Teams channel (something like #Renty - Dev)

- You may want to install the kubernetes dashboard GUI with [**kubernetes_dashboard_install.sh**](./dev-setup/kubernetes_dashboard_install.sh) (optional)

- Then deploy all services and app this repo contains with `make serviceinstall`

## 🔧 Troubleshooting

- Can't get kube to pull my images ...
Check if your authentication token(s) are in the `.credential` file, if not add them and rerun `k3s_registry_secrets.sh`

- Can't push to container registry `repo.treescale.com/dave-lopeur/kubebeber/*` ...
Sorry it's my (@dave-lopeur) fault, i've forgot to put the common gitlab registry in all images references , but i've made a **change_registry.sh** script for that, have a look , udpate if needed and run it !

## Services documentation

1. [MinIO](/minio/)

2. [API](/api/)

3. [Renty-web](/renty-web/)

4. [OpenFAAS](/openfaas/)

5. [Spacy](/spacy/)

6. [Inception](/inception/)

7. [Scrappers](/scrappers/)



