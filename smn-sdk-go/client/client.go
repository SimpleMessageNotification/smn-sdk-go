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
	"net/http"
	"smn-sdk-go/commom"
	"smn-sdk-go/auth"
	"fmt"
	"encoding/json"
)

var version = "0.0.1"

type SmnClient struct {
	httpClient       *http.Client
	smnConfiguration *commom.SmnConfiguration
	auth             *auth.Auth
}

func NewClient(userName, domainName, password, regionName string) (client SmnClient, err error) {
	smnConfiguration := &commom.SmnConfiguration{Username: userName, DomainName: domainName, Password: password, RegionName: regionName}
	client.smnConfiguration = smnConfiguration
	httpClient := &http.Client{}
	client.httpClient = httpClient
	client.auth = auth.NewAuth(smnConfiguration, httpClient)
	return
}

func (client *SmnClient) sendRequest(request SmnRequest, response SmnResponse) (err error) {
	if err := client.buildHeaderAndParam(request); err != nil {
		return err
	}

	requestMethod := request.GetMethod()
	requestUrl, err := request.GetUrl()
	if err != nil {
		return
	}

	body, err := request.GetBodyReader()
	if err != nil {
		return
	}
	httpRequest, err := http.NewRequest(requestMethod, requestUrl, body)
	if err != nil {
		return
	}
	for key, value := range request.GetHeaders() {
		httpRequest.Header[key] = []string{value}
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return
	}
	err = unmarshal(response, httpResponse)
	return
}

func unmarshal(response SmnResponse, httpResponse *http.Response) (err error) {
	err = response.parseHttpResponse(httpResponse)
	if err != nil {
		return
	}
	err = json.Unmarshal(response.GetContentBytes(), response)
	return
}

func (client *SmnClient) buildHeaderAndParam(request SmnRequest) (err error) {
	token, projectId, err := client.auth.GetTokenAndProject()
	if err != nil {
		return
	}
	if projectId == "" || token == "" {
		return fmt.Errorf("token or projectId is invalid")
	}
	request.SetProjectId(projectId)
	request.SetRegionName(client.smnConfiguration.RegionName)
	request.addHeaderParam("Content-Type", "application/json; charset=UTF-8")
	request.addHeaderParam("region", client.smnConfiguration.RegionName)
	request.addHeaderParam("X-Auth-Token", token)
	request.addHeaderParam("X-Project-Id", projectId)
	request.addHeaderParam("User-Agent", "smn-sdk-go/"+version)
	return
}
