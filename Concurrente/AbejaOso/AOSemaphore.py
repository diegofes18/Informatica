import threading
import time
import random

H = 5 # capacitat màxima del pot
numAbelles = 20 # nombre d'abelles
pot = 0 # capacitat actual del pot
transportAbelles = 1 # Màxim d'unitats que pot transportar una abella al pot
semaforOs = threading.Semaphore(0)
semaforAbella = threading.Semaphore(1)

def os():
    global pot
    while True:
        print(f"L'os demana per entrar a menjar")
        semaforOs.acquire()
        print(f"L'os comença a menjar.")
        # temps que tarda l'ós en menjar
        time.sleep(random.uniform(0.005, 2.0))
        pot = 0
        print(f"L'ós ha acabat de menjar.")
        semaforAbella.release()

def abella():
    global pot
    threadId = threading.current_thread().name
    while True:
        semaforAbella.acquire()
        print(f"L'abella {threadId} entra al pot a ficar la Mel")
        time.sleep(random.uniform(0.005, 2.0))
        pot += 1
        if pot == H:
            print(f"L'abella {threadId} va a despertar a l'ós")
            time.sleep(random.uniform(0.005, 2.0))
            semaforOs.release()
        else:
            print(f"L'abella {threadId} ha depositat la Mel i dona pas a la següent abella")
            semaforAbella.release()
            time.sleep(random.uniform(0.005, 2.0))

def main():
    fils = []
    # Inicialitzam el consumidor --> OS
    o = threading.Thread(target=os)
    o.setName("0")
    fils.append(o)

    # Inicialitzam les abelles
    for i in range(numAbelles):
        a = threading.Thread(target=abella)
        a.setName(f"{i+1}")
        fils.append(a)

    # Executam tots els fils
    for t in fils:
        t.start()

    # Esperam a que acabin tots els processos
    for t in fils:
        t.join()

    print("End")

if __name__ == "__main__":
    main()
