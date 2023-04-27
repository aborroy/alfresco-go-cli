package httpclient

import (
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/aborroy/alfresco-cli/nativestore"
)

type Format string

const (
	Json    Format = "json"
	Content Format = "content"
	None    Format = "none"
)

type HttpExecution struct {
	Method             string
	Data               string
	Url                string
	Parameters         url.Values
	Format             Format
	ResponseBodyOutput io.Writer
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func Execute(execution *HttpExecution) error {

	var storedServer, username, password = nativestore.GetDetails()

	var payload io.Reader
	if execution.Format == Json {
		payload = bytes.NewBufferString(execution.Data)
	}
	var urlStr = storedServer + execution.Url
	if execution.Parameters != nil {
		urlStr = urlStr + "?" + execution.Parameters.Encode()
	}
	request, err := http.NewRequest(execution.Method, urlStr, payload)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", basicAuth(username, password))
	if execution.Format == Json {
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(storedServer + execution.Url + " - Failed to close response body - " + err.Error())
		}
	}()
	_, err = io.Copy(execution.ResponseBodyOutput, response.Body)
	return err
}

func ExecuteUploadContent(execution *HttpExecution) error {

	var storedServer, username, password = nativestore.GetDetails()

	r, w := io.Pipe()
	request, err := http.NewRequest(execution.Method, storedServer+execution.Url, r)
	if err != nil {
		log.Println(err)
		return err
	}
	request.Header.Add("Authorization", basicAuth(username, password))

	go func() {
		defer w.Close()
		file, err := os.Open(execution.Data)
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		if _, err = io.Copy(w, file); err != nil {
			log.Println(err)
			return
		}
	}()

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(storedServer + execution.Url + " - Failed to close response body - " + err.Error())
		}
	}()
	_, err = io.Copy(execution.ResponseBodyOutput, response.Body)
	return err

}
