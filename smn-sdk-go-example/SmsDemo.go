/*
 * Copyright (C) 2017. Huawei Technologies Co., LTD. All rights reserved.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of Apache License, Version 2.0.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * Apache License, Version 2.0 for more details.
 */
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

	// send sms
	SmsPublish(&tempClient)
	// list sms signs
	ListSmsSigns(&tempClient)
	// delete sms sign
	DeleteSmsSign(&tempClient)
	// list sms msg report
	ListSmsMsgReport(&tempClient)
	// get sms message content
	GetSmsMessage(&tempClient)
	// list sms callback event
	ListSmsEvent(&tempClient)
	// update sms callback event
	UpdateSmsEvent(&tempClient)
}

// send sms
func SmsPublish(smnClient *client.SmnClient) {
	request := smnClient.NewSmsPublishRequest()
	request.EndPoint = "+8613688807587"
	request.SignId = "6be340e91e5241e4b5d85837e6709104"
	request.Message = "您的验证码是:1234，请查收"
	response, err := smnClient.SmsPublish(request)

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

// list sms signs
func ListSmsSigns(smnClient *client.SmnClient) {
	request := smnClient.NewListSmsSignsRequest()
	response, err := smnClient.ListSmsSigns(request)
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

// delete sms sign
func DeleteSmsSign(smnClient *client.SmnClient) {
	request := smnClient.NewDeleteSmsSignRequest()
	request.SignId = "d886b2a9a2cf44499e9c5e6f283b3ae9"
	response, err := smnClient.DeleteSmsSign(request)
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

func ListSmsMsgReport(smnClient *client.SmnClient) {
	request := smnClient.NewListSmsMsgReportRequest()
	request.StartTime = "1512625955366"
	request.EndTime = "1512712355850"
	request.SignId = "6be340e91e5241e4b5d85837e6709104"
	request.Limit = 10
	request.Offset = 0
	response, err := smnClient.ListSmsMsgReport(request)
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

func GetSmsMessage(smnClient *client.SmnClient) {
	request := smnClient.NewGetSmsMessageRequest()

	request.MessageId = "8fab2e49115642e8bcf26bdc58ab297d"
	response, err := smnClient.GetSmsMessage(request)
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

func ListSmsEvent(smnClient *client.SmnClient) {
	request := smnClient.NewListSmsEventRequest()
	request.EventType = "sms_reply_event"
	response, err := smnClient.ListSmsEvent(request)
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

func UpdateSmsEvent(smnClient *client.SmnClient) {
	request := smnClient.NewUpdateSmsEventRequest()
	callback1 := client.CallbackInfo{EventType: "sms_reply_event", TopicUrn: "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_liuqiangqiang_send_success"}
	request.Callback = append(request.Callback, callback1)

	response, err := smnClient.UpdateSmsEvent(request)
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