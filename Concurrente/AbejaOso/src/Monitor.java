public class Monitor  {
    static final int N_SIZE = 1000;
    static int currentPot = 0;
    public Monitor(){

    }

    synchronized public void Produce(){
        while (currentPot == N_SIZE){
            try {
                this.wait();
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }
        currentPot++;
        //System.out.println("CURRENT POT " + currentPot);
        this.notifyAll();
    }

    synchronized public void Consume(){
        while (currentPot < N_SIZE){
            try {
                this.wait();
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }
        System.out.println("ES BUIDA EL POT " + currentPot);
        currentPot = 0;
        notifyAll();
    }
}
