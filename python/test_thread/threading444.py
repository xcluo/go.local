import threading
import time
 
data = 0
lock = threading.Lock()
 
def func():
    global data
    print '%s acquire lock...' % threading.currentThread().getName()
    
    if lock.acquire():
    # if True:
        print '%s get the lock.' % threading.currentThread().getName()
        data += 1
        time.sleep(2)
        print '%s release lock...' % threading.currentThread().getName()
        
        lock.release()
    print 'data = ', data
 
t1 = threading.Thread(target=func)
t2 = threading.Thread(target=func)
t3 = threading.Thread(target=func)
t1.start()
t2.start()
t3.start()
