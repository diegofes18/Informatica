//CLASSE COTXE
public class Cotxe implements Runnable{
    //Numero per identificar el cotxe
    private int id;
    //Direcció cotxe
    //true -> NORD
    //false -> SUD
    private Direccio direccio;
    //Monitor per cridar als procediments
    private MonitorPont monitor;

    //CONSTRUCTOR
    public Cotxe(int id, MonitorPont monitor, Direccio direccio) {
        this.id = id;
        this.monitor = monitor;
        this.direccio=direccio;
    }

    @Override
    public void run() {
        try{
            //INICIAM RUTA
            Thread.sleep((long) (Math.random() * 200));
            //Si anam cap al nord demanamNord() i si no, demanamSud()
            switch (direccio){
                case NORD:
                    //DEMANAM ACCÉS AL PONT
                    System.out.println("El cotxe "+id+" està en ruta direcció NORD");
                    monitor.demanaNord(id);
                    break;
                case SUD:
                    //DEMANAM ACCÉS AL PONT
                    System.out.println("El cotxe "+id+" està en ruta direcció SUD");
                    monitor.demanaSud(id);
                    break;
            }
            //CONDUÏM PEL PONT
            Thread.sleep((long) (Math.random() * 1000));
            //SORTIM
            monitor.surt(id);
        }catch(InterruptedException ex){
            System.out.println(this.id + " EXEPCIÓ!" + ex.getMessage());
        }

    }
}
