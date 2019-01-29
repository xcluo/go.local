import threading
import time
 
 
def worker():
    threadname = threading.currentThread().getName()
    print 'Thread %s Start', threadname
    print 'Thread %s: %f' % (threadname, time.time())
    print 'Thread %s finished!' % threadname
 

if __name__ == "__main__":
    for i in range(5):
        t = threading.Thread(target=worker)
        t.start()
