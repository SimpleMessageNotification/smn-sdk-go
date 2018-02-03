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

	// create message template
	CreateMessageTemplate(&tempClient)
	// query message template detail
	QueryMessageTemplateDetail(&tempClient)
	// list message template
	ListMessageTemplate(&tempClient)
	// update message template
	UpdateMessageTemplate(&tempClient)
	// delete message template
	DeleteMessageTemplate(&tempClient)
}

func CreateMessageTemplate(smnClient *client.SmnClient) {
	request := smnClient.NewCreateMessageTemplateRequest()
	request.Protocol = "sms"
	request.MessageTemplateName = "template_test_by_zhangyx_for_go"
	request.Content = "this is the first template for go {time}"
	response, err := smnClient.CreateMessageTemplate(request)
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

func QueryMessageTemplateDetail(smnClient *client.SmnClient) {
	request := smnClient.NewQueryMessageTemplateDetailRequest()
	request.MessageTemplateId = "d4c4968551534a298233c92b6c4916c7"
	response, err := smnClient.QueryMessageTemplateDetail(request)
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

func ListMessageTemplate(smnClient *client.SmnClient) {
	request := smnClient.NewListMessageTemplateRequest()
	request.Offset = 0
	request.Limit = 10
	request.MessageTemplateName = "template_test_by_zhangyx_for_go"
	response, err := smnClient.ListMessageTemplate(request)
	fmt.Println(request.GetUrl())

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

func UpdateMessageTemplate(smnClient *client.SmnClient) {
	request := smnClient.NewUpdateMessageTemplateRequest()
	request.MessageTemplateId = "d4c4968551534a298233c92b6c4916c7"
	request.Content = "this is the first template for go {time} v2"
	response, err := smnClient.UpdateMessageTemplate(request)
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

func DeleteMessageTemplate(smnClient *client.SmnClient) {
	request := smnClient.NewDeleteMessageTemplateRequest()
	request.MessageTemplateId = "d4c4968551534a298233c92b6c4916c7"
	response, err := smnClient.DeleteMessageTemplate(request)
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
