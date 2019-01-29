import threading
import time
 
 
def worker1():
    print 'Thread Worker1 Start'
    print 'Thread Worker1: ', time.time()
    print 'Thread Worker1 finished!'
 
 
def worker2():
    print 'Thread Worker2 Start'
    print 'Thread Worker2: ', time.time()
    print 'Thread Worker2 finished!'
 

if __name__ == "__main__":
    t1 = threading.Thread(target=worker1)
    t2 = threading.Thread(target=worker2)
    t1.start()
    t2.start()
