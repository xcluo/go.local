import multiprocessing  
import time  
  
def worker(num):  
    print "Worker", num  
    time.sleep(0.1)  
    return  
  
if __name__ == '__main__':  
    jobs = []  
    for i in range(5):  
        p = multiprocessing.Process(target=worker, args=(i, ))  
        jobs.append(p)  
        p.start() 
