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
	"io"
	"github.com/smn-sdk-go/smn-sdk-go/util"
)

// the base request info
type BaseRequest struct {
	projectId  string `json:"-"`
	url        string `json:"-"`
	method     string `json:"-"`
	regionName string `json:"-"`

	Headers map[string]string `json:"-"`
}

// the interface for request
type SmnRequest interface {
	GetUrl() (string, error)
	GetMethod() string
	GetHeaders() map[string]string
	GetBodyReader() (io.Reader, error)
	SetProjectId(projectId string)
	SetRegionName(regionName string)
	addHeaderParam(key, value string)
}

// get the request headers
func (request *BaseRequest) GetHeaders() map[string]string {
	return request.Headers
}

// add the header to request
func (request *BaseRequest) addHeaderParam(key, value string) {
	request.Headers[key] = value
}

// set project
func (request *BaseRequest) SetProjectId(projectId string) {
	request.projectId = projectId
}

// set region name
func (request *BaseRequest) SetRegionName(regionName string) {
	request.regionName = regionName
}

// get smn service url
func (request *BaseRequest) GetSmnServiceUrl() string {
	return util.HttpsPrefix + util.Smn + "." + request.regionName + "." + util.Endpoint +
		util.UrlDelimiter
}
