# encoding: UTF-8
import threading
import time
 

event = threading.Event()

 
def worker():
    print '%s waiting for event...' % threading.currentThread().getName()
    event.wait()
    print '%s received event.' % threading.currentThread().getName()
 

if __name__ == "__main__":
    t1 = threading.Thread(target=worker)
    t2 = threading.Thread(target=worker)
    t1.start()
    t2.start()
 
    time.sleep(2)
    print 'Event Start.'
    event.set()
