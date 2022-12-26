public class Main {

    static MonitorCondition monitor = new MonitorCondition();

    public static void main(String[] args) throws InterruptedException {

        Thread os = new Thread(new Os(monitor));
        os.start();

        Thread[] abelles = new Thread[4];
        for (int i = 0; i < abelles.length; i++) {
            abelles[i] = new Thread(new Abella(i,monitor));
            abelles[i].start();
        }
        os.join();
        for (int i = 0; i < abelles.length; i++) {
            abelles[i].join();
        }

    }
}