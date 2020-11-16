#!/usr/bin/env bash
set -e

get_latest_release() {
  curl -s "https://api.github.com/repos/$1/releases/latest" |       # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}

download_kubesec_bin() {
    release=$(get_latest_release controlplaneio/kubesec)
    echo -e "[INFO] Found release $release"
    wget "https://github.com/controlplaneio/kubesec/releases/download/$release/kubesec_linux_amd64.tar.gz"
    tar xvf kubesec_linux_amd64.tar.gz && rm kubesec_linux_amd64.tar.gz && chmod +x kubesec
}

echo -e "[INFO] Installing kubesec cli if missing"
[ -f kubesec ] || download_kubesec_bin
echo -e "[INFO] Test if everything all right"
./kubesec --help
