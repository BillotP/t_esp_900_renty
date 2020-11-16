#!/usr/bin/env bash
bold="\e[1m"
clear="\e[0m"
HOST=$(ifconfig wlp2s0 | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p')
HOSTURL=$(echo -n $HOST | sed 's/\./-/g')
echo -e "[INFO] 🟢 Updating $bold api $clear ingress"
cd backend/api && \
    sed -i  "s/api.*.sslip.io/api.$HOSTURL.sslip.io/g" deployment.yaml && \
    k3s kubectl apply -f deployment.yaml && \
    cd -
echo -e "[INFO] 🟢 Updating $bold minio bucket $clear ingress"
cd infra/minio && \
    sed -i  "s/bucket.*.sslip.io/bucket.$HOSTURL.sslip.io/g" minio-proxy.yaml && \
    k3s kubectl apply -f minio-proxy.yaml && \
    cd -
echo -e "[INFO] 🟢 Updating $bold my-app $clear ingress"
cd frontend/renty-web && \
    sed -i  "s/app.*.sslip.io/app.$HOSTURL.sslip.io/g" deployment.yaml && \
    k3s kubectl apply -f deployment.yaml && \
    cd -
echo -e "[INFO] 🟢 Updating $bold openfaas gateway $clear ingress"
cd infra/openfaas && \
    sed -i  "s/openfaas.*.sslip.io/openfaas.$HOSTURL.sslip.io/g" ingress.yaml && \
    k3s kubectl apply -f ingress.yaml && \
    cd -
# echo -e "[INFO] 🟢 Updating $bold auth function $clear ingress"
# cd openfaas/auth && \
#     sed -i  "s/openfaas.*.sslip.io/openfaas.$HOSTURL.sslip.io/g" ingress.yaml && \
#     k3s kubectl apply -f ingress.yaml && \
#     cd -
echo -e "[INFO] 🚀 All done enjoy your services on new domain =>$bold$HOSTURL.sslip.io$clear<="