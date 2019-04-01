package util

import (
	"github.com/joho/godotenv"
	"encoding/json"
	"net/url"
	"strings"
	"log"
	"os"
)

type Contact struct {
	Type string
	Phone string
	Message string
	Username string
	Password string
	Prefix string
}

// Parse response from JSON string into struct. Main application
// must send json string via Message Broker.
func ParseResponse(jsonBuffer []byte) (Contact) {
	var contact Contact

	// Map json string into struct data type under `Contact` struct
	err := json.Unmarshal(jsonBuffer, &contact)
	if err != nil {
		FailResponse(err, "Failed to parse json string into readable struct format")
	}

	return contact
}

// Show error response along with Exit Code 1.
func FailResponse(err error, msg string) {
	log.Fatalf("%s : %s", msg, err)
}

// Validate phone number coming from response.
func (contact *Contact) ValidatedPhone() (Contact) {
	// If passed phone data doesnt have `6` in front, then add prefix `6`
	if !strings.HasPrefix(contact.Phone, "6") {
		contact.Phone = "6" + contact.Phone
	}

	return *contact
}

// Load environment variables.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		FailResponse(err, "Failed to load environment variable. Make sure you create .env file specific to your project.")
	}
}

// Populate URL and form query string to call the final ISMS api.
func PopulateUrl(contact *Contact) string {
	ismsUrl := os.Getenv("ISMS_BASE_URL")

	baseUrl, err := url.Parse(ismsUrl)

	if err != nil {
		FailResponse(err, "Malformed URL")
	}

	queryString := url.Values{}

	if contact.Username == "" {
		queryString.Set("un", os.Getenv("ISMS_USERNAME"))
	}

	if contact.Password == "" {
		queryString.Set("pwd", os.Getenv("ISMS_PASSWORD"))
	}

	queryString.Set("dstno", contact.Phone)
	queryString.Set("msg", contact.Prefix + ":" + contact.Message)
	queryString.Set("type", "1")

	baseUrl.RawQuery = queryString.Encode()

	return baseUrl.String()
}
