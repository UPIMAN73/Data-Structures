/**
 * @file queue.h
 * @author Joshua Calzadillas (jcalzadillas.job@gmail.com)
 * @brief 
 * @version 0.1
 * @date 2023-09-04
 * 
 * @copyright Copyright (c) 2023
 * 
 */
#ifndef QUEUE_H_
#define QUEUE_H_

#include <stdlib.h>

#define Queue struct queue_t
#define QueueDeclare(STRUCT_NAME, T) \
    STRUCT_NAME {\
        size_t size;\
        size_t capacity;\
        T bucket;\
    }

// 
QueueDeclare(Queue, int);

// Initiate the queue using a malloc
#define NewQueue() (Queue *) malloc(sizeof(Queue *))

// Peek at the first item
void * Peek(Queue * queue) {
    if (queue != NULL) {
        return (void *) queue;
    }
    else {
        return NULL;
    }
}

// Test function that demonstrates the use of the data-structure
void TestQueue() {
    Queue * queue = NewQueue();
    queue->capacity = 6;
    Peek(queue);
}

#endif