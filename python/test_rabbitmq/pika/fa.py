#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Longgeek <longgeek@gmail.com>

import pika, uuid

class Center(object):
    """控制中心类"""

    def __init__(self):
        self.connection = pika.BlockingConnection(pika.ConnectionParameters(
                host='192.168.8.239'))

        self.channel = self.connection.channel()

        #定义接收返回消息的队列
        result = self.channel.queue_declare(exclusive=True)
        self.callback_queue = result.method.queue

        self.channel.basic_consume(self.on_response,
                                   no_ack=True,
                                   queue=self.callback_queue)

        #返回的结果都会存储在该字典里
        self.response = {}

    #定义接收到返回消息的处理方法
    def on_response(self, ch, method, props, body):
        self.response[props.correlation_id] = body

    def request(self, n):
        corr_id = str(uuid.uuid4())
        self.response[corr_id] = None

        #发送计算请求，并设定返回队列和correlation_id
        self.channel.basic_publish(exchange='',
                                   routing_key='docker_task',
                                   properties=pika.BasicProperties(
                                         reply_to = self.callback_queue,
                                         correlation_id = corr_id,
                                         ),
                                   body=str(n))
        #接收返回的数据
        while self.response[corr_id] is None:
            self.connection.process_data_events()
        print '----------------'
        print self.response[corr_id]

center = Center()

center.request('center fa from xcluo')
