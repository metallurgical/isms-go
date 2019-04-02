package main

import (
	"github.com/streadway/amqp"
	"isms/sms"
	"isms/sms/util"
	"log"
	"os"
)

func main() {
	// Load environment variables.
	util.LoadEnv()
	connection, err := amqp.Dial(os.Getenv("CLOUD_AMQP_URL"))
	defer connection.Close()

	if err != nil {
		util.FailResponse(err, "Cannot connect to CloudAMQP service provider")
	}

	channel, err := connection.Channel()
	defer channel.Close();

	if err != nil {
		util.FailResponse(err, "Failed to initialed channel")
	}

	queueName := "isms"

	q, err := channel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		util.FailResponse(err, "Failed to declare a queue")
	}

	messages, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		util.FailResponse(err, "Queue failed to bind to consumer")
	}

	// Declare boolean channel
	forever := make(chan bool)

	go func() {
		for response := range messages {
			// Parse Json String into struct for easy to access properties.
			contact := util.ParseResponse(response.Body)

			switch contact.Type {
				case "send_sms":
					sms.SendDirectSMS(contact)
				case "check_balance":
					sms.CheckBalance(contact);
				default:
					sms.SendDirectSMS(contact)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<- forever
}
