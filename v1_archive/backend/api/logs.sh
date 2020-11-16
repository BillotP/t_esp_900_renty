#!/usr/bin/env bash
api_id=$(k3s kubectl get pods | grep api | cut -d " " -f1)
k3s kubectl logs --follow $api_id
