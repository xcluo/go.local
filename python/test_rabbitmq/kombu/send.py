#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Longgeek <longgeek@gmail.com>

import kombu
import sys

connection = kombu.Connection('amqp://guest:guest@192.168.8.215:5672//')
simple_queue = connection.SimpleQueue('World')

if len(sys.argv) < 2:
    print 'message is empty!'
    sys.exit(0)

message = sys.argv[1]
simple_queue.put(message)
print '[X] send: %s \n' % message
simple_queue.close()

