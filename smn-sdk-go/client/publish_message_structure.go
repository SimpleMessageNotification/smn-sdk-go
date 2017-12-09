package client

import (
	"fmt"
	"smn-sdk-go/util"
	"io"
)

type PublishMessageStructureRequest struct {
	*BaseRequest
	TopicUrn         string            `json:"-"`
	Subject          string            `json:"subject"`
	MessageStructure string `json:"message_structure"`
}
type PublishMessageStructureResponse struct {
	*BaseResponse
	MessageId string `json:"message_id"`
}

func (client *SmnClient) PublishMessageStructure(request *PublishMessageStructureRequest) (response *PublishMessageStructureResponse, err error) {
	response = &PublishMessageStructureResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewPublishMessageStructureRequest() (request *PublishMessageStructureRequest) {
	request = &PublishMessageStructureRequest{
		BaseRequest:      &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *PublishMessageStructureRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	if !util.ValidateSubject(request.Subject) {
		return "", fmt.Errorf("subject is invalid or oversize 512bytes")
	}

	if len(request.MessageStructure) == 0 {
		return "", fmt.Errorf("message structure is null")
	}

	if !util.ValidateMessageStructureDefault(request.MessageStructure) {
		return "", fmt.Errorf("message structure not contain defalult message")

	}

	if !util.ValidateMessageStructureLength(request.MessageStructure) {
		return "", fmt.Errorf("message structure is oversize")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter + util.Publish, nil
}

func (request *PublishMessageStructureRequest) GetMethod() string {
	return util.POST
}

func (request *PublishMessageStructureRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
