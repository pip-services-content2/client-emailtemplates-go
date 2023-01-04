package version1

import "github.com/pip-services3-gox/pip-services3-commons-gox/data"

type EmailTemplateV1 struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	From    string            `json:"from"`
	ReplyTo string            `json:"reply_to"`
	Subject map[string]string `json:"subject"`
	Text    map[string]string `json:"text"`
	Html    map[string]string `json:"html"`
	Status  string            `json:"status"`
}

func NewEmailTemplateV1(name string, subject, text, html map[string]string, status string) *EmailTemplateV1 {
	c := &EmailTemplateV1{
		Id:      data.IdGenerator.NextLong(),
		Name:    name,
		Subject: subject,
		Text:    text,
		Html:    html,
	}

	if status == "" {
		c.Status = EmailTemplateStatusNew
	}

	return c
}
