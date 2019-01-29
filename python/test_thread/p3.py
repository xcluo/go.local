import multiprocessing

  
def worker(x):  
    name = multiprocessing.current_process().name
    print 'Process %s x =%d'% (name, x)
  

pool = multiprocessing.Pool(processes=5)  
pool.map(worker, range(10))  
