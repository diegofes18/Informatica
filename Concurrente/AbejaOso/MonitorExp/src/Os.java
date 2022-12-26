import javax.management.monitor.Monitor;

public class Os implements Runnable{

    private MonitorCondition monitor;
    public Os(MonitorCondition monitor) {
        this.monitor=monitor;
    }
    @Override
    public void run(){
        System.out.println("Bon dia som el os i vull menjar");
        try {
            monitor.demanaMel();
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }

        monitor.deixaPot();
        System.out.println( "ja estic ple");
    }
}
