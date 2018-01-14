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

	createTopicReqeust := smnClient.NewCreateTopicRequest()
	createTopicReqeust.Name = "zhangyxTestCreate01_go"
	createTopicReqeust.DisplayName = "zhangyxTestCreate01"

	createTopicResponse, err := smnClient.CreateTopic(createTopicReqeust)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !createTopicResponse.IsSuccess() {
		fmt.Printf("%#v\n", createTopicResponse.ErrorResponse)
		return
	}

	subscribeRequest := smnClient.NewSubscribeRequest()
	subscribeRequest.Endpoint = "13688807587"
	subscribeRequest.TopicUrn = createTopicResponse.TopicUrn
	subscribeRequest.Protocol = "sms"
	subscribeRequest.Remark = "email"
	response, err := smnClient.Subscribe(subscribeRequest)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	fmt.Printf("%#v\n", response)
}
