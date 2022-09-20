import threading, time

THREADS = 4

def thread():
    time.sleep(0.5)
    print("Hi, I'm the thread {}".format(threading.current_thread().name))

def main():
    threads = []

    for i in range(THREADS):
        # Create new threads
        t = threading.Thread(target=thread)
        threads.append(t)
        t.start() # start the thread

    # Wait for all threads to complete
    for t in threads:
        t.join()

    print("End")

if __name__ == "__main__":
    main()