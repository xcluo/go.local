#!/usr/bin/python
#!coding:utf-8
import threading
import time
 
def worker():
    print "Fuvism Threading Name"
    time.sleep(1)
    return
 

if __name__ == "__main__":
    t = threading.Thread(target=worker)
    t.start()

"""
大多数线程程序并不适合使用 print 来调试，比如在生成环境，比如守护线程，都在使用 print 的有限制(需要终端)。

logging 模块支持将线程名嵌入到日志中，并且通过 level 控制不通的日志记录级别。方便对线程，尤其是多线程程序的调试追踪。

logging 是线程安全的，也就是说来自不同的线程的消息，输出的时候会有区别，而不是搅合在一起，这个问题我们在实践中可以逐步感受。
"""
import logging
import threading
import time


logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)s] (%(threadName)-10s) %(message)s',
    )

 
def worker():
    logging.debug("Fuvism Threading")
    time.sleep(1)
    return
 

if __name__ == "__main__":
    t = threading.Thread(target=worker)
    t.start()
