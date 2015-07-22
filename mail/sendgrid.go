package mail

import (
	"fmt"
	"net/mail"

	"gopkg.in/sendgrid/sendgrid-go.v2"
)

type Message interface {
	From() *mail.Address
	ReplyTo() *mail.Address
	To() []*mail.Address
	Cc() []*mail.Address
	Bcc() []*mail.Address
	Subject() string
	Text() string
	HTML() string
}
type SendGridOptions interface {
	APIKey() string
}

type SendGridSender interface {
	Init(SendGridOptions)
	Send(Message) error
}

func NewSendGridOptions(apiKey string) SendGridOptions {
	return &sendGridOptions{apiKey: apiKey}
}

type message struct {
	from    string
	replyTo string
	to      string
	cc      string
	bcc     string
	subject string
	text    string
	html    string
}

type sendGridOptions struct {
	apiKey string
}

type sender struct {
	Options SendGridOptions
}

func (m *message) From() *mail.Address    { return getSingleAddress(m.from) }
func (m *message) ReplyTo() *mail.Address { return getSingleAddress(m.replyTo) }
func (m *message) To() []*mail.Address    { return getMultipleAddress(m.to) }
func (m *message) Cc() []*mail.Address    { return getMultipleAddress(m.cc) }
func (m *message) Bcc() []*mail.Address   { return getMultipleAddress(m.bcc) }
func (m *message) Subject() string        { return m.subject }
func (m *message) Text() string           { return m.text }
func (m *message) HTML() string           { return m.html }

func NewMessage(from, to, replyto, cc, bcc, subject, text, html string) Message {
	return &message{
		from:    from,
		to:      to,
		replyTo: replyto,
		cc:      cc,
		bcc:     bcc,
		subject: subject,
		text:    text,
		html:    html,
	}
}

func (s *sender) Init(options SendGridOptions) {
	s.Options = options
}

func (s *sender) Send(m Message) error {
	if s.Options == nil {
		return fmt.Errorf("must initialize sender first")
	}
	sg := sendgrid.NewSendGridClientWithApiKey(s.Options.APIKey())
	msg := sendgrid.NewMail()
	msg.SetFromEmail(m.From())
	msg.SetReplyToEmail(m.ReplyTo())
	for _, to := range m.To() {
		msg.AddTo(to.Address)
		msg.AddToName(to.Name)
	}
	msg.AddCcRecipients(m.Cc())
	msg.AddBccRecipients(m.Cc())
	msg.SetSubject(m.Subject())
	msg.SetText(m.Text())
	msg.SetHTML(m.HTML())
	return sg.Send(msg)
}

func (s *sendGridOptions) APIKey() string { return s.apiKey }

func getSingleAddress(raw string) *mail.Address {
	addr, err := mail.ParseAddress(raw)
	if err != nil {
		return nil
	}
	return addr
}
func getMultipleAddress(raw string) []*mail.Address {
	addr, err := mail.ParseAddressList(raw)
	if err != nil {
		return nil
	}
	return addr
}

var _ SendGridOptions = (*sendGridOptions)(nil)
var _ SendGridSender = (*sender)(nil)
var _ Message = (*message)(nil)
