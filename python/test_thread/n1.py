import threading
import time
 

product = None
cond = threading.Condition()

 
def produce():
    global product

    if cond.acquire():
        while True:
            if product is None:
                print 'Produce...'
                product = 'anything'

                cond.notify()

            cond.wait()
            time.sleep(2)
 

def consume():
    global product
    
    if cond.acquire():
        while True:
            if product is not None:
                print 'Consume...'
                product = None

                cond.notify()

            cond.wait()
            time.sleep(2)
 

if __name__ == "__main__":
    t1 = threading.Thread(target=produce)
    t2 = threading.Thread(target=consume)
    t2.start()
    t1.start()
