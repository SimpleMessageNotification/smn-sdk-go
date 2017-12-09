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
	"fmt"
)

type DeleteSmsSignRequest struct {
	*BaseRequest
	SignId string `json:"sign_id"`
}

type DeleteSmsSignResponse struct {
	*BaseResponse
}

func (client *SmnClient) DeleteSmsSign(request *DeleteSmsSignRequest) (response *DeleteSmsSignResponse, err error) {
	response = &DeleteSmsSignResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewDeleteSmsSignRequest() (request *DeleteSmsSignRequest) {
	request = &DeleteSmsSignRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *DeleteSmsSignRequest) GetUrl() (string, error) {
	if request.SignId == "" {
		return "", fmt.Errorf("signId is null")
	}
	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmsSignature + util.UrlDelimiter +
		request.SignId, nil
}

func (request *DeleteSmsSignRequest) GetMethod() string {
	return util.DELETE
}

func (request *DeleteSmsSignRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
