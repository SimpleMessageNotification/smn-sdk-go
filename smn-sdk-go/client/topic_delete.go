package client

import (
	"smn-sdk-go/util"
	"io"
	"fmt"
)

type DeleteTopicRequest struct {
	*BaseRequest
	TopicUrn string
}

type DeleteTopicResponse struct {
	*BaseResponse
}

func (client *SmnClient) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
	response = &DeleteTopicResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewDeleteTopicRequest() (request *DeleteTopicRequest) {
	request = &DeleteTopicRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *DeleteTopicRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn, nil
}

func (request *DeleteTopicRequest) GetMethod() string {
	return util.DELETE
}

func (request *DeleteTopicRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
