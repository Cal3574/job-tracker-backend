package email

import (
	"fmt"
	"log"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v3"
)

// SendEmail sends an email using Mailjet API.
func SendEmail(to, name, subject, body string) error {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "Callumhall65@gmail.com",
				Name:  "JobTrackr",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
					Name:  name,
				},
			},
			Subject:  subject,
			TextPart: body,
			HTMLPart: fmt.Sprintf(
				`<h3>Dear %s. You have an upcoming interview tomorrow!</a>!</h3><br /> %s`,
				name, body,
			),
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
	return nil
}
