#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Longgeek <longgeek@gmail.com>

from kombu import Connection, Exchange, Queue

actiontype = 'create_container'
types = {
    'create_container': {
        'queue_name': 'create_container_queue',
        'router_key': 'create.container.router',
        'exchange_name': 'container',
    },
}


queue_name = types[actiontype]['queue_name']
router_key = types[actiontype]['router_key']
exchange_name = types[actiontype]['exchange_name']
exchange = Exchange(exchange_name,
                    type='topic',
                    durable=True)
queue = Queue(queue_name,
              exchange=exchange,
              routing_key=router_key)


def process_media(body, message):
    print 'start'
    print body
    # print body
    message.ack()
    print 'end'

# connections
with Connection('amqp://guest:guest@localhost:5672//') as conn:
    # consume
    with conn.Consumer(queue, callbacks=[process_media]) as consumer:
        while True:
            conn.drain_events()
