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

	//list subscriptions
	ListSubscriptions(&tempClient)
	//list subscriptions by topic
	ListSubscriptionsByTopic(&tempClient)
	// un subscribe
	Unsubscribe(&tempClient)
	//Subscribe
	Subscribe(&tempClient)
}

func ListSubscriptions(smnClient *client.SmnClient) {
	request := smnClient.NewListSubscriptionsRequest()
	request.Limit = 10
	request.Offset = 0
	response, err := smnClient.ListSubscriptions(request)
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

func ListSubscriptionsByTopic(smnClient *client.SmnClient) {
	request := smnClient.NewListSubscriptionsByTopicRequest()
	request.Limit = 10
	request.Offset = 0
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	response, err := smnClient.ListSubscriptionsByTopic(request)
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

func Subscribe(smnClient *client.SmnClient) {
	request := smnClient.NewSubscribeRequest()
	request.Endpoint = "375831500@qq.com"
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	request.Protocol = "email"
	request.Remark = "email"
	response, err := smnClient.Subscribe(request)
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

func Unsubscribe(smnClient *client.SmnClient) {
	request := smnClient.NewUnsubscribeRequest()
	request.SubscriptionUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go:9e52cb38e107402d83cae90fa8b9f6ed"
	response, err := smnClient.Unsubscribe(request)
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
