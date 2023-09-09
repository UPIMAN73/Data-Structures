/**
 * @file linkedlist.h
 * @author Joshua Calzadillas (jcalzadillas.job@gmail.com)
 * @brief 
 * @version 0.1
 * @date 2023-09-04
 * 
 * @copyright Copyright (c) 2023
 * 
 */
#ifndef LINKEDLIST_H_
#define LINKEDLIST_H_

#include <stdlib.h>
#include <stdio.h>

// Name a list
#define LinkedList struct linkedlist_t

// Declare a list
#define LinkedListDeclare(T) \
    LinkedList {\
        T value;\
        LinkedList * previous;\
        LinkedList * next;\
    }

// Linked list
#define NewLinkedList() (LinkedList *) malloc(sizeof(LinkedList *))
LinkedListDeclare(int);

// Get the size of the list (based on a point in the list)
size_t GetSize(LinkedList * node) {
    if (node != NULL) {
        return GetSize(node->next) + 1;
    }
    else {
        return 0;
    }
}

// Test function that demonstrates the use of the datastructure
void TestLinkedList() {
    LinkedList * head = NewLinkedList();

    // Allocate and set 
    int i = 1;
    int targetValue = 10;
    LinkedList * currentNode = head;
    while (i <= targetValue) {
        currentNode->value = i;
        currentNode->next = NewLinkedList();
        currentNode->next->previous = currentNode;
        currentNode = currentNode->next;
        i++;
    }
}
#endif