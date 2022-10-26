import threading
import random
import time

MAX_COUNT=250
ABELLES=10

class Monitor(object):
    MAX=20
    mutex=threading.Lock()

    def __init__(self):
        self.data=0
        self.potProduir = threading.Condition(self.mutex)
        self.canEat = threading.Condition(self.mutex)

    def getCapacitat(self):
        return self.data
    
    def afegeix(self):
        with Monitor.mutex:
            #mentre pot sigui ple
            while self.data==Monitor.MAX:
                self.potProduir.wait()
            time.sleep(random.uniform(0.0500,1.0))
            self.data=self.data+1
            if(self.data==Monitor.MAX):
                self.canEat.notify()
            
    
    def buida(self):
        with Monitor.mutex:
            
            while(self.data!=Monitor.MAX):
                self.canEat.wait()
            self.data=0
            time.sleep(random.uniform(0.0500,1.0))
            self.potProduir.notify()
            

def abella(buffer,id):
    print("Hola soc l'abella {}".format(id))
    for i in range(MAX_COUNT//ABELLES):
        buffer.afegeix()
        print("L'abella {} ha produit mel: {}/{}".format(id,buffer.getCapacitat(),Monitor.MAX))
    

def bear(buffer):
    print("Hola soc l'ós")
    for i in range(MAX_COUNT//Monitor.MAX):
        buffer.buida()
        print("L'ós acaba de menjar {}".format(i))



def main():
    threads=[]

    pot=Monitor()

    os=threading.Thread(target=bear,args=(pot,))
    threads.append(os)

    for i in range(ABELLES):
        a=threading.Thread(target=abella,args=(pot,i))
        threads.append(a)

    # Start all threads
    for t in threads:
        t.start()
    
    # Wait for all threads to complete
    for t in threads:
        t.join()
    
    print("End")
    
    
if __name__ == "__main__":
    main()
