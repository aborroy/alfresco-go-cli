package node

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
)

const NodeDownloadFolderCmdId string = "[NODE DOWNLOAD FOLDER]"

const DefaultMaxItems int = 10

func getChildren(folderId string, folderName string, skipCount int, maxItems int) {
	response := ListNode(folderId, strconv.Itoa(skipCount), strconv.Itoa(maxItems))
	var nodeList NodeList
	json.Unmarshal(response.Bytes(), &nodeList)
	for _, node := range nodeList.List.Entries {
		if node.Entry.IsFolder {
			os.Mkdir(folderNameDownload+"/"+folderName+"/"+node.Entry.Name, os.ModePerm)
			getChildren(node.Entry.ID, folderName+"/"+node.Entry.Name, 0, DefaultMaxItems)
		} else {
			wgDownload.Add(1)
			go downloadContent(node.Entry.ID, folderName, node.Entry.Name)
		}
	}

	var hasMoreItems bool = nodeList.List.Pagination.HasMoreItems
	for hasMoreItems {
		skipCount = skipCount + maxItems
		response := ListNode(folderId, strconv.Itoa(skipCount), strconv.Itoa(maxItems))
		var nodeList NodeList
		json.Unmarshal(response.Bytes(), &nodeList)
		for _, node := range nodeList.List.Entries {
			if node.Entry.IsFolder {
				os.Mkdir(folderNameDownload+"/"+folderName+"/"+node.Entry.Name, os.ModePerm)
				getChildren(node.Entry.ID, folderName+"/"+node.Entry.Name, 0, DefaultMaxItems)
			} else {
				wgDownload.Add(1)
				go downloadContent(node.Entry.ID, folderName, node.Entry.Name)
			}
		}
		hasMoreItems = nodeList.List.Pagination.HasMoreItems
	}
}

var folderNameDownload string
var wgDownload sync.WaitGroup
var nodeDownloadFolderCmd = &cobra.Command{
	Use:   "download-folder",
	Short: "Download Alfresco Repository folder to local folder",
	Run: func(command *cobra.Command, args []string) {

		log.Println(NodeUploadFolderCmdId,
			"Downloading Alfresco folder "+nodeId+" to local folder "+folderNameDownload)

		var response = GetNodeProperties(nodeId)
		var node Node
		json.Unmarshal(response.Bytes(), &node)
		os.Mkdir(folderNameDownload+"/"+node.Entry.Name, os.ModePerm)
		getChildren(nodeId, node.Entry.Name, 0, DefaultMaxItems)

		wgDownload.Wait()

		log.Println(NodeUploadFolderCmdId,
			"Downloaded Alfresco folder "+nodeId+" to local folder "+folderNameDownload)
	},
}

func downloadContent(nodeId string, folderName string, fileName string) {
	GetNodeContent(nodeId, folderName, fileName)
	log.Println(NodeUploadFolderCmdId, "File "+folderName+"/"+fileName+" has been downloaded")
	wgDownload.Done()
}

func init() {
	nodeCmd.AddCommand(nodeDownloadFolderCmd)
	nodeDownloadFolderCmd.Flags().StringVarP(&folderNameDownload, "directory", "d", "", "Folder to download Alfresco content")
}
