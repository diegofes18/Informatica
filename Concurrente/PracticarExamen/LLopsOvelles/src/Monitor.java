import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 *
 * @author Diego
 */
 //Solo subo mi clase monitor porque lo otro es igual

public class Monitor {

    private final Lock lock = new ReentrantLock();
    private int llops, ovelles, esperant;
    private final Condition perEntrarOvelles,perSortirOvelles;

    public Monitor(){

        this.llops=0;
        this.ovelles=0;
        this.esperant=0;
        this.perEntrarOvelles=lock.newCondition();
        this.perSortirOvelles=lock.newCondition();

    }
    public void perSortirLlop(int id) {
        lock.lock();
        try {
            System.out.println("El llop: " + id + " vol sortir");
            llops--;
        }finally{
            lock.unlock();
        }
    }

    public void perEntrarLlop(int id) {

        lock.lock();
        try {
            System.out.println("El llop :"+id+" vol entrar");
            llops++;
            if(ovelles<2){
                System.out.println("Ñam");
            }
        }finally{
            lock.unlock();
        }
    }

    public void perEntrarOvella(int id) throws InterruptedException {
        lock.lock();
        try {
            System.out.println("    L'ovella : "+id+" vol entrar");
            esperant++;
            if(esperant>=2){
                perEntrarOvelles.signal();
            }else{
                while((esperant<2)&&(ovelles==0)){
                    perEntrarOvelles.await();
                }
            }
            ovelles++;
            esperant--;
        }finally{
            lock.unlock();
        }

    }

    synchronized void perSortirOvella(int id) throws InterruptedException {

        lock.lock();
        try {
            System.out.println("    L'ovella : "+id+" vol sortir");
            while(ovelles<2){
                perSortirOvelles.await();
            }
            ovelles--;
            perSortirOvelles.signal();
        }finally{
            lock.unlock();
        }

    }
}
