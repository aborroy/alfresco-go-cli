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

echo "[INTEGRATION TESTS for PEOPLE commands]"

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

# Get Name for user "admin"
ADMIN_NAME=$($ALF people list | grep admin | awk -F ' ' '{print $2}')
echo "Retrieved admin name using "list": $ADMIN_NAME"

if [[ "$ADMIN_NAME" != "Administrator" ]]; then
  echo "Got $ADMIN_NAME but expecting Administrator"
  exit -1
fi

# Create user "alfresco"
ALFRESCO_USER_ID=$($ALF people create -i alfresco --password alfresco -f Alfresco -l Hyland -e alfresco@alfresco.com -o id)
echo "Created new user alfresco using "get": $ALFRESCO_USER_ID"

if [[ "$ALFRESCO_USER_ID" != "alfresco" ]]; then
  echo "Got $ALFRESCO_USER_ID but expecting alfresco"
  exit -1
fi

# Modify user "alfresco"
ALFRESCO_USER_NAME=$($ALF people update -i alfresco -f Hyland | grep alfresco | awk -F ' ' '{print $2}')
echo "Modified user alfresco name using "update": $ALFRESCO_USER_NAME"
ALFRESCO_USER_NAME=$($ALF people get -i alfresco | grep alfresco | awk -F ' ' '{print $2}')

if [[ "$ALFRESCO_USER_NAME" != "Hyland" ]]; then
  echo "Got $ALFRESCO_USER_NAME but expecting Hyland"
  exit -1
fi

# Remove user "alfresco"
$ALF people delete -i $ALFRESCO_USER_ID
echo "Person $ALFRESCO_USER_ID has been deleted"

# Verify the person has been effectively removed (it will provoke an error)
set +o errexit
GET_NEW_NODE_ERROR=$($ALF people get -i $ALFRESCO_USER_ID)
set -o errexit

if [[ "${GET_NEW_NODE_ERROR}" != *"404"* ]]; then
  echo "Got $GET_NEW_NODE_ERROR but expecting 404 (Not Found)"
  exit -1
fi
echo "Person $ALFRESCO_USER_ID delete verification ok"

# Stop Alfresco
docker compose down