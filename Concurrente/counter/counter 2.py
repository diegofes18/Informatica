#! /usr/bin/env python
#
# An example of defining a thread as a class

import threading

MAX_COUNT = 10000000
THREADS = 2
n = 0

class CountThread(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)
    def run(self):
        global n
        print("Thread {}".format(threading.current_thread().name))
        for i in range (MAX_COUNT/THREADS):
            n += 1
            
def main():
    threads = []
    for i in range(THREADS):
        # Create new threads
        t = CountThread()
        threads.append(t)
        t.start() # start the thread
    # Wait for all threads to complete
    for t in threads:
        t.join()
    print("Counter value: {} Expected: {}\n".format(n, MAX_COUNT))
if __name__ == "__main__":
    main()