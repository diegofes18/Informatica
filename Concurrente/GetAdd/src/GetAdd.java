

import java.util.concurrent.atomic.AtomicInteger;

public class GetAdd implements Runnable{
    static final int THREADS = 4;
    static final int MAX_COUNT = 1000000;
    static AtomicInteger mutex;
    static AtomicInteger torn;
    static volatile int n = 0;
    int id;

    public GetAdd(int id) {
        this.id = id;
    }

    @Override
    public void run() {
        int fi = MAX_COUNT / THREADS;
        int current;
        System.out.println("Thread " + id);
        for (int i = 0; i < fi; i++) {
            // Get&Set
            current = mutex.getAndAdd(1);
            while (current!=torn.get()){

            }
            n += 1; //SC
            torn.getAndAdd(1);
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[THREADS];
        mutex = new AtomicInteger(0);
        torn = new AtomicInteger(0);
        int i;
        for (i = 0; i < THREADS; i++) {
            threads[i] = new Thread(new GetAdd(i));
            threads[i].start();
        }
        for (i = 0; i < THREADS; i++) {
            threads[i].join();
        }
        float error = (MAX_COUNT - n) / (float) MAX_COUNT * 100;
        System.out.printf("Counter value: %d Expected: %d Error: %3.6f%%\n", n, MAX_COUNT, error);
    }
}

