#!/usr/bin/env bash
set -e
bold="\e[1m"
clear="\e[0m"
service_dirs=("backend/api" "backend/scrappers/leboncoin" "backend/seed")
funct_dirs=$(find backend/openfaas -maxdepth 1 -not -path "*/build" -type d -printf '%f\n')
for VARIABLE in $funct_dirs
do
    [[ $VARIABLE == "openfaas" || $VARIABLE == "template" ]] && continue
    fpath="backend/openfaas/$VARIABLE"
    echo -e "[INFO] Will update go deps in $bold$fpath$clear"
	  cd $fpath && go get -u || exit
    cd -
done
for path in ${service_dirs[@]};
do
    echo -e "[INFO] Will update go deps in $bold$path$clear"
	cd $path && go get -u || exit
    cd -
done

