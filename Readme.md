## GO RabbitMQ Example

#### This is repository how to use RabbitMQ in golang

### Prerequisits
- Install Golang 1.13 or higher version
- install Docker
- Knowing about Go Modules
- Knowing about RabbitMQ, and pub-sub concept

### How to run

- Clone this repo
```bash
git clone https://github.com/satriyoaji/go-rabbitmq-example
```

- Run the RabbitMQ instance locally
 ```bash
  docker run -d --hostname localhost-rabbit --name test-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

- Running the publisher and consumer
```bash
# consumer
go run consumer/main.go

# publisher
go run publisher/main.go
```

- Lookup the dashboard of current running client
`Open in browser http://localhost:15672/` (the host & port based on first exposed port when you run the RabbitMQ docker image)