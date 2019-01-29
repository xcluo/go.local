import os, time, sys
pipe_name = 'pipe_test'
 

def server( ):
    pipein = open(pipe_name, 'r')
    while True:
        line = pipein.readline()[:-1]
        print('[Recieved] Parent %d got "%s" at %s' % (os.getpid(), line, time.time( )))
        time.sleep(1)


if __name__ == "__main__":
    if not os.path.exists(pipe_name):
        os.mkfifo(pipe_name)  
    
    pid = os.fork()    
    server()
