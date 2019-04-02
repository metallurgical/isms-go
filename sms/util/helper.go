package util

import (
	"github.com/joho/godotenv"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

// Check balance URL.
func (contact *Contact) CheckBalanceUrl() (*url.URL) {
	smsUrl := os.Getenv("SMS_BASE_URL") + os.Getenv("SMS_CHECK_BALANCE_URL")
	url, queryString := contact.DefaultQueryString(smsUrl)
	url.RawQuery = queryString.Encode()

	return url
}

// Sending Simple SMS and TAC url.
func (contact *Contact) DirectUrl() (*url.URL) {
	smsUrl := os.Getenv("SMS_BASE_URL") + os.Getenv("SMS_SEND_URL")
	url, queryString := contact.DefaultQueryString(smsUrl)
	queryString.Set("mobile", contact.Phone)
	queryString.Set("message", contact.Prefix + ":" + contact.Message)
	queryString.Set("type", "1")
	queryString.Set("sender", os.Getenv("SMS_USERNAME"))
	url.RawQuery = queryString.Encode()

	return url
}

// Populate URL and form query string to call the final SMS api.
func (contact *Contact) DefaultQueryString(baseUrl string) (*url.URL, url.Values) {
	smsUrl, err := url.Parse(baseUrl)

	if err != nil {
		FailResponse(err, "Malformed URL")
	}

	queryString := url.Values{}

	if contact.Username == "" {
		queryString.Set("username", os.Getenv("SMS_USERNAME"))
	}

	if contact.Password == "" {
		queryString.Set("password", os.Getenv("SMS_PASSWORD"))
	}

	if contact.Prefix == "" {
		contact.Prefix = os.Getenv("SMS_MESSAGE_PREFIX")
	}

	return smsUrl, queryString
}

// Send sms into mobile number doesnt matter by
// direct SMS or TAC. Just send it anyway.
func (contact *Contact) Send(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		FailResponse(err, "Error sending SMS or checking balance")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	return body
}
