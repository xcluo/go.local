#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Longgeek <longgeek@gmail.com>

from kombu import Connection, Exchange, Queue

media_exchange = Exchange('media', 'direct', durable=True)
video_queue = Queue('video', exchange=media_exchange, routing_key='video')

# connection
with Connection('amqp://guest:guest@localhost:5672//') as conn:
    # produce
    producer = conn.Producer(serializer='json')
    producer.publish({'name': '/tmp/locat1.avi', 'size': 13020},
                     exchange=media_exchange, routing_key='video', on_return=True,
                     declare=[video_queue])
