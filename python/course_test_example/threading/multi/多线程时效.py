#!/usr/bin/python
#!coding:utf-8
import threading
import time
 
def worker():
    print "Fuvism Threading"
    time.sleep(1)
    return
 

if __name__ == "__main__":
    print '开始于=', time.ctime()
    # for i in xrange(5):
    #     t = threading.Thread(target=worker)
    #     t.start()

    for i in xrange(5):
        worker()

    print '结束于=', time.ctime()

"""
函数执行的时候，是一个一个的，需要5秒
而线程执行的时候，是同时执行的，感觉像只花费了1秒
"""
