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
package client

import (
	"fmt"
	"io"
	"github.com/SimpleMessageNotification/smn-sdk-go/smn-sdk-go/util"
)

// the request data of publish sms with diff message
type SmsBatchPublishWithDiffMessageRequest struct {
	*BaseRequest
	SmsMessage []SmsPublishMessage `json:"sms_message"`
}

type SmsPublishMessage struct {
	Message                string `json:"message""`
	EndPoint               string `json:"endpoint"`
	SignId                 string `json:"sign_id"`
	MessageIncludeSignFlag bool   `json:"message_include_sign_flag"`
}

//the response data of batch publish sms with diff message
type SmsBatchPublishWithDiffMessageResponse struct {
	*BaseResponse
	Result []SmsBatchPublishWithDiffMessageResult `json:"result"`
}

type SmsBatchPublishWithDiffMessageResult struct {
	MessageId string `json:"message_id"`
	endpoint  string `json:"endpoint"`
	code      string `json:"code"`
	message   string `json:"message"`
}

// send request to batch publish sms with diff message
func (client *SmnClient) SmsBatchPublishWithDiffMessage(request *SmsBatchPublishWithDiffMessageRequest) (response *SmsBatchPublishWithDiffMessageResponse, err error) {
	response = &SmsBatchPublishWithDiffMessageResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list sms batch publish with diff message request struct
func (client *SmnClient) NewSmsBatchPublishWithDiffMessageRequest() (request *SmsBatchPublishWithDiffMessageRequest) {
	request = &SmsBatchPublishWithDiffMessageRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the list sms batch publish with diff message request
func (request *SmsBatchPublishWithDiffMessageRequest) GetUrl() (string, error) {
	if request.SmsMessage == nil || len(request.SmsMessage) == 0 {
		return "", fmt.Errorf("smsMessage is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnProtocolSmsBatch, nil
}

// get the http method of the sms batch publish with diff message request
func (request *SmsBatchPublishWithDiffMessageRequest) GetMethod() string {
	return util.POST
}

// get the body params of the sms batch publish with diff message request
func (request *SmsBatchPublishWithDiffMessageRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
