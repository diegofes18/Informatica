


/**
 *
 * @author miquelmascarooliver
 */
class Monitor {

    static final int N = 5;
    private int plazas;
    private int acabats;
    
    public Monitor() {
        this.acabats=0;
        this.plazas = 0;
    }
    synchronized boolean tancat(){
        return acabats==3;
    }
    synchronized void surt(int id){
        plazas--;
        this.notifyAll();
    }
    synchronized void entra(int id) throws InterruptedException {
        while(plazas==5){
            this.wait();
        }
        plazas++;
    }
    synchronized void acabat(int id){
        acabats++;
    }




}
