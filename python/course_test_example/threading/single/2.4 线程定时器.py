#!/usr/bin/python
#!coding:utf-8
import logging
import threading
import time
 

logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)s] (%(threadName)-10s) %(message)s',
    )

 
def worker():
    logging.debug("Fuvism Threading")
    time.sleep(2)
    return
 

if __name__ == "__main__":
    print 'Start...'
    t = threading.Timer(3, worker)
    t.start()


"""
线程定时，是在启动一个线程的时候，指定一个线程多少秒后执行。通过 threading.Timer 实现。
"""
