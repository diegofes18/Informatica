/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package peterson;

import java.time.LocalTime;
import java.util.Random;

/**
 *
 * @author diegofes
 */
public class Peterson implements Runnable {

    static final int THREADS = 2;
    static final int MAX_COUNT = 20;
    static volatile int[] want = new int[THREADS];
    static volatile int last = 0;
    static volatile int n = 0;
    int id;
   
    private float avgTime;

    public Peterson(int id) {
        this.id = id;
        this.avgTime = 0;
    }
   

    @Override
    public void run() {
        Random rand = new Random();
        long time = 0;
        int max = MAX_COUNT / THREADS;
        int altreProces = (id + 1) % THREADS;
        System.out.println("Entrada " + id);
        for (int i = 0; i < max; i++) {
            // Peterson
             try {
                Thread.sleep(rand.nextInt(5000));
            } catch (InterruptedException ex) {
                System.out.println("Error durant la interrupció aleatòria.");
            }
            want[id] = 1;
            last = id;
            while (want[altreProces] == 1 && last == id) {
            }
            System.out.println("Porta: "+id+" "+(i+1)+" entrades de : "+(n+1)+ " Temps: "+LocalTime.now());

 //           System.out.println("Index: " + i + " Contador " + n + " Fil: " + id + " Altre procés: " + altreProces);
            n += 1; //SC
            if(i>0){
                avgTime += (System.currentTimeMillis()-time);
            }
            time = System.currentTimeMillis();
            want[id] = 0;
            
        }
        avgTime = (avgTime/max)/1000;
    }

    public float getAvgTime() {
        return avgTime;
    }
    

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[THREADS];
        Peterson[] counter = new Peterson[THREADS];
        want[0] = 0;
        want[1] = 0;
        
        int i;
        for (i = 0; i < THREADS; i++) {
            counter[i] = new Peterson(i);
            threads[i] = new Thread(counter[i]);
            threads[i].start();
        }
        for (i = 0; i < THREADS; i++) {
            threads[i].join();
        }
        System.out.println("Entrades totals: "+n);
        for (i = 0;  i< THREADS; i++) {
            System.out.println("Temps mig porta "+i+": "+counter[i].getAvgTime()+" segons");
            
        }
    }
}