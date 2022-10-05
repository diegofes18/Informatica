from operator import truediv
import threading
import time

THREADS = 2
MAX_COUNT = 10000000
estat = [THREADS]
n = 0
start_time = time.time()

def thread():
    global estat
    global n
    max=MAX_COUNT//THREADS
    print("Hi, I'm the thread {}".format(threading.current_thread().name))
    id=int(threading.current_thread().name)
    nextId=(int(threading.current_thread().name)+1)%THREADS
    for i in range(max):
        estat[id]=1
        while estat[nextId]==1:
            estat[id]=0
            estat[id]=1
        n=n+1
        estat[id]=0

def main():
    global estat
    threads = []
    for i in range(THREADS):
        # Create new threads
        t = threading.Thread(target=thread)
        threads.append(t)
        estat.append(0)
        t.name=i
        t.start() # start the thread

    # Wait for all threads to complete
    for t in threads:
        t.join()

    print("Total time: --- %s seconds ---" % (time.time() - start_time))
    print("Expected "+str(MAX_COUNT)+", Real value "+str(n))

if __name__ == "__main__":
    main()
