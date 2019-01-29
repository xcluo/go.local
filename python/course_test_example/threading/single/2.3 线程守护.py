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
    logging.debug("Fuvism Threading Start")
    time.sleep(2)
    logging.debug("Fuvism Threading Exist")
    return
 

if __name__ == "__main__":
    t = threading.Thread(target=worker)
    t.setDaemon(False)        # 没有输出
    # t.setDaemon(True)
    t.start()
    """
    守护线程，就是指线程可以一直运行而不阻塞主程序退出。
    
    默认情形，或者到目前为止，所有的范例都是需要等待所有线程完成工作后才能退出(control + c 除外)。
    
    如果需要创建一个线程作为守护线程，也就是 daemon 程序，是的服务不能自动终止，以保证可以随时提供服务。比如监控程序，就需要一直运行，每秒/每分钟...处理监控和报警，这个时候程序就不能只运行一次而终止。
    
    要使用守护线程，只需要标志一个线程为守护线程即可，使用方法:
    
    setDaemon()
    """

    # t.join()
    """
    输出中没有守护线程的 Exist 消息，因为守护线程在执行之前，本身所在的程序已经结束了。如果需要主程序等待守护线程的执行完成，那么需要 join() 函数。
    """

    # t.join(1)
    """
    但是join() 会无限制的阻塞下去。可以通过给 join 传递一个浮点数字，来控制党线程变为不活跃状态所需的时间，单位是秒。
    也就是说，党线程在这个指定的时间段没有结束，那么 join() 也会返回。

    这里可以看出来，线程只输出了 "Fuvism Threading Start"，而没有 "Fuvism Threading Exist" 信息，是由于我们限制了 join() 也就是等待时间是1秒。

    这样做的目的是，防止线程无限制的死掉，而导致服务器资源浪费在一些僵死的线程中。
    当然如果不需要等待结果，或者不好预测线程结束时间，并且通过编程尽量保证了线程不死(大多生产情形)，但依然可以在测试的时候，通过join() 来观察输出，还是很有帮助的。
    """
