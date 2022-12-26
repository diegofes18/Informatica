import java.lang.management.MonitorInfo;
import java.util.concurrent.TimeUnit;

public class Main  {
static final int N_Abelles = 10;
static final int Os = 1;
static final int TO_CONSUME = 5;
static final int TO_PRODUCE = (TO_CONSUME * 1000) /N_Abelles;

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[N_Abelles+Os];
        Monitor mn = new Monitor();
        int i;
        for (i=0; i< N_Abelles; i++) {
            threads[i] = new Thread(new Abella(mn,i,TO_PRODUCE));
            threads[i].start();
        }
        threads[i] = new Thread(new Os(mn,TO_CONSUME));
        threads[i].start();
        for (i=0; i< N_Abelles; i++) {
            threads[i].join();
        }
        threads[i].join();
    }
}

  class Abella implements Runnable {
    Monitor mn;
    int id;
    int produce;
    public Abella(Monitor m, int i,int p){
        mn = m;
        id = i;
        produce = p;
    }

      @Override
      public void run() {
          System.out.println("ABELLA N: " + id);
          for (int i = 0; i< produce; i++){
              try {
                  TimeUnit.SECONDS.sleep((long) 0.5);
              } catch (InterruptedException e) {
                  throw new RuntimeException(e);
              }
              mn.Produce();
              System.out.println("ABELLA " + id + " PRODUEIX:   " + i + "/" + produce);
          }
          System.out.println("ABELLA :" + id + " FINAL");
      }
  }

  class Os implements Runnable {
    Monitor mn;
    int consume;
    public Os(Monitor m,int c){
        mn = m;
        consume = c;
    }

      @Override
      public void run() {
          System.out.println("OS");
          for (int i = 0; i< consume; i++){
              mn.Consume();
              System.out.println("SOC L'OS I HE MENJAT " + i + " PICS");
          }
      }
  }