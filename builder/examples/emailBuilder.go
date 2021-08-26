// Package builder: builder parameter example.
//Used for hiding structs/objects implementation from the API users
package builder

import "strings"

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{email{}}
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	eb.email.from = from
	return eb
}

func (eb *EmailBuilder) To(to string) *EmailBuilder {
	eb.email.to = to
	return eb
}

func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.subject = subject
	return eb
}

func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}

func (eb *EmailBuilder) Build() *email {
	return &eb.email
}

func sendMailImpl(email *email) {
	// sending mail
}

type build func(builder *EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}