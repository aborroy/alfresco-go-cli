#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

# Cleanup resources
function cleanup {
  EXIT_CODE=$?
  set +e
  docker compose down
  exit $EXIT_CODE
}
trap cleanup EXIT

echo "[INTEGRATION TESTS for GROUP commands]"

# Start Alfresco
docker compose up -d
echo "Starting Alfresco ..."
bash -c 'while [[ "$(curl -s -o /dev/null -w ''%{http_code}'' http://localhost:8080/alfresco/s/api/server)" != "200" ]]; do sleep 5; done'

# Alfresco CLI program location
ALF="../alfresco"

# Configure Alfresco CLI for a local ACS in port 8080 with default credentials (admin/admin)
# ACS must be up & ready before running this script
$ALF config set -s http://localhost:8080/alfresco -u admin -p admin
echo "Credentials stored"

# Get Name for group "GROUP_ALFRESCO_ADMINISTRATORS"
ADMIN_NAME=$($ALF group list | grep GROUP_ALFRESCO_ADMINISTRATORS | awk -F ' ' '{print $2}')
echo "Retrieved admin group using "list": $ADMIN_NAME"

if [[ "$ADMIN_NAME" != "ALFRESCO_ADMINISTRATORS" ]]; then
  echo "Got $ADMIN_NAME but expecting ALFRESCO_ADMINISTRATORS"
  exit -1
fi

# Create group "alfresco"
ALFRESCO_GROUP_ID=$($ALF group create -i GROUP_ALFRESCO -d Alfresco -o id)
echo "Created new group alfresco using "get": $ALFRESCO_GROUP_ID"

if [[ "$ALFRESCO_GROUP_ID" != "GROUP_ALFRESCO" ]]; then
  echo "Got $ALFRESCO_GROUP_ID but expecting GROUP_ALFRESCO"
  exit -1
fi

# Modify group "alfresco"
ALFRESCO_GROUP_NAME=$($ALF group update -i GROUP_ALFRESCO -d Hyland | grep GROUP_ALFRESCO | awk -F ' ' '{print $2}')
echo "Modified group alfresco name using "update": $ALFRESCO_GROUP_NAME"
ALFRESCO_GROUP_NAME=$($ALF group get -i GROUP_ALFRESCO | grep GROUP_ALFRESCO | awk -F ' ' '{print $2}')

if [[ "$ALFRESCO_GROUP_NAME" != "Hyland" ]]; then
  echo "Got $ALFRESCO_GROUP_NAME but expecting Hyland"
  exit -1
fi

# Add user "admin" to group "alfresco" as member
ALFRESCO_USER_NAME=$($ALF group add -i GROUP_ALFRESCO -a admin -t PERSON | grep admin | awk -F ' ' '{print $2}')
echo "Added user admin to group alfresco"

if [[ "$ALFRESCO_USER_NAME" != "admin" ]]; then
  echo "Got $ALFRESCO_USER_NAME but expecting admin"
  exit -1
fi

# Remove group "alfresco"
$ALF group delete -i $ALFRESCO_GROUP_ID
echo "Group $ALFRESCO_GROUP_ID has been deleted"

# Stop Alfresco
docker compose down