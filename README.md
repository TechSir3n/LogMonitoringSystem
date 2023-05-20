# Logs Monitoring

***Log monitoring system based on Golang, using the RabbitMQ message broker for message processing and Nginx for load balancing***


## Got Notification to email in case critical error in system app 
Description: an error occurred while sending a message to rabbitmq<br>
From: myusername1234@mail.ru<br>
To: recipient@mail.ru<br>
Subject: Logs Notification<br>
MIME-Version: 1.0<br>
Content-Type: text/plain; charset="utf-8"<br>
Content-Transfer-Encoding: quoted-printable<br>
Level: error <br>

## Logs Example
```json
{"level":"info","msg":"Success connected to mongoDB","time":"2023-05-20T16:52:15+03:00"}

{"level":"info","msg":"Connected to rabbitMq Success [Producer]","time":"2023-05-20T17:03:53+03:00"}

{"level":"error","msg":"cannot transform type bson.Raw to a BSON Document: length read exceeds number of bytes available. length=37 bytes=1701585531","time":"2023-05-20T17:32:49+03:00"}


{"level":"fatal","msg":"Couldn't connect to rabbitMq [Producer]dial tcp 127.0.0.1:5673: connect: connection refused","time":"2023-05-20T17:17:45+03:00"}
```

## Task need to do 
1. Implement log filtering by severity level (debug, info, warning, error, critical) and send only necessary logs. **Done**
2. Add the ability to save logs to the database for further analysis. **Done**
3. Write a script to automatically delete old logs so as not to take up much disk space.
4. Implement the ability to send notifications of critical errors via email or SMS. **Done**
5. Add support for various log formats (JSON, CSV, XML, etc.).
6. Implement application performance monitoring based on logs (request processing speed, response time, etc.).
7. Add the ability to aggregate logs from multiple sources (several servers or applications).
8. Implement automatic detection of application problem areas based on logs and offer solutions to improve performance.
9. Add support for multilingual logs (localization).
10. Implement the ability to configure the rules for processing logs through the web interface or the configuration file.