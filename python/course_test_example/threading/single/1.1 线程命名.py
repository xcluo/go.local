#!/usr/bin/python
#!coding:utf-8
import threading
import time
 
def worker():
    print "Fuvism Threading Name=", threading.currentThread().getName()
    time.sleep(1)
    return
 

if __name__ == "__main__":
    # t = threading.Thread(target=worker)       # 未改名
    t = threading.Thread(name='MyThreading', target=worker)     # 重命名
    t.start()

"""
默认的，每个线程实例都有一个名词，只不过这个名称是以 Thread 开头，附加一个递增的数字构成，缺少友好度。

在线程编程中，尤其是包含多线程编程的时候，通常设定一个别名，方便调试和查找。
"""
