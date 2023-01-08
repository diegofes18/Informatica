

public class Sortida implements Runnable{
    
    private int id;
    private Monitor monitor;

    public Sortida(int i, Monitor monitor){
        this.id = i;
        this.monitor = monitor;
    }

    @Override
    public void run() {
        boolean continua = true;
        while(continua){
            //delay
            monitor.surt(id);
            if(monitor.tancat()){
                continua = false;
            }

        }


    }
}