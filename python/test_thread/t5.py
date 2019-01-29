import threading
import time


lock = threading.Lock()


def worker1():
    if lock.acquire():
        time.sleep(2)
        print 'worker1'
        lock.release()
 

def worker2():
    if lock.acquire():
        time.sleep(1)
        print 'worker2'
        lock.release()



if __name__ == "__main__":
    t1 = threading.Thread(target=worker1)
    t2 = threading.Thread(target=worker2)

    t1.start()
    time.sleep(1)
    t2.start()



