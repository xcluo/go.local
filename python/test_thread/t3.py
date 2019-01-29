import threading
import time


data = []
res = 0
 
 
def worker1():
    global data
    for i in range(100):
        data.append(i)
        time.sleep(0.1)
 

def worker2():
    global data
    global res
    time.sleep(5)
    res = sum(data)

if __name__ == "__main__":
    t1 = threading.Thread(target=worker1)
    t2 = threading.Thread(target=worker2)
    t1.start()
    t2.start()

    for i in range(10):
        print 'res = ', res
        time.sleep(1)
