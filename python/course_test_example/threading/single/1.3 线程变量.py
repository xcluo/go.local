#!/usr/bin/python
#!coding:utf-8
import threading
import time

 
def worker():
    print "Fuvism Threading"
    print threading.currentThread()
    print dir(threading.currentThread())
    time.sleep(1)
    return
 

if __name__ == "__main__":
    t = threading.Thread(target=worker)
    t.start()

"""
前面的重命名部分，我们使用 threading.currentThread()。

它的作用，是在线程内部，获取当前线程的实例变量：

<Thread(Thread-1, started 139700257748736)>

通过 dir() 函数，可以看到它包含的属性和方法，常用的有：

  getName()
  setName()

这两个我们已经实践过了，还包括了：

  isAlive()：监测线程是否 "活着"
  isDaemon()：是否是一个守护线程(后面我们将研究守护线程)
  name: 属性，获取线程名，和 getName() 返回结果一样
  join()：阻塞，多线程部分我们深入研究
  run()：线程执行内容的内容，主要是在 threading.Thread 类中重写线程部分的代码
  start()：用来启动线程

"""
