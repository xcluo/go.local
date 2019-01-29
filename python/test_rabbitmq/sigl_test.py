#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Longgeek <longgeek@gmail.com>

import memcache
import time

class TT(object):
    """ 单例类，只初始化一次

    获取 Memorycache 连接对象，并初始化课程数据到 Memorycache
    """

    _instance = None
    conn = None

    def __init__(self):
        print '__init__'
        pass
    def __del__():
        print 'del...'

    def __new__(cls, *args, **kwargs):
        print '__new__'
        if not cls._instance:
            cls._instance = super(TT, cls).__new__(
                cls, *args, **kwargs)

            # Initialize all Course Memory Cache Data
            cls.conn = memcache.Client(['127.0.0.1:11211'], debug=0)
            cls.conn.set('a', time.time())
            print 'init in __new__'

        return cls._instance
    def get(self):
        #time.sleep(3)
        return self.conn.get('a')


if __name__ == "__main__":
    tt = TT()
    print tt.get()
    print '--------------------'
    time.sleep(5)
    tt = TT()
    print tt.get()





