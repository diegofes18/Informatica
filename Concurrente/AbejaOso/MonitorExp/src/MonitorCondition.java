public class MonitorCondition {

    private int size;
    private boolean menjant;
    private final int MAX = 3;

    public MonitorCondition(){
        this.size=0;
        this.menjant=false;
    }

    synchronized void demanaMel() throws InterruptedException {
        while(size<MAX){
            wait();
        }
        menjant = true;
        size = 0;
    }
    synchronized void deixaPot(){
        menjant=false;
        notifyAll();
    }

    synchronized void demanaProduir() throws InterruptedException {
        while(menjant || size>MAX){
            wait();
        }
        size++;
    }

}
