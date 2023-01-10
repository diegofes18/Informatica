import java.util.Random;

public class Main {
    //Instància del monitor per pasar per paràmetre al constructor dels fils
    static MonitorPont monitor = new MonitorPont();
    //Nombre total de cotxes
    private static final int COTXES = 5;
    //Cotxes que van al NORD
    private static int nord;
    //Cotxes que van al SUD
    private static int sud;


    public static void main(String[] args) throws InterruptedException {
        //Inicialitzam els nombres dels cotxes
        //Com a mínim hem de tenir-ne un de cada
        nord=getRandomNumber(1,4);
        sud=COTXES-nord; //Aquesta variable no la utilitzam pero serveix per comprobar

        //Fil per al procés ambulància
        Thread a = new Thread(new Ambulancia(101,monitor));
        a.start();

        //Array de fils per als cotxes
        Thread[] cotxes = new Thread[COTXES];

        //Inicialitzam els cotxes cap al NORD
        for (int i = 0; i < nord; i++) {
            cotxes[i] = new Thread(new Cotxe(i, monitor,Direccio.NORD));
            cotxes[i].start();
        }
        //Inicialitzam els fils cap al SUD
        for (int i = nord; i < COTXES; i++) {
            cotxes[i] = new Thread(new Cotxe(i, monitor,Direccio.SUD));
            cotxes[i].start();
        }
        //Ambulància
        a.join();
        //Cotxes
        for (int i = 0; i < COTXES; i++) {
            cotxes[i].join();
        }
    }
    //Funció auxiliar per treure números aleatoris en un rang
    public static int getRandomNumber(int min, int max) {
        return (int) ((Math.random() * (max - min)) + min);
    }
}
