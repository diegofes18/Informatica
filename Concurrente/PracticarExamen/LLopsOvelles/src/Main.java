
public class Main {

    static Monitor monitor = new Monitor();
    private static final int LLOPS = 4;
    private static final int OVELLES = 6;

    public static void main(String[] args) throws InterruptedException {

        Thread[] ovelles = new Thread[OVELLES];
        for (int i = 0; i < ovelles.length; i++) {
            ovelles[i] = new Thread(new Ovella(i, monitor));
            ovelles[i].start();
        }

        Thread[] llops = new Thread[LLOPS];
        for (int i = 0; i < llops.length; i++) {
            llops[i] = new Thread(new Llop(i, monitor));
            llops[i].start();
        }
        for (int i = 0; i < ovelles.length; i++) {
            ovelles[i].join();
        }

        for (int i = 0; i < llops.length; i++) {
            llops[i].join();
        }

    }
}