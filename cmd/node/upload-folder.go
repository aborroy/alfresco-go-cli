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
var outputResponseBody bytes.Buffer
var wgUpload sync.WaitGroup
var nodeUploadFolderCmd = &cobra.Command{
	Use:   "upload-folder",
	Short: "Upload local folder to Alfresco Repository",
	Run: func(command *cobra.Command, args []string) {
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
							response := CreateNode(parentId, info.Name(), TypeFolder, nil, nil, "")
							var node Node
							json.Unmarshal(response.Bytes(), &node)
							tree[path] = node.Entry.ID
							if folderNameUpload == path {
								outputResponseBody = response
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
		var format, _ = command.Flags().GetString("output")
		outputNode(outputResponseBody.Bytes(), format)
		log.Println(NodeUploadFolderCmdId,
			"Uploaded local folder "+folderNameUpload+" to Alfresco Repository folder "+nodeId)
	},
}

func createFile(parentId string, path string, info fs.DirEntry) {
	response := CreateNode(parentId, info.Name(), TypeContent, nil, nil, path)
	var node Node
	json.Unmarshal(response.Bytes(), &node)
	log.Println(NodeUploadFolderCmdId, "File "+path+" has been uploaded")
	wgUpload.Done()
}

func init() {
	nodeCmd.AddCommand(nodeUploadFolderCmd)
	nodeUploadFolderCmd.Flags().StringVarP(&folderNameUpload, "directory", "d", "", "Local folder to be uploaded (complete or local path)")
}
