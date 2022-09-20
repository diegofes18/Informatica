package counter;

public class Counter implements Runnable {
//public class Counter extends Thread {
    static final int THREADS = 2;
    static final int MAX_COUNT = 10000000;
    static volatile int n = 0;
//    static  int n = 0; // El compilador manté una còpia local a cada fil
    int id;

    public Counter(int id) {
        this.id = id;
    }

    @Override
    public void run() {
        int max = MAX_COUNT/THREADS;
        System.out.printf("Thread %d\n", id);
        for (int i =0; i < max; i++) {
            n += 1;
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[THREADS];
        int i;

        for (i=0; i< THREADS; i++) {
            threads[i] = new Thread(new Counter(i));
            threads[i].start();
        }

        for (i=0; i< THREADS; i++) {
            threads[i].join();
        }
        System.out.printf("Counter value: %d Expected: %d\n", n, MAX_COUNT);
    }
}
