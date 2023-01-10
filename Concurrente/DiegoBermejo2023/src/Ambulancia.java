//CLASSE AMBULÀNCIA
public class Ambulancia implements Runnable{
    //Monitor per cridar els procediments
    private MonitorPont monitor;
    //Número per identificar a l'ambulància
    private int id;
    //CONSTRUCTOR
    public Ambulancia(int id, MonitorPont monitor) {
        this.monitor = monitor;
        this.id = id;
    }
    @Override
    public void run() {
        try {
            //INICIAM RUTA
            Thread.sleep((long) (Math.random() *200));
            System.out.println("L'Ambulància "+id+" esta en ruta");
            //DEMANAM ACCÉS AL PONT
            monitor.demanaAmbulancia(id);
            //CONDUÏM PEL PONT
            Thread.sleep((long) (Math.random() * 10000));
            //SORTIM
            monitor.surt(id);
        } catch (InterruptedException ex) {
            System.out.println(this.id + " EXEPCIÓ!" + ex.getMessage());
        }

    }
}
