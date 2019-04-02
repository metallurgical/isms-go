# isms-golang
ISMS Microservice + RabbitMQ + CloudAMPQ which still Work in Progress :). This package should be run as a background process. 

The ideas is simple, imagine we have one main web application running on top of PHP or any other languages. Instead of calling directly SMS API using HTTP handler from main application(also called the producer), the main application should invoke the RabbitMQ message broker to queue a message along with the "receiver - receiver here is the one who will receive the SMS".

The consumer(which is in this case is our GO package) background services will always listen and waiting to the incoming queue. Any queue received will run as goroutine to call the SMS API. Note that the producer, consumer and broker do not have to reside on the same host, it could be hosted on different servers which eventually decoupling some of the features from main application.

## Requirements
- CloudAMQP account(Free account will do - www.cloudamqp.com)
- ISMS account(www.isms.com.my)
- Tested GO on version go1.11.4 darwin/amd64
- Tested RabbitMQ on version 3.7.12


## Environment Variables
This project used environment variables to perform some of the task. So, please create `.env` and paste below code:
```
CLOUD_AMQP_URL="amqp://<username>:<password>@<amqp-url>/<vhost>"
SMS_USERNAME=ACNE
SMS_PASSWORD=SECRET
SMS_BASE_URL="https://www.isms.com.my"
SMS_SEND_URL="/isms_send.php"
SMS_CHECK_BALANCE_URL="/isms_balance.php"
SMS_MESSAGE_PREFIX="CompanyName"
```

## Usages

Called via main application using RabbitMQ services with the following json data running on top of PHP language:

```php
json_encode(['phone' => '0169344497', 'message' => 'Wat la weii', 'type' => 'send_sms']);
```

