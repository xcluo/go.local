#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Longgeek <longgeek@gmail.com>

import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()
channel.queue_declare(queue='hello')

def callback(ch, method, properties, body):
    print "Received %r" % (body,)

channel.basic_consume(callback, queue='hello', no_ack=True)
print 'Waiting for messages. to exit press CTRL+C'
channel.start_consuming()
