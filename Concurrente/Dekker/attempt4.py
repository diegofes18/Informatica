import threading 

#Constants
THREADS = int(input("Introdueix el numero de fils:"))
MAX_COUNT = 1000000

counter = 0
wantx = [False for i in range(THREADS)]


def count():
    
    global counter
    num = getTorn(threading.current_thread().name)-1
    
    #PREPROTOCOL
    wantx[num] = True

    while wantx[(num+1)%2]:
        wantx[num] = False
        wantx[num] = True
    
    #SECCIÓ CRITICA
    print("Thread {}".format(getTorn(threading.current_thread().name)))
    for i in range(MAX_COUNT//THREADS):
        counter += 1 

    #POSTPROTOCOL
    #print(f"El procés {num+1} ha comptat fins a {counter}")
    wantx[num] = False 

    

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
        
    
