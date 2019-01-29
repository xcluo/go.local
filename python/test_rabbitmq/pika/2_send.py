#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Longgeek <longgeek@gmail.com>

import pika
import sys

message = ' '.join(sys.argv[1:]) or "Hello World!!!!"
connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()
channel.queue_declare(queue='task_queue', durable=True)
channel.basic_publish(exchange='', routing_key='hello', body=message, properties=pika.BasicProperties(delivery_mode = 2,))
print "Sent 'Hello World!'"
connection.close()
