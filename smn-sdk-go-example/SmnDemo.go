package main

import (
	"smn-sdk-go/client"
	"fmt"
)

var smnClient *client.SmnClient = nil

func main() {
	// init client
	tempClient, err := client.NewClient("youUserName", "youDomainName", "youPassword", "cn-north-1")
	if err != nil {
		panic(err)
	}

	smnClient = &tempClient
	// send sms
	SmsPublish()
	// list sms signs
	ListSmsSigns()
	// delete sms sign
	DeleteSmsSign()
	// list sms msg report
	ListSmsMsgReport()
	// get sms message content
	GetSmsMessage()
	// list sms callback event
	ListSmsEvent()
	// update sms callback event
	UpdateSmsEvent()

	// create topic
	CreateTopic()
	// update topic
	UpdateTopic()
	// delete topic
	DeleteTopic()
	// list topics
	ListTopic()
	// query topic detail
	QueryTopicDetail()
	// update topic attribute
	UpdateTopicAttribute()
	// list topic attributes
	ListTopicAttributes()
	// delete topic attribute by name
	DeleteTopicAttributeByName()
	// delete all topic attributes
	DeleteTopicAttributes()

	//publish message
	PublishMessage()
	// publish message structure
	PublishMessageStructure()
	// publish message template
	PublishMessageTemplate()
	// list sms signs
	ListSmsSigns()
	// delete sms sign
	DeleteSmsSign()
	// list sms msg report
	ListSmsMsgReport()
	// get sms message content
	GetSmsMessage()
	// list sms callback event
	ListSmsEvent()
	// update sms callback event
	UpdateSmsEvent()

	//list subscriptions
	ListSubscriptions()
	//list subscriptions by topic
	ListSubscriptionsByTopic()
	// un subscribe
	Unsubscribe()
	//Subscribe
	Subscribe()

	// create message template
	CreateMessageTemplate()
	// query message template detail
	QueryMessageTemplateDetail()
	// list message template
	ListMessageTemplate()
	// update message template
	UpdateMessageTemplate()
	// delete message template
	DeleteMessageTemplate()
}

