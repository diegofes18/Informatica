public class Llop implements Runnable{

    private int id;
    private Monitor monitor;

    public Llop(int i, Monitor monitor){
       this.id=i;
       this.monitor=monitor;
    }
    @Override
    public void run() {
        monitor.perEntrarLlop(id);
        //delay
        try {
            System.out.println("El llop: "+id+" beu aigua");
            Thread.sleep((long) (Math.random() * 1000));
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
        monitor.perSortirLlop(id);
    }
}
