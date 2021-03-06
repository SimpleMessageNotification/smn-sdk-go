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
	"github.com/SimpleMessageNotification/smn-sdk-go/smn-sdk-go/util"
	"io"
	"fmt"
)

// the request data of update sms event
type UpdateSmsEventRequest struct {
	*BaseRequest
	Callback []CallbackInfo `json:"callback"`
}

// the response data of update sms event
type UpdateSmsEventResponse struct {
	*BaseResponse
}

// send request to update sms event
func (client *SmnClient) UpdateSmsEvent(request *UpdateSmsEventRequest) (response *UpdateSmsEventResponse, err error) {
	response = &UpdateSmsEventResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new update sms event request struct
func (client *SmnClient) NewUpdateSmsEventRequest() (request *UpdateSmsEventRequest) {
	request = &UpdateSmsEventRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the update sms event request
func (request *UpdateSmsEventRequest) GetUrl() (urlStr string, err error) {
	if request.Callback == nil || len(request.Callback) == 0 {
		return "", fmt.Errorf("callbacks is invalid")
	}

	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnProtocolSms +
		util.UrlDelimiter + util.CallBack
	return
}

// get the http method of the update sms event request
func (request *UpdateSmsEventRequest) GetMethod() string {
	return util.PUT
}

// get the body params of the update sms event request
func (request *UpdateSmsEventRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
