package httpclient

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/nativestore"
)

type Format string

const (
	Json    Format = "json"
	Content Format = "content"
	None    Format = "none"
)

const HttpClientId string = "[HTTP]"

type HttpExecution struct {
	Method             string
	Data               string
	Url                string
	Parameters         url.Values
	Format             Format
	ResponseBodyOutput io.Writer
}

var validHttpResponse = map[int]bool{
	http.StatusOK:        true,
	http.StatusCreated:   true,
	http.StatusNoContent: true,
}

func setBasicAuthHeader(request *http.Request, username, password string) {
	if cmd.UsernameParam != "" {
		username = cmd.UsernameParam
		password = cmd.PasswordParam
	}
	auth := username + ":" + password
	request.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))
}

func createHttpClient(tlsEnabled bool, insecureAllowed bool) *http.Client {
	var tlsConfig *tls.Config
	if tlsEnabled && insecureAllowed {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return &client
}

func checkStatusResponse(httpStatus int) {
	if !validHttpResponse[httpStatus] {
		fmt.Println("HTTP ERROR", httpStatus, http.StatusText(httpStatus))
		log.Println("HTTP ERROR", httpStatus, http.StatusText(httpStatus))
		os.Exit(1)
	}
}

func Execute(execution *HttpExecution) error {

	var storedServer, username, password, tlsEnabled, insecureAllowed = nativestore.GetDetails()

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
		cmd.ExitWithError(HttpClientId, err)
	}
	setBasicAuthHeader(request, username, password)
	if execution.Format == Json {
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	response, err := createHttpClient(tlsEnabled, insecureAllowed).Do(request)
	if err != nil {
		cmd.ExitWithError(HttpClientId, err)
	}
	checkStatusResponse(response.StatusCode)
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(storedServer + execution.Url + " - Failed to close response body - " + err.Error())
		}
	}()
	_, err = io.Copy(execution.ResponseBodyOutput, response.Body)
	return err
}

func ExecuteUploadContent(execution *HttpExecution) error {

	var storedServer, username, password, tlsEnabled, insecureAllowed = nativestore.GetDetails()

	r, w := io.Pipe()
	request, err := http.NewRequest(execution.Method, storedServer+execution.Url, r)
	if err != nil {
		cmd.ExitWithError(HttpClientId, err)
	}
	setBasicAuthHeader(request, username, password)

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

	response, err := createHttpClient(tlsEnabled, insecureAllowed).Do(request)
	if err != nil {
		cmd.ExitWithError(HttpClientId, err)
	}
	checkStatusResponse(response.StatusCode)
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(storedServer + execution.Url + " - Failed to close response body - " + err.Error())
		}
	}()
	_, err = io.Copy(execution.ResponseBodyOutput, response.Body)
	return err

}

func ExecuteDownloadContent(execution *HttpExecution) error {

	var storedServer, username, password, tlsEnabled, insecureAllowed = nativestore.GetDetails()

	out, err := os.Create(execution.Data)
	if err != nil {
		cmd.ExitWithError(HttpClientId, err)
	}
	defer out.Close()

	var payload io.Reader
	request, err := http.NewRequest(execution.Method, storedServer+execution.Url, payload)
	if err != nil {
		log.Println(err)
		return err
	}
	setBasicAuthHeader(request, username, password)
	response, err := createHttpClient(tlsEnabled, insecureAllowed).Do(request)
	if err != nil {
		cmd.ExitWithError(HttpClientId, err)
	}
	checkStatusResponse(response.StatusCode)
	io.Copy(out, response.Body)
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(storedServer + execution.Url + " - Failed to close response body - " + err.Error())
		}
	}()
	_, err = io.Copy(execution.ResponseBodyOutput, response.Body)
	return err

}
