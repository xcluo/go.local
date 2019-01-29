#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Longgeek <longgeek@gmail.com>

from kombu import Connection, Exchange

types = {
    'create_container': {
        'queue_name': 'create_container_queue',
        'router_key': 'create.container.router',
        'exchange_name': 'container',
    },
    'atest': {
        'queue_name': 'network',
        'router_key': 'network',
        'exchange_name': 'test'
    }
}

def send(type, messages):
    router_key = types[type]['router_key']
    exchange_name = types[type]['exchange_name']

    exchange = Exchange(exchange_name,
                        type='topic',
                        durable=True)
    # connection
    with Connection('amqp://guest:guest@localhost:5672//') as conn:
        # produce
        producer = conn.Producer(serializer='json', exchange=exchange)
        producer.publish(messages,
                         exchange=exchange,
                         routing_key=router_key,)
        print 'done'
if __name__ == "__main__":
    send('atest', "{'a': 1}")
