package client

import (
	"smn-sdk-go/util"
	"io"
	"fmt"
)

type CreateTopicRequest struct {
	*BaseRequest
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type CreateTopicResponse struct {
	*BaseResponse
	TopicUrn string `json:"topic_urn"`
}

func (client *SmnClient) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	response = &CreateTopicResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewCreateTopicRequest() (request *CreateTopicRequest) {
	request = &CreateTopicRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *CreateTopicRequest) GetUrl() (string, error) {
	if !util.ValidateTopicName(request.Name) {
		return "", fmt.Errorf("topic name is invalid")
	}
	if !util.ValidateTopicDiplayName(request.DisplayName) {
		return "", fmt.Errorf("topic displayName is overSize")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics, nil
}

func (request *CreateTopicRequest) GetMethod() string {
	return util.POST
}

func (request *CreateTopicRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
