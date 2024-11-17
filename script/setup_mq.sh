#!/bin/bash
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_VHOST=test -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password  rabbitmq:4.0-management