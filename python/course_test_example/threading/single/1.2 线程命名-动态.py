#!/usr/bin/python
#!coding:utf-8
import threading
import time
 
def worker():
    print "Fuvism Threading Name=", threading.currentThread().getName()
    threading.currentThread().setName('My Fuvism Threading')
    print "Fuvism Threading Name=", threading.currentThread().getName()
    time.sleep(1)
    return
 

if __name__ == "__main__":
    t = threading.Thread(target=worker)
    t.start()

"""
通过 setName() 可以动态给线程命名。

注意：这个方法的使用，是在线程内部。
"""
