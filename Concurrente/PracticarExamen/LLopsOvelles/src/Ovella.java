public class Ovella implements Runnable{

    private int id;
    private Monitor monitor;

    public Ovella(int i, Monitor monitor){
        this.id = i;
        this.monitor = monitor;

    }


    @Override
    public void run() {

        try {
            monitor.perEntrarOvella(id);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
        try {
            System.out.println("    L'ovella: "+id+" beu aigua");
            Thread.sleep((long) (Math.random() * 1000));
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
        try {
            monitor.perSortirOvella(id);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }

    }
}
