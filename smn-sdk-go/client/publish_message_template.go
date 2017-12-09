package client

import (
	"fmt"
	"smn-sdk-go/util"
	"io"
)

type PublishMessageTemplateRequest struct {
	*BaseRequest
	TopicUrn            string            `json:"-"`
	Subject             string            `json:"subject"`
	MessageTemplateName string            `json:"message_template_name"`
	Tags                map[string]string `json:"tags"`
}
type PublishMessageTemplateResponse struct {
	*BaseResponse
	MessageId string `json:"message_id"`
}

func (client *SmnClient) PublishMessageTemplate(request *PublishMessageTemplateRequest) (response *PublishMessageTemplateResponse, err error) {
	response = &PublishMessageTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewPublishMessageTemplateRequest() (request *PublishMessageTemplateRequest) {
	request = &PublishMessageTemplateRequest{
		Tags:        make(map[string]string),
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *PublishMessageTemplateRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	if !util.ValidateSubject(request.Subject) {
		return "", fmt.Errorf("subject is invalid or oversize 512bytes")
	}

	if request.MessageTemplateName == "" {
		return "", fmt.Errorf("message template name is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter + util.Publish, nil
}

func (request *PublishMessageTemplateRequest) GetMethod() string {
	return util.POST
}

func (request *PublishMessageTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
