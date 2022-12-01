import threading
import time
import collections

MAX = 10

TO_PRODUCE = 3
PRODUCERS = 2

CONSUMERS = 5
COMPRAS = 3

class Maquina(object):
    def __init__(self, max): #constructor
        #deque doble cua permet treure i afegir left right

        self.size = 5
        self.mutex = threading.Lock()
        self.notFull = threading.Semaphore(max)
        self.notEmpty = threading.Semaphore(0)

    def append(self):
        self.notFull.acquire()
        with self.mutex:
            self.size += 3
        self.notEmpty.release()

    def take(self):
        self.notEmpty.acquire()
        with self.mutex:
            self.size -= 1
        self.notFull.release()
        return self.size


def producer(maquina):
    id = threading.current_thread().name
    print("Reponedor {}".format(id))
    for i in range(TO_PRODUCE):
        
        maquina.append()
        print("        {} HE REPUESTO: {}".format(id, maquina.size))

def consumer(maquina):
    id = threading.current_thread().name
    print("COMPRO {}".format(id))
    for i in range(COMPRAS):
        size = maquina.take()
        print("{} CONSUMEIX: {}".format(id, size))


def main():
    threads = []

    maquina = Maquina(MAX)
    for i in range(CONSUMERS):
        c = threading.Thread(target=consumer, args=(maquina,))
        threads.append(c)

    for i in range(PRODUCERS):
        p = threading.Thread(target=producer, args=(maquina,))
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
