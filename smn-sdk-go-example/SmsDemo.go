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

	//// send notify/verify sms
	//SmsPublish(&tempClient)
	// batch send notify/verify sms
	SmsBatchPublish(&tempClient)
	// batch send notify/verify sms with diff message
	SmsBatchPublishWithDiffMessage(&tempClient)
	//// send promotion sms
	//PromotionSmsPublish(&tempClient)
	//// create sms template
	//CreateSmsTemplate(&tempClient)
	////delete sms template
	//DeleteSmsTemplate(&tempClient)
	//// list sms template
	//ListSmsTemplate(&tempClient)
	//// query sms template detail
	//GetSmsTemplateDetail(&tempClient)
	//// list sms signs
	//ListSmsSigns(&tempClient)
	//// delete sms sign
	//DeleteSmsSign(&tempClient)
	//// list sms msg report
	//ListSmsMsgReport(&tempClient)
	//// get sms message content
	//GetSmsMessage(&tempClient)
	//// list sms callback event
	//ListSmsEvent(&tempClient)
	//// update sms callback event
	//UpdateSmsEvent(&tempClient)
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

// batch send notify/verify sms
func SmsBatchPublish(smnClient *client.SmnClient) {
	request := smnClient.NewSmsBatchPublishRequest()
	request.EndPoints = []string{"8613688807587"}
	request.SignId = "6be340e91e5241e4b5d85837e6709104"
	request.Message = "您的验证码是:123455，请查收"
	response, err := smnClient.SmsBatchPublish(request)

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

// batch send notify/verify sms with diff message
func SmsBatchPublishWithDiffMessage(smnClient *client.SmnClient) {
	request := smnClient.NewSmsBatchPublishWithDiffMessageRequest()
	smsMessage1 := client.SmsPublishMessage{Message: "验证码12355测试", EndPoint: "13688807587", SignId: "6be340e91e5241e4b5d85837e6709104"}
	smsMessage2 := client.SmsPublishMessage{Message: "验证码12355测试", EndPoint: "17727904831", SignId: "6be340e91e5241e4b5d85837e6709104"}
	request.SmsMessage = append(request.SmsMessage, smsMessage1, smsMessage2)
	response, err := smnClient.SmsBatchPublishWithDiffMessage(request)

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

func PromotionSmsPublish(smnClient *client.SmnClient) {
	request := smnClient.NewPromotionSmsPublishRequest()
	request.SignId = "47f86cf7c9a7449d98ee61cf193a1060"
	request.SmsTemplateId = "bfda25c6406e42ddabad74b4a20f6d05"
	request.EndPoints = []string{"8613688807587"}

	response, err := smnClient.PromotionSmsPublish(request)
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

func CreateSmsTemplate(smnClient *client.SmnClient) {
	request := smnClient.NewCreateSmsTemplateRequest()
	request.SmsTemplateType = 1
	request.SmsTemplateName = "go语言测试"
	request.SmsTemplateContent = "go语言测试12341234"
	request.Remark = "helloword"
	response, err := smnClient.CreateSmsTemplate(request)
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

func GetSmsTemplateDetail(smnClient *client.SmnClient) {
	request := smnClient.NewGetSmsTemplateDetailRequest()
	request.SmsTemplateId = "bfda25c6406e42ddabad74b4a20f6d05"

	response, err := smnClient.GetSmsTemplateDetailRequest(request)
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

func DeleteSmsTemplate(smnClient *client.SmnClient) {
	request := smnClient.NewDeleteSmsTemplateRequest()
	request.SmsTemplateId = "5341edf7f0aa410684737511b26d72b9"

	response, err := smnClient.DeleteSmsTemplate(request)
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

func ListSmsTemplate(smnClient *client.SmnClient) {
	request := smnClient.NewListSmsTemplatesRequest()
	request.Offset = "0"
	request.Limit = "10"
	request.SmsTemplateType = 1
	request.SmsTemplateName = "模板"
	response, err := smnClient.ListSmsTemplates(request)
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
