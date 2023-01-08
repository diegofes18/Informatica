public class Entrada implements Runnable{

    private int id;
    private Monitor monitor;
    private int entradas;

    public Entrada(int i, Monitor monitor){
        this.id=i;
        this.monitor=monitor;
        this.entradas = 4;
    }
    @Override
    public void run() {
        for (int i = 0; i < entradas; i++) {
            try {
                monitor.entra(id);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }
        System.out.println("He acabado CLOSE");
        monitor.acabat(id);

    }
}