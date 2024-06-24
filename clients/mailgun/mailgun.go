package mailgun

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"api.default.marincor.pt/adapters/logging"
	"api.default.marincor.pt/app/appinstance"
	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	"api.default.marincor.pt/pkg/helpers"
)

type Mailgun struct {
	APIHost string
	APIKey  string
}

type mailgunResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func New() *Mailgun {
	return &Mailgun{
		APIHost: fmt.Sprintf("https://api.mailgun.net/v3/%s/messages", appinstance.Data.Config.MailGunDomain),
		APIKey:  fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("api:"+appinstance.Data.Config.MailGunKey))),
	}
}

func parseTemplate(templateName string, args ...map[string]interface{}) (string, error) {
	var (
		err     error
		tmpl    *template.Template
		content bytes.Buffer
	)

	tmpl, err = template.ParseFiles(fmt.Sprintf("%s/%s", constants.TemplatesFolder, templateName))
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message: "error to get email template",
			Request: templateName,
			Reason:  err.Error(),
		}, "critical", nil)

		return "", err
	}

	if len(args) > 0 {
		err = tmpl.Execute(&content, args[0])
		if err != nil {
			logging.Log(&entity.LogDetails{
				Message: "error to bind args in email template",
				Reason:  err.Error(),
				Request: args,
			}, "critical", nil)

			return "", err
		}
	}

	return content.String(), nil
}

func (mailgun *Mailgun) Send(to string, messageAttr *entity.MessageAttributes) { //nolint:varnamelen
	content, err := parseTemplate(messageAttr.Template, messageAttr.Args)
	if err != nil {
		return
	}

	emailData := map[string]string{
		"from":    fmt.Sprintf("%s <%s>", appinstance.Data.Config.EmailSenderLabel, appinstance.Data.Config.EmailSenderAddress),
		"to":      to,
		"subject": messageAttr.Subject,
		"html":    content,
	}

	form, contentType, err := helpers.WriteFormData(emailData)
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "error to parse mailgun email body",
			Reason:  err.Error(),
			Request: fmt.Sprintf("%+v", form),
		}, "critical", nil)

		return
	}

	TODO: MAKE REQUEST HTTP
	// request := requestsnippet.Request{
	// 	Method: http.MethodPost,
	// 	URI:    mailgun.APIHost,
	// 	Headers: []requestsnippet.Header{
	// 		{Key: "Authorization", Value: mailgun.APIKey},
	// 		{Key: "Content-Type", Value: contentType},
	// 	},
	// 	Body: form,
	// }

	response, err := request.Call()
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message:  "error to send email through mailgun",
			Reason:   err.Error(),
			Request:  request,
			Response: response,
		}, "critical", nil)

		return
	}

	if !helpers.Contains(constants.HTTPStatusesOk, fmt.Sprintf("%d", response.StatusCode)) {
		logging.Log(&entity.LogDetails{
			Message:    "status code error while sending email through mailgun",
			Request:    request,
			Response:   string(response.Message),
			StatusCode: response.StatusCode,
		}, "critical", nil)

		return
	}

	var responseMessage mailgunResponse
	err = json.Unmarshal(response.Message, &responseMessage)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message:  "error to unmarshal mailgun response",
			Reason:   err.Error(),
			Request:  request,
			Response: string(response.Message),
		}, "critical", nil)

		return
	}

	logging.Log(&entity.LogDetails{
		Message:  "email successfully sent through mailgun",
		Request:  request,
		Response: responseMessage,
	}, "info", nil)
}
