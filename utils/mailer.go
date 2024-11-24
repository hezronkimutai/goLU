package mailer

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var (
	clientID     = os.Getenv("MAIL_CLIENT_ID")
	clientSecret = os.Getenv("MAIL_CLIENT_SECRET")
	redirectURL  = os.Getenv("MAIL_REDIRECT_URL")
	accessToken  = os.Getenv("MAIL_ACCESS_TOKEN")
	refreshToken = os.Getenv("MAIL_REFRESH_TOKEN")
)

// Initialize OAuth2 configuration
func getOAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
		Endpoint:     google.Endpoint,
	}
}

// Create token object
func getToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       time.Now().Add(time.Hour),
		TokenType:    "Bearer",
	}
}

// SendMail sends an email using the Gmail API
func SendMail(content, to, from, subject string) {
	ctx := context.Background()
	oauthConfig := getOAuth2Config()
	token := getToken()

	// Create HTTP client with OAuth2 token
	client := oauthConfig.Client(ctx, token)

	// Use NewService with the HTTP client
	service, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Gmail service: %v", err)
	}

	// Compose the email
	email := fmt.Sprintf(
		"Content-Type: text/html; charset=\"UTF-8\"\nMIME-Version: 1.0\nContent-Transfer-Encoding: 7bit\nto: %s\nfrom: %s\nsubject: %s\n\n%s",
		to, from, subject, content,
	)

	// Encode email in base64 (URL-safe)
	rawMessage := base64.URLEncoding.EncodeToString([]byte(email))

	// Send email
	message := &gmail.Message{
		Raw: rawMessage,
	}
	_, err = service.Users.Messages.Send("me", message).Do()
	if err != nil {
		log.Fatalf("Unable to send email: %v", err)
	}

	fmt.Println("Email sent successfully!")
}
