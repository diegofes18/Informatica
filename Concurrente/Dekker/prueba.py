import threading 

#Constants
THREADS = 2
MAX_COUNT = 1000000

counter = 0
torn = 1 

def count():

    global counter
    #SECCIÓ CRITICA
    print(" {}".format(getTorn(threading.current_thread().name)))
    for i in range(MAX_COUNT//THREADS):
        counter += 1 

def getTorn(str):
    return str[7]


def main():
    threads = []
    for i in range(THREADS):
        t = threading.Thread(target=count)
        threads.append(t) 
        t.start()

    for t in threads:
        t.join()
    print("Counter value: {} Expected: {}\n".format(counter,MAX_COUNT))

if __name__ == "__main__":
    main()
        
    