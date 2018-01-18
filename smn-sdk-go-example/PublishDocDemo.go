package main

import (
	"github.com/SimpleMessageNotification/smn-sdk-go/smn-sdk-go/client"
	"fmt"
)

func main() {
	// init client
	smnClient, err := client.NewClient(
		"YourAccountUserName",
		"YourAccountDomainName",
		"YourAccountPassword",
		"YourRegionName")
	if err != nil {
		panic(err)
	}

	request := smnClient.NewPublishMessageRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01_go"
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