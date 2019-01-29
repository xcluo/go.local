import time
import multiprocessing  
from  multiprocessing import Process,Queue

  
def worker1(queue, i): 
    name = multiprocessing.current_process().name
    if queue.get():
        print 'Process %s i =%d'% (name, i)


def worker2(queue, i): 
    name = multiprocessing.current_process().name
    if queue.get():
        print 'Process %s i =%d'% (name, i)

  
if __name__ == '__main__':  
    queue = multiprocessing.Queue()

    p1 = multiprocessing.Process(target=worker1, args=(queue, 1))  
    p1.start()  

    p2 = multiprocessing.Process(target=worker2, args=(queue, 2))  
    p2.start()  
  
  
    time.sleep(3)
    queue.put('data1')
    queue.put('data2')
    queue.close()  
    p1.join()
    p2.join()
