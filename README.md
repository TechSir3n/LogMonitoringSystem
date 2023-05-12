# Logs Monitoring

***Log monitoring system based on Golang, using the RabbitMQ message broker for message processing and Nginx for load balancing***


## Got Notification to email in case critical error in system app 
Description: an error occurred while sending a message to rabbitmq
From: myusername1234@mail.ru
To: recipient@mail.ru
Subject: Logs Notification
MIME-Version: 1.0
Content-Type: text/plain; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Level: error 

## Task need to do 
1. Implement log filtering by severity level (debug, info, warning, error, critical) and send only necessary logs. **Done**
2. Add the ability to save logs to the database for further analysis.
3. Write a script to automatically delete old logs so as not to take up much disk space.
4. Implement the ability to send notifications of critical errors via email or SMS. **Done**
5. Add support for various log formats (JSON, CSV, XML, etc.).
6. Implement application performance monitoring based on logs (request processing speed, response time, etc.).
7. Add the ability to aggregate logs from multiple sources (several servers or applications).
8. Implement automatic detection of application problem areas based on logs and offer solutions to improve performance.
9. Add support for multilingual logs (localization).
10. Implement the ability to configure the rules for processing logs through the web interface or the configuration file.