// send sms
func SmsPublish() {
	request := smnClient.NewSmsPublishRequest()
	request.EndPoint = "+8613688807587"
	request.SignId = "6be340e91e5241e4b5d85837e6709104"
	request.Message = "您的验证码是:1234，请查收"
	response, err := smnClient.SmsPublish(request)

	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

// list sms signs
func ListSmsSigns() {
	request := smnClient.NewListSmsSignsRequest()
	response, err := smnClient.ListSmsSigns(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

// delete sms sign
func DeleteSmsSign() {
	request := smnClient.NewDeleteSmsSignRequest()
	request.SignId = "d886b2a9a2cf44499e9c5e6f283b3ae9"
	response, err := smnClient.DeleteSmsSign(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)

}

func ListSmsMsgReport() {
	request := smnClient.NewListSmsMsgReportRequest()
	request.StartTime = "1512625955366"
	request.EndTime = "1512712355850"
	request.SignId = "6be340e91e5241e4b5d85837e6709104"
	response, err := smnClient.ListSmsMsgReport(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func GetSmsMessage() {
	request := smnClient.NewGetSmsMessageRequest()

	request.MessageId = "8fab2e49115642e8bcf26bdc58ab297d"
	response, err := smnClient.GetSmsMessage(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func ListSmsEvent() {
	request := smnClient.NewListSmsEventRequest()
	request.EventType = "sms_reply_event"
	response, err := smnClient.ListSmsEvent(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func UpdateSmsEvent() {
	request := smnClient.NewUpdateSmsEventRequest()
	callback1 := client.CallbackInfo{EventType: "sms_reply_event", TopicUrn: "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_liuqiangqiang_send_success"}
	request.Callback = append(request.Callback, callback1)

	response, err := smnClient.UpdateSmsEvent(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func CreateTopic() {
	request := smnClient.NewCreateTopicRequest()
	request.Name = "zhangyxTestCreate01"
	request.DisplayName = "zhangyxTestCreate01"

	response, err := smnClient.CreateTopic(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func UpdateTopic() {
	request := smnClient.NewUpdateTopicRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	request.DisplayName = "zhangyxTestCreate001"

	response, err := smnClient.UpdateTopic(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func DeleteTopic() {
	request := smnClient.NewDeleteTopicRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"

	response, err := smnClient.DeleteTopic(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func ListTopic() {
	request := smnClient.NewListTopicRequest()
	request.Limit = "10"
	request.Offset = "0"
	response, err := smnClient.ListTopic(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func QueryTopicDetail() {
	request := smnClient.NewQueryTopicDetailRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"

	response, err := smnClient.QueryTopicDetail(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func UpdateTopicAttribute() {
	request := smnClient.NewUpdateTopicAttributeRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	request.Name = "access_policy"
	request.Value = `{
         "Version": "2016-09-07",
         "Id": "__default_policy_ID",
         "Statement": [
            {
              "Sid": "__user_pub_0",
              "Effect": "Allow",
              "Principal": {
                "CSP": [
                         "urn:csp:iam::1040774eae344b78b14f2939863d4ede:root"
                       ]
                 },
              "Action": ["SMN:Publish","SMN:QueryTopicDetail"],
              "Resource": "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
              }
             ]
          }`
	response, err := smnClient.UpdateTopicAttribute(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func ListTopicAttributes() {
	request := smnClient.NewListTopicAttributesRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	request.Name = "access_policy"
	response, err := smnClient.ListTopicAttributes(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func DeleteTopicAttributeByName() {
	request := smnClient.NewDeleteTopicAttributeByNameRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	request.Name = "access_policy"
	response, err := smnClient.DeleteTopicAttributeByName(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func DeleteTopicAttributes() {
	request := smnClient.NewDeleteTopicAttributesRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:zhangyxTestCreate01"
	response, err := smnClient.DeleteTopicAttributes(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func PublishMessage() {
	request := smnClient.NewPublishMessageRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	request.Message = "test by zhangyx"
	response, err := smnClient.PublishMessage(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func PublishMessageStructure() {
	request := smnClient.NewPublishMessageStructureRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	request.MessageStructure = "{" +
		"\"default\":\"test by zhangyx structure\"," +
		"\"email\":\"test by zhangyx structure email\"," +
		"\"sms\":\"test by zhangyx structure _sms\"}"
	response, err := smnClient.PublishMessageStructure(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func PublishMessageTemplate() {
	request := smnClient.NewPublishMessageTemplateRequest()
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	request.MessageTemplateName = "createMessageTemplate"
	request.Tags["year"] = "2016"
	request.Tags["company"] = "hellokitty"
	response, err := smnClient.PublishMessageTemplate(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}


func ListSubscriptions() {
	request := smnClient.NewListSubscriptionsRequest()
	request.Limit = "10"
	request.Offset = "0"
	response, err := smnClient.ListSubscriptions(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func ListSubscriptionsByTopic() {
	request := smnClient.NewListSubscriptionsByTopicRequest()
	request.Limit = "10"
	request.Offset = "0"
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	response, err := smnClient.ListSubscriptionsByTopic(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func Subscribe() {
	request := smnClient.NewSubscribeRequest()
	request.Endpoint = "375831500@qq.com"
	request.TopicUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go"
	request.Protocol = "email"
	request.Remark = "email"
	response, err := smnClient.Subscribe(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func Unsubscribe() {
	request := smnClient.NewUnsubscribeRequest()
	request.SubscriptionUrn = "urn:smn:cn-north-1:cffe4fc4c9a54219b60dbaf7b586e132:test_zhangyx_go:9e52cb38e107402d83cae90fa8b9f6ed"
	response, err := smnClient.Unsubscribe(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}


func CreateMessageTemplate() {
	request := smnClient.NewCreateMessageTemplateRequest()
	request.Protocol = "sms"
	request.MessageTemplateName = "template_test_by_zhangyx_for_go"
	request.Content = "this is the first template for go {time}"
	response, err := smnClient.CreateMessageTemplate(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func QueryMessageTemplateDetail() {
	request := smnClient.NewQueryMessageTemplateDetailRequest()
	request.MessageTemplateId = "d4c4968551534a298233c92b6c4916c7"
	response, err := smnClient.QueryMessageTemplateDetail(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func ListMessageTemplate() {
	request := smnClient.NewListMessageTemplateRequest()
	request.Offset = "0"
	request.Limit = "10"
	request.MessageTemplateName = "template_test_by_zhangyx_for_go"
	response, err := smnClient.ListMessageTemplate(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func UpdateMessageTemplate() {
	request := smnClient.NewUpdateMessageTemplateRequest()
	request.MessageTemplateId = "d4c4968551534a298233c92b6c4916c7"
	request.Content = "this is the first template for go {time} v2"
	response, err := smnClient.UpdateMessageTemplate(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}

func DeleteMessageTemplate() {
	request := smnClient.NewDeleteMessageTemplateRequest()
	request.MessageTemplateId = "d4c4968551534a298233c92b6c4916c7"
	response, err := smnClient.DeleteMessageTemplate(request)
	if err != nil {
		fmt.Println("the request is error ", err)
		return
	}

	if !response.IsSuccess() {
		fmt.Printf("%#v\n", response.ErrorResponse)
		return
	}

	fmt.Printf("%#v\n", response)
}
