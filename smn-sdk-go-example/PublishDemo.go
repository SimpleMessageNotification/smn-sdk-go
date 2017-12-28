package main

import (
	"github.com/SimpleMessageNotification/smn-sdk-go/smn-sdk-go/client"
	"fmt"
	//"net/http"
	//"github.com/SimpleMessageNotification/smn-sdk-go/smn-sdk-go/commom"
	//"time"
	//"net/url"
	//"crypto/tls"
)

func main() {
	// init client
	tempClient, err := client.NewClient(
		"YourAccountUserName",
		"YourAccountDomainName",
		"YourAccountPassword",
		"YourRegionName")

	// if you want to config http. you can use like this
	//proxy := func(_ *http.Request) (*url.URL, error) {
	//	return url.Parse("http://127.0.0.1:8080")
	//}
	//transport := &http.Transport{
	//	Proxy: proxy,
	//	TLSClientConfig:&tls.Config{InsecureSkipVerify:true}}
	//config := commom.NewClientConfig()
	//config.SetTimeout(20 * time.Second)
	//config.SetTransport(transport)
	//
	//tempClient, err := client.NewClientWithConfig(
	//	"YourAccountUserName",
	//	"YourAccountDomainName",
	//	"YourAccountPassword",
	//	"YourRegionName",
	//	config)
	if err != nil {
		panic(err)
	}

	//publish message
	PublishMessage(&tempClient)
	// publish message structure
	PublishMessageStructure(&tempClient)
	// publish message template
	PublishMessageTemplate(&tempClient)
}

func PublishMessage(smnClient *client.SmnClient) {
	request := smnClient.NewPublishMessageRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	request.Message = "test by zhangyx"
	response, err := smnClient.PublishMessage(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func PublishMessageStructure(smnClient *client.SmnClient) {
	request := smnClient.NewPublishMessageStructureRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	request.MessageStructure = "{" +
		"\"default\":\"test by zhangyx structure\"," +
		"\"email\":\"test by zhangyx structure email\"," +
		"\"sms\":\"test by zhangyx structure _sms\"}"
	response, err := smnClient.PublishMessageStructure(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func PublishMessageTemplate(smnClient *client.SmnClient) {
	request := smnClient.NewPublishMessageTemplateRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	request.MessageTemplateName = "createMessageTemplate"
	request.Tags["year"] = "2016"
	request.Tags["company"] = "hellokitty"
	response, err := smnClient.PublishMessageTemplate(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}
