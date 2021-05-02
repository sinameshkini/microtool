package rest

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	netUrl "net/url"
	"github.com/spf13/viper"
)

type Api struct {
	Scheme      string
	Host        string
	Port        string
	ApiPrefix   string
	Url         string
	contentType string
	accessToken string
}

func New(serviceConn string) *Api {
	var (
		schema      = viper.GetString(serviceConn + "schema")
		host        = viper.GetString(serviceConn + "host")
		port        = viper.GetString(serviceConn + "port")
		apiPrefix   = viper.GetString(serviceConn + "api_prefix")
		contentType = viper.GetString(serviceConn + "content_type")
		token       = viper.GetString(serviceConn + "token")
	)

	if port != "" {
		port = ":" + port
	}

	restClient := &Api{
		Scheme:      schema,
		Host:        host,
		Port:        port,
		ApiPrefix:   apiPrefix,
		accessToken: token,
		contentType: contentType,
	}

	restClient.Url = fmt.Sprintf("%s://%s%s%s",
		restClient.Scheme,
		restClient.Host,
		restClient.Port,
		restClient.ApiPrefix,
	)

	return restClient
}

type HttpResponse struct {
	StatusCode int
	Status     string
	Body       []byte
}

func (a *Api) Get(reqUrl string) (response *HttpResponse, err error) {
	var (
		url      string
		resp     *http.Response
		respBody []byte
		httpReq  *http.Request
	)

	url = a.Url + reqUrl

	if httpReq, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}

	httpReq.Header.Add("Authorization", a.accessToken)
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//httpReq.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	httpReq.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err = client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	response = &HttpResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       respBody,
	}

	if viper.GetBool("debug") {
		fmt.Println("__________________ Get Request To: ________________________")
		fmt.Println(url)
		fmt.Println("__________________ Response: ________________________")
		fmt.Println("status code: ", response.StatusCode)
		fmt.Println("status: ", response.Status)
		fmt.Println("body: ", string(response.Body))
	}

	return response, nil
}

func (a *Api) Post(req map[string]string, methodPath string) (response *HttpResponse, err error) {
	var (
		httpReq  *http.Request
		resp     *http.Response
		respBody []byte
		url      string
	)

	url = a.Url + methodPath
	data := netUrl.Values{}
	for key, param := range req {
		data.Set(key, param)
	}

	httpReq, err = http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Authorization", a.accessToken)
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	httpReq.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err = client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response = &HttpResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       respBody,
	}

	if viper.GetBool("debug") {
		fmt.Println("__________________ Post Request To: ________________________")
		fmt.Println(url)
		fmt.Println(req)
		fmt.Println("__________________ Response: ________________________")
		fmt.Println("status code: ", response.StatusCode)
		fmt.Println("status: ", response.Status)
		fmt.Println("body: ", string(response.Body))
	}

	return response, nil
}

func (a *Api) PostFile(url string, relativePath string) (response *HttpResponse, err error) {
	var (
		method   = "POST"
		payload  = &bytes.Buffer{}
		writer   = multipart.NewWriter(payload)
		file     *os.File
		part     io.Writer
		client   = &http.Client{}
		req      *http.Request
		resp     *http.Response
		respBody []byte
	)

	file, err = os.Open(relativePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	part, err = writer.CreateFormFile("file", filepath.Base(relativePath))
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	response = &HttpResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       respBody,
	}

	if viper.GetBool("debug") {
		fmt.Println("__________________ Post Request To: ________________________")
		fmt.Println(url)
		fmt.Println(req)
		fmt.Println("__________________ Response: ________________________")
		fmt.Println("status code: ", response.StatusCode)
		fmt.Println("status: ", response.Status)
		fmt.Println("body: ", string(response.Body))
	}

	return response, nil
}
