import threading
import time
 
sema = threading.Semaphore(2)
 

def worker():
    name = threading.currentThread().getName()
    print '%s acquire semaphore...' % name

    if sema.acquire():
        print '%s get semaphore' % threading.currentThread().getName()
        time.sleep(1)
        
        print '%s release semaphore' % threading.currentThread().getName()
        sema.release()
 

if __name__ == "__main__":
    for i in range(5):
        t = threading.Thread(target=worker)
        t.start()
