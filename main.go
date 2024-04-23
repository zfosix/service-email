package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func sendMailSimple(subject string, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"udinudinsedunia68@gmail.com",
		"pydaxjawijnlpwyi",
		"smtp.gmail.com",
	)

	msg := "Subject:" + subject + "\n\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"udinudinsedunia68@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func sendMailSimpleHTML(subject string, templatePath string, to []string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	err = t.Execute(&body, struct{ Name string }{Name: "Fajar"})
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	auth := smtp.PlainAuth(
		"",
		"udinudinsedunia68@gmail.com",
		"pydaxjawijnlpwyi",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject:" + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"udinudinsedunia68@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func sendGomail(templatePath string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	err = t.Execute(&body, struct{ Name string }{Name: "Fajar"})
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "udinudinsedunia68@gmail.com")
	m.SetHeader("To", "udinudinsedunia68@gmail.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	m.Attach("./self.jpeg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "udinudinsedunia68@gmail.com", "pydaxjawijnlpwyi")

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}

// func sendSendgrid() {
// 	from := mail.NewEmail("Fajar", "hello@fajar.com")
// 	subject := "Sending with Twilio SendGrid is Fun"
// 	to := mail.NewEmail("Example User", "udinudinsedunia68@gmail.com")
// 	plainTextContent := "and easy to do anywhere, even with Go"
// 	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
// 	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
// 	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
// 	response, err := client.Send(message)
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		fmt.Println(response.StatusCode)
// 		fmt.Println(response.Body)
// 		fmt.Println(response.Headers)
// 	}
// }

func main() {
	// sendMailSimple("Another subject", "Another body", []string{"udinudinsedunia68@gmail.com"})
	// sendMailSimpleHTML("Another subject", "./test.html", []string{"udinudinsedunia68@gmail.com"})
	sendGomail("./test.html")
	// sendSendgrid()
}
