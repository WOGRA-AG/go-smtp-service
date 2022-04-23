# go-smtp-service

The go smtp service is a webservice for sending mails to recipients. This is a first trial to develop such a service in go.

## Configuration

You need a config.yml in the directory of execution. 

A sample content of the config file is here:

```
  User: "smtp_user"
  Password: "smtp_password"
  Smtpserver: "example.smtp.server.com"
  Smtpport: "25"
  SenderAddress: "example@example.com"
```

To make sure that the service will only used by granted user you need an accestoken.yml file in the execution directory:

```
  Accesstokens:
    - token: "1"
      user: "internal"
      password: "internal"
    - token: "2"
      user: "internal"
      password: "internal"
    - token: "3"
      user: "internal"
      password: "internal"
```

At the moment the default port is 8080. We will add a further server config in the next steps.

Here a sample call for the service:

Post Service:

```
  localhost:8080/sendMail/
```

Body of the Postservice:

```
  {
          "sender": "example@example.com",
          "receiver": ["first_recipient@example.com", "second_recipient@example.com"],
          "subject": "This is a Test",
          "message": "I want to test this rest service for sending mails",
          "accesstoken": {"accesstoken":"1", "user":"internal", "password": "internal" }
  }
```

The accesstoken given in the body will be checked to make sure that the caller has the right to call this service.

