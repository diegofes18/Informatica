//CLASSE MONITOR
public class MonitorPont {
    //Variables amb els cotxes esperant a cada direcció
    private int esperaNord, esperaSud;
    //Variables booleanes
    //  esperaAmbulancia: Hi ha una ambulancia esperant
    //  ocupat: el pont está ocupat
    private boolean esperaAmbulancia, ocupat;

    //CONSTRUCOR
    //  Inicialitzam les variables d'estat del monitor
    public MonitorPont(){
        this.esperaNord=0;
        this.esperaSud=0;
        this.esperaAmbulancia=false;
        this.ocupat=false;
    }

    synchronized void demanaNord(int id) throws InterruptedException {
        //Esperam al pont NORD
        esperaNord++;
        System.out.println("    El cotxe "+id+" espera a l'entrada NORD, esperen al NORD "+esperaNord);
        /*Ens bloquetjam si:
            -El pont está ocupat
            -L'altre accés té més cotxes esperant
            -Hi ha una ambuláncia esperant
        */
        while(ocupat || (esperaSud>esperaNord) || esperaAmbulancia){
            this.wait();
        }
        //Ocupam el pont i deixam d'esperar
        ocupat=true;
        esperaNord--;
        System.out.println("    El cotxe "+id+" es al pont, esperen al NORD "+esperaNord);
    }
    synchronized void demanaSud(int id) throws InterruptedException {
        //Esperam al pont SUD
        esperaSud++;
        System.out.println("    El cotxe "+id+" espera a l'entrada SUD, esperen al SUD "+esperaSud);
        /*Ens bloquetjam si:
            -El pont está ocupat
            -L'altre accés té més cotxes esperant
            -Hi ha una ambuláncia esperant
        */
        while(ocupat || (esperaNord>=esperaSud) || esperaAmbulancia){
            this.wait();
        }
        //Ocupam el pont i deixam d'esperar
        ocupat=true;
        esperaSud--;
        System.out.println("    El cotxe "+id+" es al pont, esperen al SUD "+esperaSud);
    }
    synchronized void demanaAmbulancia(int id) throws InterruptedException {
        //Esperam
        System.out.println("++++++Ambulància "+id+" espera per entrar");
        esperaAmbulancia=true;
        //L'ambulància només espera a que el pont sigui buit
        while(ocupat){
            this.wait();
        }
        //deixam d'esperar i ocupam el pont
        esperaAmbulancia=false;
        ocupat=true;
        System.out.println("++++++Ambulància "+id+" és al pont");
    }

    synchronized void surt(int id) {
        //DEIXAM D'OCUPAR EL PONT
        ocupat=false;
        System.out.println("----->El vehicle "+id+" surt del pont");
        this.notifyAll();
    }

}
