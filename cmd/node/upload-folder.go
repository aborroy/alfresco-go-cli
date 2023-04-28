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

var folderName string
var outputResponseBody bytes.Buffer
var wg sync.WaitGroup
var nodeUploadFolderCmd = &cobra.Command{
	Use:   "upload-folder",
	Short: "Upload local folder to Alfresco Repository",
	Run: func(command *cobra.Command, args []string) {
		log.Println(NodeUploadFolderCmdId,
			"Uploading local folder "+folderName+" to Alfresco Repository folder "+nodeId)
		tree := make(map[string]string)
		var hiddenPaths []string
		err := filepath.WalkDir(folderName,
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
							response := CreateNode(command, parentId, info.Name(), TypeFolder, nil, nil, "")
							var node Node
							json.Unmarshal(response.Bytes(), &node)
							tree[path] = node.Entry.ID
							if folderName == path {
								outputResponseBody = response
							}
							log.Println(NodeUploadFolderCmdId, "Folder "+path+" has been uploaded")
						} else {
							parentId := tree[parentPath]
							wg.Add(1)
							go createFile(command, parentId, path, info)
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
		wg.Wait()
		if err != nil {
			cmd.ExitWithError(NodeUploadFolderCmdId, err)
		}
		var format, _ = command.Flags().GetString("output")
		outputNode(outputResponseBody.Bytes(), format)
		log.Println(NodeUploadFolderCmdId,
			"Uploaded local folder "+folderName+" to Alfresco Repository folder "+nodeId)
	},
}

func createFile(cmd *cobra.Command, parentId string, path string, info fs.DirEntry) {
	response := CreateNode(cmd, parentId, info.Name(), TypeContent, nil, nil, path)
	var node Node
	json.Unmarshal(response.Bytes(), &node)
	log.Println(NodeUploadFolderCmdId, "File "+path+" has been uploaded")
	wg.Done()
}

func init() {
	nodeCmd.AddCommand(nodeUploadFolderCmd)
	nodeUploadFolderCmd.Flags().StringVarP(&folderName, "directory", "d", "", "Folder to be uploaded (complete or local path)")
}
