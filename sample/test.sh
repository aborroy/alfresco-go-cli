#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

# Alfresco CLI program location
ALF="../alfresco"

# Configure Alfresco CLI for a local ACS in port 8080 with default credentials (admin/admin)
# ACS must be up & ready before running this script
$ALF config set -s http://localhost:8080/alfresco -u admin -p admin
echo "Credentials stored"

# Get Node ID for folder "Shared", that is one of the childrens of node "Root"
SHARED_ID=$($ALF node list -i -root- | grep Shared | awk -F ' ' '{print $1}')
echo "Retrieved Shared Folder NodeId: $SHARED_ID"

# Create a sample file to be uploaded
echo 'This is a test file' > file.txt

# Create a new file under folder "Shared" with some properties and the content of file.txt
NEW_NODE_ID=$($ALF node create -n file.txt \
-i $SHARED_ID \
-t cm:content -p cm:title="Title" -p cm:description="Description" \
-f file.txt | grep file.txt | awk -F ' ' '{print $1}')
echo "Created new node under folder with NodeId: $SHARED_ID"

# Remove sample file
rm -rf file.txt

# Get Node ID for new node
GET_NEW_NODE_ID=$($ALF node get -i $NEW_NODE_ID -o id | awk -F ' ' '{print $1}')
echo "Get NodeId: $GET_NEW_NODE_ID"

if [[ "$GET_NEW_NODE_ID" != "$NEW_NODE_ID" ]]; then
  echo "Got $GET_NEW_NODE_ID but expecting $NEW_NODE_ID"
  exit -1
fi

# Remove node created before
$ALF node delete -i $GET_NEW_NODE_ID
echo "Node $GET_NEW_NODE_ID has been deleted"

# Verify the node has been effectively removed (it will provoke an error)
set +o errexit
GET_NEW_NODE_ERROR=$($ALF node get -i $NEW_NODE_ID)
set -o errexit

if [ "${GET_NEW_NODE_ERROR}" == *"404"* ]; then
  echo "Got $GET_NEW_NODE_ERROR but expecting 404 (Not Found)"
  exit -1
fi

echo "Node $GET_NEW_NODE_ID delete verification ok"

# Remove local configuration file created by "config" command
rm .alfresco