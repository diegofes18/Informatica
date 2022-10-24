import threading
import collections
import time
import random

MAX_COUNT = 250
PRODUCERS = 10


#Monitor
class ProducerConsumer(object):
    H = 20
    def __init__(self):
        self.pot = 0
        self.mutex = threading.Lock()
        self.notFull = threading.Condition(self.mutex)
        self.canEat = threading.Condition(self.mutex)

    #Servir
    def append(self):
        with self.mutex:
            #if si no esta lleno x2
            while self.pot == self.H:
                self.notFull.wait()
            time.sleep(random.uniform(0.0500,1.0))
            #Metemos "data"
            self.pot += 1
             #Notificacion que ya hay "data"
            if self.pot == self.H:
                self.canEat.notify() 

    #Coger
    def take(self):
        with self.mutex:
            while self.pot != self.H:
                self.canEat.wait()
            self.pot = 0 #Sacar el valor mas a la izquierda de la cola
            time.sleep(random.uniform(0.0500,1.0))
            self.notFull.notify()
            

def abella(buffer,):
    id = threading.current_thread().name
    print("abeja quiere producir {}".format(id))
    for i in range(MAX_COUNT//PRODUCERS):
        buffer.append()
        print("        {} Produce".format(id))

def os(buffer):
    id = threading.current_thread().name
    print("oso quiere comer {}".format(id))
    for i in range(MAX_COUNT//ProducerConsumer.H):
        buffer.take()
        print("        {} Come".format(i))

    


def main():
    threads = []

    buffer = ProducerConsumer()
 
    c = threading.Thread(target=os, args=(buffer,))
    threads.append(c)
    for i in range(PRODUCERS):
        p = threading.Thread(target=abella, args=(buffer,))
        threads.append(p)
    
    # Start all threads
    for t in threads:
        t.start()

    # Wait for all threads to complete
    for t in threads:
        t.join()

    print("End")


if __name__ == "__main__":
    main()
