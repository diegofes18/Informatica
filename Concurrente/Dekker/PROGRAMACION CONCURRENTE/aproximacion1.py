import threading
import time

THREADS = 2
MAX_COUNT = 1000000
torn=0
n = 0
start_time = time.time()

def thread():
    global torn
    global n
    max=MAX_COUNT//THREADS
    print("Hi, I'm the thread {}".format(threading.current_thread().name))
    id=int(threading.current_thread().name)
    nextId=(int(threading.current_thread().name)+1)%THREADS
    print(id)
    print(nextId)
    for i in range(max):
        while torn!=id:
            pass
        n=n+1
        torn=nextId
    print("He acabado"+str(id))

def main():
    threads = []

    for i in range(THREADS):
        # Create new threads
        t = threading.Thread(target=thread)
        threads.append(t)
        t.name=i
        t.start() # start the thread

    # Wait for all threads to complete
    for t in threads:
        t.join()

    print("Total time: --- %s seconds ---" % (time.time() - start_time))
    print("Expected "+str(MAX_COUNT)+", Real Value "+str(n))

if __name__ == "__main__":
    main()
