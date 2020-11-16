#!/usr/bin/env bash
set -e
fname="todo_$(date +'%m-%d-%y').txt"
curl -s 'http://api.192-168-1-34.sslip.io/api/query' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Origin: http://api.192-168-1-34.sslip.io' -d '{"query":"query {\n  rentoffers {\n    description {\n      value\n    }\n  }\n}\n"}' --compressed \
    | jq '.data.rentoffers' > todo.json
cat todo.json | jq -r '.[].description[0].value' > "to_annotate/$fname"
# awk 'NF' "to_annotate/$fname"
sed -i '/^$/d' "to_annotate/$fname"
rm todo.json