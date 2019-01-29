import os, time, sys
pipe_name = 'pipe_test'
 

def client():
    pipeout = os.open(pipe_name, os.O_WRONLY)

    counter = 0
    while True:
        msg = b'Number %03d\n' % counter
        os.write(pipeout, msg)
        counter = (counter+1) % 5
        print("[Send]: ", msg)
        time.sleep(1)
 
 
if __name__ == "__main__":
    if not os.path.exists(pipe_name):
        os.mkfifo(pipe_name)  
    
    pid = os.fork()    
    client()
