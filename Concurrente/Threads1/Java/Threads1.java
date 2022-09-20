package Threads1.Java;

public class Threads1 implements Runnable {
    static final int THREADS = 4;
    int id;

    public Threads1(int id) {
        this.id = id;
    }

    @Override
    public void run() {
        System.out.printf("Hi, I'm thread %d\n", id);
    }

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[THREADS];
        int i;
        
        for (i=0; i< THREADS; i++) {
            threads[i] = new Thread(new Threads1(i));
            threads[i].start();
        }

        for (i=0; i< THREADS; i++) {
            threads[i].join();
        }
        
    }
}