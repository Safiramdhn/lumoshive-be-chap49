package utils

import (
	"context"
	"fmt"
	"golang-chap49/models"
	"time"

	"github.com/mailersend/mailersend-go"
)

func SendOTPEmail(mailSenderAPI, emailClient string, otpReq models.OTPRequst) error {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(mailSenderAPI)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := "OTP Verification"
	html := "<p>This is your OTP{{otp_verification}}. Valid until {{time_experied}}.</p>"

	from := mailersend.From{
		Name:  "Your Name",
		Email: emailClient,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Your Client",
			Email: otpReq.UserEmail,
		},
	}

	formatedExpirationTime := otpReq.ExpiredTime.Format("02/01/06 15:04")
	personalization := []mailersend.Personalization{
		{
			Email: otpReq.UserEmail,
			Data: map[string]interface{}{
				"otp_verification": otpReq.OTP,
				"time_experied":    formatedExpirationTime,
			},
		},
	}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)

	message.SetPersonalization(personalization)

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Printf("Message sent successfully! ID: %s\n", res.Header.Get("X-Message-Id"))
	return nil
}
