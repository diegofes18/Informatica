// Compilat amb $cc -pthread counter.c

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>


#define NUM_THREADS  2
#define MAX_COUNT 10000000

// Just used to send the index of the id
struct tdata {
    int tid;
};

volatile int counter = 0;

void *count(void *ptr) {
    long i, max = MAX_COUNT/NUM_THREADS;
    int tid = ((struct tdata *) ptr)->tid;
    for (i=0; i < max; i++) {
        counter += 1; // The global variable, i.e. the critical section
    }
    printf("End %d counter: %d\n", tid, counter);
    pthread_exit(NULL);
}

int main (int argc, char *argv[]) {
    pthread_t threads[NUM_THREADS];
    int rc, i;
    struct tdata id[NUM_THREADS];

    for(i=0; i<NUM_THREADS; i++){
        id[i].tid = i;
        rc = pthread_create(&threads[i], NULL, count, (void *) &id[i]);
    }

    for(i=0; i<NUM_THREADS; i++){
        pthread_join(threads[i], NULL);
    }

    printf("Counter value: %d Expected: %d\n", counter, MAX_COUNT);
    pthread_exit(NULL);
}
