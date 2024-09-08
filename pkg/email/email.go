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
				`<html>
					<head>
						<style>
							body { font-family: Arial, sans-serif; color: #333; }
							.container { max-width: 600px; margin: auto; padding: 20px; border: 1px solid #ddd; border-radius: 8px; background-color: #f9f9f9; }
							h3 { color: #007bff; }
							p { line-height: 1.6; }
							.footer { margin-top: 20px; font-size: 0.8em; color: #888; }
						</style>
					</head>
					<body>
						<div class="container">
							<h3>Dear %s,</h3>
							<p>I hope you are feeling prepared!</p>
							<p>%s</p>
							<p>Best of luck.</p>
							<div class="footer">
								<p>Thank you for using JobTrackr.</p>
								<p>Regards,<br />The JobTrackr Team</p>
							</div>
						</div>
					</body>
				</html>`, name, body,
			),
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}
	log.Printf("Email sent successfully: %+v", res)
	return nil
}
