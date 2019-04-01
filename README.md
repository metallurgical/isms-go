# isms-golang
ISMS Microservice + RabbitMQ + CloudAMPQ which still Work in Progress :)

## Environment Variables

This project used environment variables to perform some of the task. So, please create `.env` and paste below code:
```
CLOUD_AMQP_URL="amqp://<username>:<password>@<amqp-url>/<vhost>"
ISMS_USERNAME=ACNE
ISMS_PASSWORD=SECRET
ISMS_BASE_URL="https://www.isms.com.my/isms_send.php"
```

## Usages

Called via PHP using RabbitMQ with following json data :

```php
json_encode(['phone' => '0169344497', 'message' => 'Wat la weii', 'type' => 'send_sms']);
```

