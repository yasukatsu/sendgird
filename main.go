package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {

	// helperPost()
	// notHelperPost()

	// DynamicTemplateを活用してのメール送信
	sendDynamicTemplateEmail()

}

var myEmail = os.Getenv("EMAIL")

func helperPost() {

	var (
		SubjectTopStr            = "送信したユーザーより"
		OfctourOrWktrialAppendix = "体験"
		SubjectPremiumStr        = ""
		EntryMgrNo               = 12345
	)

	from := mail.NewEmail("Example User", "test@example.com")
	subject := fmt.Sprintf("%s%s応募がありました%s。[問い合わせNo：%d]", SubjectTopStr, OfctourOrWktrialAppendix, SubjectPremiumStr, EntryMgrNo)
	to := mail.NewEmail("Example User", myEmail)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<h1>and easy to do anywhere, even with Go</h1>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func sendDynamicTemplateEmail() {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = dynamicTemplateEmail()
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func dynamicTemplateEmail() []byte {
	m := mail.NewV3Mail()

	// from
	address := "test@example.com"
	name := "Example User"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	tempateID := os.Getenv("TEMPLATE_ID_1")
	m.SetTemplateID(tempateID)

	var (
		SubjectTopStr            = "送信したユーザーより"
		OfctourOrWktrialAppendix = "体験"
		SubjectPremiumStr        = ""
		EntryMgrNo               = 12345
	)
	subject := fmt.Sprintf("%s%s応募がありました%s。[問い合わせNo：%d]", SubjectTopStr, OfctourOrWktrialAppendix, SubjectPremiumStr, EntryMgrNo)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("Example User", myEmail),
	}
	p.AddTos(tos...)

	imgURL := "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.rbbtoday.com%2Farticle%2F2016%2F03%2F04%2F140287.html&psig=AOvVaw10U9FxxRFUPvABLOFNMbw9&ust=1590392902820000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCJChtrWBzOkCFQAAAAAdAAAAABAD"
	p.SetDynamicTemplateData("img", imgURL)

	p.SetDynamicTemplateData("subject", subject)
	p.SetDynamicTemplateData("NAME_KANA", "氏名（ふりがな）")
	p.SetDynamicTemplateData("AGE", "年齢")
	p.SetDynamicTemplateData("SEX", "性別")
	p.SetDynamicTemplateData("TEL1", "電話番号1")
	p.SetDynamicTemplateData("TEL2", "電話番号2")
	p.SetDynamicTemplateData("EMAIL1", "メールアドレス1")
	p.SetDynamicTemplateData("EMAIL2", "メールアドレス2")
	p.SetDynamicTemplateData("ADDRESS", "住所")

	m.AddPersonalizations(p)
	return mail.GetRequestBody(m)
}

func notHelperPost() {

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = []byte(` {
	"personalizations": [
		{
			"to": [
				{
					"email": {{myaddress}}
				}
			],
			"subject": "' + FakePass + '{$SUBJECT_TOP_STR}{$OFCTOUR_OR_WKTRIAL_APPENDIX}応募がありました{$SUBJECT_PREMIUM_STR}。[問い合わせNo：{$EntryMgrNo}]"
		}
	],
	"from": {
		"email": "test@example.com"
	},
	"content": [
		{
			"type": "text/html",
			"value": |
			"<h1>Hello World</h1><h2>hello world</h2>"
		}
	]
	}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
