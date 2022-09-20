import threading 

#Constants
THREADS = 2
MAX_COUNT = 1000000

counter = 0
wantx = [False for i in range(THREADS)]


def count():
    
    global counter
    num = getTorn(threading.current_thread().name)-1
    
    #PREPROTOCOL
    wantx[(num+1)%2] = True

    while wantx[num]:
        pass
    
    #SECCIÓ CRITICA
    print("Thread {}".format(getTorn(threading.current_thread().name)))
    for i in range(MAX_COUNT//THREADS):
        counter += 1 

    #POSTPROTOCOL
    print(f"El procés {num+1} ha comptat fins a {counter}")
    wantx[(num+1)%2] = False 

    

def getTorn(str):
    return int(str[7])

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
        
    
