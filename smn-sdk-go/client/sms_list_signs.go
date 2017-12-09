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
	"smn-sdk-go/util"
	"io"
)

type ListSmsSignsRequest struct {
	*BaseRequest
}

type ListSmsSignsResponse struct {
	*BaseResponse
	SmsSignCount int           `json:"sms_sign_count"`
	SmsSigns     []SmsSignInfo `json:"sms_signs"`
}

type SmsSignInfo struct {
	SignName   string `json:"sign_name"`
	CreateTime string `json:"create_time"`
	SignId     string `json:"sign_id"`
	Reply      string `json:"reply"`
	Status     int    `json:"status"`
}

func (client *SmnClient) ListSmsSigns(request *ListSmsSignsRequest) (response *ListSmsSignsResponse, err error) {
	response = &ListSmsSignsResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewListSmsSignsRequest() (request *ListSmsSignsRequest) {
	request = &ListSmsSignsRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *ListSmsSignsRequest) GetUrl() (string, error) {
	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmsSignature, nil
}

func (request *ListSmsSignsRequest) GetMethod() string {
	return util.GET
}

func (request *ListSmsSignsRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
