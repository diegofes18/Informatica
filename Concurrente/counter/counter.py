#! /usr/bin/env python

import threading

THREADS = 2
MAX_COUNT = 10000000
n = 0

def count():
    global n
    print("Thread {}".format(threading.current_thread().name))
    for i in range(MAX_COUNT//THREADS):
        n += 1
    print(f"El procés {threading.current_thread().name} ha comptat fins a {n}")
def main():
    threads = []
    for i in range(THREADS):
        # Create new threads
        t = threading.Thread(target=count)
        threads.append(t)
        t.start() # start the thread
    # Wait for all threads to complete
    for t in threads:
        t.join()
    print("Counter value: {} Expected: {}\n".format(n, MAX_COUNT))

if __name__ == "__main__":
    main()
