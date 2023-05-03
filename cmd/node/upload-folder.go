package node

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/util"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

const NodeUploadFolderCmdId string = "[NODE UPLOAD FOLDER]"

var folderNameUpload string
var wgUpload sync.WaitGroup
var nodeUploadFolderCmd = &cobra.Command{
	Use:   "upload-folder",
	Short: "Upload local folder to Alfresco Repository",
	Run: func(command *cobra.Command, args []string) {

		if relativePath != "" {
			nodeId = GetNodeId(nodeId, relativePath)
		}

		log.Println(NodeUploadFolderCmdId,
			"Uploading local folder "+folderNameUpload+" to Alfresco Repository folder "+nodeId)

		tree := make(map[string]string)
		var hiddenPaths []string
		err := filepath.WalkDir(folderNameUpload,
			func(path string, info fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if hidden, _ := util.IsHiddenFile(info.Name()); !hidden {
					parentPath := string(path[0:strings.LastIndex(path, "/")])
					if !slices.Contains(hiddenPaths, parentPath) {
						if info.IsDir() {
							parentId := tree[parentPath]
							if parentId == "" {
								parentId = nodeId
							}
							var localResponseBody bytes.Buffer
							CreateNode(parentId, info.Name(), TypeFolder, "", nil, nil, "", &localResponseBody)
							var node Node
							json.Unmarshal(localResponseBody.Bytes(), &node)
							tree[path] = node.Entry.ID
							if folderNameUpload == path {
								responseBody = localResponseBody
							}
							log.Println(NodeUploadFolderCmdId, "Folder "+path+" has been uploaded")
						} else {
							parentId := tree[parentPath]
							wgUpload.Add(1)
							go createFile(parentId, path, info)
						}
					} else {
						if info.IsDir() {
							hiddenPaths = append(hiddenPaths, path)
							log.Println(NodeUploadFolderCmdId, path+"path is hidden, it will be ignored")
						}
					}
				} else {
					hiddenPaths = append(hiddenPaths, path)
					log.Println(NodeUploadFolderCmdId, path+"path is hidden, it will be ignored")
				}
				return nil
			})
		wgUpload.Wait()
		if err != nil {
			cmd.ExitWithError(NodeUploadFolderCmdId, err)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(NodeUploadFolderCmdId,
			"Uploaded local folder "+folderNameUpload+" to Alfresco Repository folder "+nodeId)
	},
}

func createFile(parentId string, path string, info fs.DirEntry) {
	defer wgUpload.Done()
	var localResponseBody bytes.Buffer
	CreateNode(parentId, info.Name(), TypeContent, "", nil, nil, path, &localResponseBody)
	var node Node
	json.Unmarshal(localResponseBody.Bytes(), &node)
	log.Println(NodeUploadFolderCmdId, "File "+path+" has been uploaded")
}

func init() {
	nodeCmd.AddCommand(nodeUploadFolderCmd)
	nodeUploadFolderCmd.Flags().StringVarP(&nodeId, "nodeId", "i", "", "Parent Node Id in Alfresco Repository to upload local folder")
	nodeUploadFolderCmd.Flags().StringVarP(&relativePath, "relativePath", "r", "", "A path relative to the nodeId.")
	nodeUploadFolderCmd.Flags().StringVarP(&folderNameUpload, "directory", "d", "", "Local folder to be uploaded (complete path)")
	nodeUploadFolderCmd.Flags().SortFlags = false
	nodeUploadFolderCmd.MarkFlagRequired("nodeId")
	nodeUploadFolderCmd.MarkFlagRequired("directory")
}
