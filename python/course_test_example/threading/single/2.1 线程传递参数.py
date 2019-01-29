#!/usr/bin/python
#!coding:utf-8
import threading
import time

 
def worker(num):
    print "Fuvism Threading num=", num
    time.sleep(1)
    return
 

if __name__ == "__main__":
    # t = threading.Thread(target=worker, args=(10))            # 错误的
    t = threading.Thread(target=worker, args=(10, ))          # 方式1
    # t = threading.Thread(target=worker, kwargs={'num': 10})     # 方式2
    t.start()

"""
"""
