import threading
import time


def worker1():
    time.sleep(2)
    print 'worker1'
 

def worker2():
    time.sleep(1)
    print 'worker2'


if __name__ == "__main__":
    t1 = threading.Thread(target=worker1)
    t2 = threading.Thread(target=worker2)

    t1.start()
    t1.join()
    t2.start()
