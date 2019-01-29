#!/usr/bin/python
#!coding:utf-8
import threading
import time

 
def worker():
    print "Fuvism Threading isAlive=", threading.currentThread().isAlive()
    time.sleep(5)
    return
 

if __name__ == "__main__":
    t = threading.Thread(target=worker)
    print 'before start =', t.isAlive()
    t.start()
    for i in range(10):
        time.sleep(1)
        print 'after start =', t.isAlive()

"""
内部调用的时候，isAlive() 必然返回 True，否则也就不会执行代码了。
"""

"""
在外部，线程创建的时候，我们也能得到 currentThread() 实例，在 start() 前，线程的状态是 False 的。
"""

"""
当然，如果线程僵死，或执行完成，状态也将变化。
"""
