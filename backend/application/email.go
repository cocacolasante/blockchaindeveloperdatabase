package application

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Message struct {
	From           string
	FromName       string
	To             string
	Subject        string
	Attachments    []string
	AttachmentsMap map[string]string
	Data           any
	DataMap        map[string]any
	Template       string
}

//

func (m *Mail) SendEmail(msg Message) error {
	if msg.Template == "" {
		msg.Template = "mail"
	}

	if msg.From == "" {
		msg.From = m.FromAddress
	}
	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	if msg.AttachmentsMap == nil {
		msg.AttachmentsMap = make(map[string]string)
	}
	if len(msg.DataMap) == 0 {
		msg.DataMap = make(map[string]any)

	}

	msg.DataMap["message"] = msg.Data

	// build html mail
	formattedMessage, err := m.BuildHTMLMessage(msg)
	if err != nil {
		log.Println(err)
		return err
	}

	// built plain text mail
	plainMessage, err := m.BuildPlainMessage(msg)
	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.GetEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	

	smtpClient, err := server.Connect()
	if err != nil {
		log.Println("send email line 86:", err)
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)

	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, x := range msg.Attachments {
			email.AddAttachment(x)

		}
	}

	if len(msg.AttachmentsMap) > 0 {
		for key, value := range msg.AttachmentsMap {
			email.AddAttachment(value, key)
		}
	}

	err = email.Send(smtpClient)
	if err != nil {
		log.Println("send email line 111:", err)
		log.Println(err)
		return err
	}

	return nil

}

func (m *Mail) BuildHTMLMessage(msg Message) (string, error) {
	log.Println("hit build html")
	templateToRender := fmt.Sprintf("./application/templates/%s.html.gohtml", msg.Template)

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()
	formattedMessage, err = m.InlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *Mail) BuildPlainMessage(msg Message) (string, error) {
	templateToRender := fmt.Sprintf("./application/templates/%s.plain.gohtml", msg.Template)

	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMessage := tpl.String()

	return plainMessage, nil
}

func (m *Mail) InlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil

}

func (m *Mail) GetEncryption(e string) mail.Encryption {
	switch e {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}

func (app *Application) CreateMail() Mail {
	app.InfoLog.Println("Starting mailer")
	m := Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromAddress: "noreply@blockchaindatabase.com",
		FromName:    "No Reply",
		
	}

	return m
}

func (m *Mail) CreateMessage(toAddress, subject, template string, data map[string]any) (Message, error) {
	var msg Message
	msg.From = m.FromAddress
	msg.FromName = m.FromName
	msg.To = toAddress
	msg.Subject = subject
	msg.Template = template

	msg.DataMap = data

	return msg, nil

}

// func (app *Application) ListenForMail() {
// 	for {
// 		select {
// 		case msg := <-app.Mailer.MailerChan:
// 			go app.Mailer.SendMail(msg, app.Mailer.ErrorChan)
// 		case err := <-app.Mailer.ErrorChan:
// 			app.ErrorLog.Println(err)
// 		case <-app.Mailer.DoneChan:
// 			return
// 		}
// 	}
// }
