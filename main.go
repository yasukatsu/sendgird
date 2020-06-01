package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	fmt.Printf("%v\n", "start")
	r := mux.NewRouter()
	r.HandleFunc("/helper", helperPost)
	r.HandleFunc("/notHelper", notHelperPost)
	// DynamicTemplateを活用してのメール送信
	r.HandleFunc("/dynamicTemplate", sendDynamicTemplateEmail)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))

}

var myEmail = os.Getenv("EMAIL")

func helperPost(w http.ResponseWriter, r *http.Request) {

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
	byteResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	w.Write(byteResponse)
}

func sendDynamicTemplateEmail(w http.ResponseWriter, r *http.Request) {
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
	byteResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	w.Write(byteResponse)
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

	imgURL := "https://fittio.jp/wp-content/themes/fittio.jp/common/img/fittio_logoB_4c.svg"
	p.SetDynamicTemplateData("img", imgURL)

	p.SetDynamicTemplateData("subject", subject)
	p.SetDynamicTemplateData("NAME_KANA", "田中太郎")
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

func notHelperPost(w http.ResponseWriter, r *http.Request) {

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
	byteResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	w.Write(byteResponse)
}
