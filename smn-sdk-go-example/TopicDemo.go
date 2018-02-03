package main

import (
	"github.com/SimpleMessageNotification/smn-sdk-go/smn-sdk-go/client"
	"fmt"
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

	// create topic
	CreateTopic(&tempClient)
	// update topic
	UpdateTopic(&tempClient)
	// delete topic
	DeleteTopic(&tempClient)
	// list topics
	ListTopic(&tempClient)
	// query topic detail
	QueryTopicDetail(&tempClient)
	// update topic attribute
	UpdateTopicAttribute(&tempClient)
	// list topic attributes
	ListTopicAttributes(&tempClient)
	// delete topic attribute by name
	DeleteTopicAttributeByName(&tempClient)
	// delete all topic attributes
	DeleteTopicAttributes(&tempClient)

}

func CreateTopic(smnClient *client.SmnClient) {
	request := smnClient.NewCreateTopicRequest()
	request.Name = "zhangyxTestCreate01"
	request.DisplayName = "zhangyxTestCreate01"

	response, err := smnClient.CreateTopic(request)
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

func UpdateTopic(smnClient *client.SmnClient) {
	request := smnClient.NewUpdateTopicRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	request.DisplayName = "zhangyxTestCreate001"

	response, err := smnClient.UpdateTopic(request)
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

func DeleteTopic(smnClient *client.SmnClient) {
	request := smnClient.NewDeleteTopicRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"

	response, err := smnClient.DeleteTopic(request)
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

func ListTopic(smnClient *client.SmnClient) {
	request := smnClient.NewListTopicRequest()
	request.Limit = 10
	request.Offset = 0
	response, err := smnClient.ListTopic(request)
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

func QueryTopicDetail(smnClient *client.SmnClient) {
	request := smnClient.NewQueryTopicDetailRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"

	response, err := smnClient.QueryTopicDetail(request)
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

func UpdateTopicAttribute(smnClient *client.SmnClient) {
	request := smnClient.NewUpdateTopicAttributeRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	request.Name = "access_policy"
	request.Value = `{
         "Version": "2016-09-07",
         "Id": "__default_policy_ID",
         "Statement": [
            {
              "Sid": "__user_pub_0",
              "Effect": "Allow",
              "Principal": {
                "CSP": [
                         "urn:csp:iam::1040774eae344b78b14f2939863d4ede:root"
                       ]
                 },
              "Action": ["SMN:Publish","SMN:QueryTopicDetail"],
              "Resource": "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
              }
             ]
          }`
	response, err := smnClient.UpdateTopicAttribute(request)
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

func ListTopicAttributes(smnClient *client.SmnClient) {
	request := smnClient.NewListTopicAttributesRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	request.Name = "access_policy"
	response, err := smnClient.ListTopicAttributes(request)
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

func DeleteTopicAttributeByName(smnClient *client.SmnClient) {
	request := smnClient.NewDeleteTopicAttributeByNameRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	request.Name = "access_policy"
	response, err := smnClient.DeleteTopicAttributeByName(request)
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

func DeleteTopicAttributes(smnClient *client.SmnClient) {
	request := smnClient.NewDeleteTopicAttributesRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	response, err := smnClient.DeleteTopicAttributes(request)
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
