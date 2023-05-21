# Logs Monitoring

***Log monitoring system based on Golang, using the RabbitMQ message broker for message processing and Nginx for load balancing***

# Requirements
* **Go** `1.20.3`
* **MongoDB** `4.4.22`
* **Gin** `1.9.0`
* **RabbitMq** `3.11.15`
* **Logrus** `1.9.0`


## Got Notification to email in case critical error in system app 
```yaml
Description: an error occurred while sending a message to rabbitmq
From: myusername1234@mail.ru
To: recipient@mail.ru
Subject: Logs Notification
MIME-Version: 1.0
Content-Type: text/plain; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Level: error
```

## Logs Example
```json
{"level":"info","msg":"Success connected to mongoDB","time":"2023-05-20T16:52:15+03:00"}

{"level":"info","msg":"Connected to rabbitMq Success [Producer]","time":"2023-05-20T17:03:53+03:00"}

{"level":"error","msg":"cannot transform type bson.Raw to a BSON Document: length read exceeds number of bytes available. length=37 bytes=1701585531","time":"2023-05-20T17:32:49+03:00"}

{"level":"fatal","msg":"Couldn't connect to rabbitMq [Producer]dial tcp 127.0.0.1:5673: connect: connection refused","time":"2023-05-20T17:17:45+03:00"}

{"level":"info","msg":"Success inserted with ObjectID(\"6469cd3e43cffc658603027b\")","time":"2023-05-21T10:50:22+03:00"}
```

## Build / Run

```shell
git clone https://github.com/TechSir3n/LogMonitoringSystem.git
cd logs-monitoring
make build
make run 
```

## Task need to do 
1. Implement log filtering by severity level (debug, info, warning, error, critical) and send only necessary logs. **Done**
2. Add the ability to save logs to the database for further analysis. **Done**
3. Write a script to automatically delete old logs so as not to take up much disk space. **Done**
4. Implement the ability to send notifications of critical errors via email or SMS. **Done**
5. Add support for various log formats (JSON, CSV, XML, etc.).
6. Implement application performance monitoring based on logs (request processing speed, response time, etc.).
7. Add the ability to aggregate logs from multiple sources (several servers or applications).
8. Implement automatic detection of application problem areas based on logs and offer solutions to improve performance.
9. Add support for multilingual logs (localization).
10. Implement the ability to configure the rules for processing logs through the web interface or the configuration file. **Done**