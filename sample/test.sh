#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

# Alfresco CLI program location
ALF="../alfresco"

# Configure Alfresco CLI for a local ACS in port 8080 with default credentials (admin/admin)
# ACS must be up & ready before running this script
$ALF config set -s http://localhost:8080/alfresco -u admin -p admin

# Get Node ID for folder "Shared", that is one of the childrens of node "Root"
SHARED_ID=$($ALF node list -i -root- | grep Shared | awk -F ' ' '{print $1}')

# Create a sample file to be uploaded
echo 'This is a test file' > file.txt

# Create a new file under folder "Shared" with some properties and the content of file.txt
NEW_NODE_ID=$($ALF node create -n file.txt \
-i $SHARED_ID \
-t cm:content -p cm:title="Title" -p cm:description="Description" \
-f file.txt | grep file.txt | awk -F ' ' '{print $1}')

# Remove sample file
rm -rf file.txt

# Get Node ID for new node
GET_NEW_NODE_ID=$($ALF node get -i $NEW_NODE_ID -o id | awk -F ' ' '{print $1}')

if [[ "$GET_NEW_NODE_ID" != "$NEW_NODE_ID" ]]; then
  echo "Got $GET_NEW_NODE_ID but expecting $NEW_NODE_ID"
  exit -1
fi

# Remove node created before
$ALF node delete -i $GET_NEW_NODE_ID

# Verify the node has been effectively removed
GET_NEW_NODE_ID=$($ALF node get -i $NEW_NODE_ID -o id | awk -F ' ' '{print $1}')

if [ "${GET_NEW_NODE_ID}" != "" ]; then
  echo "Got $GET_NEW_NODE_ID but expecting EMPTY"
  exit -1
fi

# Remove local configuration file created by "config" command
rm .alfresco