package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

// Environment variables (replace with actual values or load from your env)
var (
	clientID     = "YOUR_CLIENT_ID"
	clientSecret = "YOUR_CLIENT_SECRET"
	redirectURL  = "YOUR_REDIRECT_URL"
	accessToken  = "YOUR_ACCESS_TOKEN"
	refreshToken = "YOUR_REFRESH_TOKEN"
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
		Expiry:       time.Now().Add(time.Hour), // Replace with actual expiry time if available
		TokenType:    "Bearer",
	}
}

// SendMail sends an email using the Gmail API
func SendMail(content, to, from, subject string) {
	ctx := context.Background()
	oauthConfig := getOAuth2Config()
	token := getToken()

	// Create OAuth2 client
	client := oauthConfig.Client(ctx, token)

	// Create Gmail service
	service, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to create Gmail service: %v", err)
	}

	// Compose the email
	email := fmt.Sprintf("Content-Type: text/html; charset=\"UTF-8\"\nMIME-Version: 1.0\nContent-Transfer-Encoding: 7bit\nto: %s\nfrom: %s\nsubject: %s\n\n%s",
		to, from, subject, content)

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

func main() {
	SendMail(
		"<h1>Hello from Go!</h1><p>This is a test email.</p>",
		"recipient@example.com",
		"your_email@example.com",
		"Test Email from Go",
	)
}
