public class Abella implements Runnable{

    private int id;
    private MonitorCondition monitor;
    public Abella(int i, MonitorCondition monitor) {
        this.id = i;
        this.monitor = monitor;
    }

    @Override
    public void run() {
        for (int i = 0; i < 2; i++) {
            try {
                System.out.println(i+": soc l'abella "+id+" i vull produir");
                monitor.demanaProduir();
                System.out.println(i+":     soc l'abella "+id+" ja he produit");
                Thread.sleep((long) (Math.random() * 10000));
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }
    }
}
