#!/usr/bin/env bash
set -e
bold="\e[1m"
clear="\e[0m"
lan=$(ifconfig wlp2s0 | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p')
echo -e "[INFO] 🟢 Will install =>$bold ArangoDB $clear<="
cd infra/arangodb && ./install.sh && cd -
echo -e "[INFO] 🟢 Will install =>$bold Gloo Gateway $clear<="
cd infra/gloo && ./install.sh && cd -
echo -e "[INFO] 🟢 Will install =>$bold MinIO $clear<="
cd infra/minio && ./install.sh && cd -
echo -e "[INFO] 🟢 Will install =>$bold OpenFAAS $clear<="
./infra/openfaas/install.sh
echo -e "[INFO] 🟢 Will deploy =>$bold frontend/renty-web $clear<="
cd frontend/renty-web && ./up.sh && cd -
echo -e "[INFO] 🟢 Will deploy =>$bold backend/api $clear<="
cd backend/api && ./up.sh && cd -
echo -e "[INFO] Will seed database"
cd backend/seed && GO111MODULE=on go run . && cd -
echo -e "[INFO] 🟢 Updating =>$bold DB_HOST $clear<="
sed -i "s/DB_HOST=.*/DB_HOST=$lan/g" backend/openfaas/.dev.env
./infra/openfaas/set_db_secrets.sh
echo -e "[INFO] 🟢 Will deploy =>$bold OpenFAAS Handlers $clear<="
cd backend/openfaas && ./up.sh && cd -

