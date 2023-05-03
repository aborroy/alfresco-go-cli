package node

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/aborroy/alfresco-cli/nativestore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const NodeDownloadFolderCmdId string = "[NODE DOWNLOAD FOLDER]"

func getChildren(folderId string, folderName string, relativePath string, skipCount int, maxItems int) {
	var responseBodyDownload bytes.Buffer
	ListNode(folderId, relativePath, skipCount, maxItems, &responseBodyDownload)
	var nodeList NodeList
	json.Unmarshal(responseBodyDownload.Bytes(), &nodeList)
	for _, node := range nodeList.List.Entries {
		if node.Entry.IsFolder {
			os.Mkdir(folderNameDownload+"/"+folderName+"/"+node.Entry.Name, os.ModePerm)
			getChildren(node.Entry.ID, folderName+"/"+node.Entry.Name, "", 0, viper.GetInt(nativestore.MaxItemsLabel))
		} else {
			wgDownload.Add(1)
			go downloadContent(node.Entry.ID, folderName, node.Entry.Name)
		}
	}

	var hasMoreItems bool = nodeList.List.Pagination.HasMoreItems
	for hasMoreItems {
		skipCount = skipCount + maxItems
		ListNode(folderId, "", skipCount, maxItems, &responseBody)
		var nodeList NodeList
		json.Unmarshal(responseBodyDownload.Bytes(), &nodeList)
		for _, node := range nodeList.List.Entries {
			if node.Entry.IsFolder {
				os.Mkdir(folderNameDownload+"/"+folderName+"/"+node.Entry.Name, os.ModePerm)
				getChildren(node.Entry.ID, folderName+"/"+node.Entry.Name, "", 0, viper.GetInt(nativestore.MaxItemsLabel))
			} else {
				wgDownload.Add(1)
				go downloadContent(node.Entry.ID, folderName, node.Entry.Name)
			}
		}
		hasMoreItems = nodeList.List.Pagination.HasMoreItems
	}
}

func downloadContent(nodeId string, folderName string, fileName string) {
	defer wgDownload.Done()
	GetNodeContent(nodeId, folderName, fileName)
	log.Println(NodeUploadFolderCmdId, "File "+folderName+"/"+fileName+" has been downloaded")
}

var folderNameDownload string
var wgDownload sync.WaitGroup
var nodeDownloadFolderCmd = &cobra.Command{
	Use:   "download-folder",
	Short: "Download Alfresco Repository folder to local folder",
	Run: func(command *cobra.Command, args []string) {

		log.Println(NodeUploadFolderCmdId,
			"Downloading Alfresco folder "+nodeId+"/"+relativePath+" to local folder "+folderNameDownload)

		GetNodeProperties(nodeId, relativePath, &responseBody)
		var node Node
		json.Unmarshal(responseBody.Bytes(), &node)
		os.Mkdir(folderNameDownload+"/"+node.Entry.Name, os.ModePerm)
		getChildren(nodeId, node.Entry.Name, relativePath, 0, viper.GetInt(nativestore.MaxItemsLabel))

		wgDownload.Wait()

	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(NodeUploadFolderCmdId,
			"Downloaded Alfresco folder "+nodeId+"/"+relativePath+" to local folder "+folderNameDownload)
	},
}

func init() {
	nodeCmd.AddCommand(nodeDownloadFolderCmd)
	nodeDownloadFolderCmd.Flags().StringVarP(&nodeId, "nodeId", "i", "", "Node Id in Alfresco Repository to download to local folder. You can also use one of these well-known aliases: -my-, -shared-, -root-")
	nodeDownloadFolderCmd.Flags().StringVarP(&relativePath, "relativePath", "r", "", "A path in Alfresco Repository relative to the nodeId.")
	nodeDownloadFolderCmd.Flags().StringVarP(&folderNameDownload, "directory", "d", "", "Folder to download Alfresco content")
	nodeDownloadFolderCmd.MarkFlagRequired("nodeId")
}
