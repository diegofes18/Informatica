
public class Main {

    static Monitor monitor = new Monitor();
    private static final int ENTRADES = 3;
    private static final int SORTIDES = 2;

    public static void main(String[] args) throws InterruptedException {

        Thread[] entrades = new Thread[ENTRADES];

        for (int i = 0; i < entrades.length; i++) {
            entrades[i] = new Thread(new Entrada(i, monitor));
            entrades[i].start();
        }

        Thread[] sortides = new Thread[SORTIDES];
        for (int i = 0; i < sortides.length; i++) {
            sortides[i] = new Thread(new Sortida(i, monitor));
            sortides[i].start();
        }
        for (int i = 0; i < entrades.length; i++) {
            entrades[i].join();
        }

        for (int i = 0; i < sortides.length; i++) {
            sortides[i].join();
        }

    }
}