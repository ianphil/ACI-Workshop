#!/usr/bin/env bash

az group create --name eastoffsite --location eastus
az container create --resource-group eastoffsite --file deploy-aci.yaml

# az container show --resource-group eastoffsite --name eastOffsiteGroup --output table
# az container logs --resource-group eastoffsite --name eastOffsiteGroup --container-name send
# az container logs --resource-group eastoffsite --name eastOffsiteGroup --container-name receive
# az container logs --resource-group eastoffsite --name eastOffsiteGroup --container-name show