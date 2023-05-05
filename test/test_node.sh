#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

# Cleanup resources
function cleanup {
  EXIT_CODE=$?
  set +e
  docker compose down
  rm -rf file-alfresco.txt file-alfresco-modified.txt file.txt
  rm -rf sample sample-2
  rm -rf .alfresco alfresco.log
  exit $EXIT_CODE
}
trap cleanup EXIT

echo "[INTEGRATION TESTS for NODE commands]"

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

# Get Node ID for folder "Shared", that is one of the childrens of node "Root"
SHARED_ID=$($ALF node list -i -root- | grep Shared | awk -F ' ' '{print $1}')
echo "Retrieved Shared Folder NodeId using "list": $SHARED_ID"

# Get Node ID for folder "-root-/Shared"
SHARED_ID_GET=$($ALF node get -i -root- -r Shared -o id)
echo "Retrieved Shared Folder NodeId using "get": $SHARED_ID_GET"

if [[ "$SHARED_ID" != "$SHARED_ID_GET" ]]; then
  echo "Got $SHARED_ID_GET but expecting $SHARED_ID"
  exit -1
fi

# Get Node ID for folder "Shared" using NodeId
SHARED_ID_GET_ID=$($ALF node get -i $SHARED_ID -o id)
echo "Retrieved Shared Folder NodeId using "get": $SHARED_ID_GET_ID"

if [[ "$SHARED_ID" != "$SHARED_ID_GET_ID" ]]; then
  echo "Got $SHARED_ID_GET_ID but expecting $SHARED_ID"
  exit -1
fi

# Create a sample file to be uploaded
echo 'This is a test file' > file.txt

# Create a new file under folder "Shared" with some properties and the content of file.txt
NEW_NODE_ID=$($ALF node create -n file-alfresco.txt \
-i $SHARED_ID \
-t cm:content -p cm:title="Title" -p cm:description="Description" \
-f file.txt | grep file-alfresco.txt | awk -F ' ' '{print $1}')
echo "Created new node $NEW_NODE_ID under folder with NodeId: $SHARED_ID"

# Download file content
DOWNLOADED_NODE_ID=$($ALF node get -i $NEW_NODE_ID -d . -o id)
cmp --silent file.txt file-alfresco.txt
echo "Downloaded file equals to original file"

# Modify node name
MODIFIED_NODE_ID=$($ALF node update -i -root- -r Shared/file-alfresco.txt -n file-alfresco-modified.txt -o id)
echo "Updated node $MODIFIED_NODE_ID to new name"

MODIFIED_NODE_ID=$($ALF node get -i $MODIFIED_NODE_ID -d . -o id)
cmp --silent file.txt file-alfresco-modified.txt
echo "Downloaded modified file equals to original file"

# Remove sample files
rm -rf file-alfresco.txt file-alfresco-modified.txt file.txt

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

if [[ "${GET_NEW_NODE_ERROR}" != *"404"* ]]; then
  echo "Got $GET_NEW_NODE_ERROR but expecting 404 (Not Found)"
  exit -1
fi
echo "Node $GET_NEW_NODE_ID delete verification ok"

# Upload local folder tree
mkdir sample sample/docs sample/files
echo 'This is a test doc' > doc.txt
echo 'This is a test file' > file.txt
echo 'This is a test' > test.txt
cp test.txt sample
cp doc.txt sample/docs/doc-1.txt
cp doc.txt sample/docs/doc-2.txt
cp file.txt sample/files/file-1.txt
cp file.txt sample/files/file-2.txt
mv doc.txt sample
mv file.txt sample
mv test.txt sample
UPLOADED_FOLDER_ID=$($ALF node upload-folder -i $SHARED_ID -d $(pwd)/sample)
echo "Folder $(pwd)/sample has been uploaded to Alfresco"

# Modify folder name
MODIFIED_UPLOADED_FOLDER_ID=$($ALF node update -i -root- -r Shared/sample -n sample-2 -o id)
echo "Updated folder $MODIFIED_UPLOADED_FOLDER_ID to new name"

# Download Repository folder
DOWNLOADED_FOLDER_ID=$($ALF node download-folder -i -root- -r Shared/sample-2 -d . -o id)
echo "Folder sample-2 has been downloaded from Alfresco"

# Compare uploaded & downloaded repository
FOLDER_DIFF=$(diff -r sample sample-2)
if [[ -n "$FOLDER_DIFF" ]]; then
  echo "Found differences between uploaded and downloaded folder: $FOLDER_DIFF"
  exit -1
fi
echo "Folder sample & sample-2 are identical"

# Remove temporal folders
rm -rf sample sample-2

# Stop Alfresco
docker compose down

# Remove local configuration file created by "config" command
rm .alfresco
rm alfresco.log