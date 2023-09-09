/**
 * @file list.h
 * @author Joshua Calzadillas (jcalzadillas.job@gmail.com)
 * @brief 
 * @version 0.1
 * @date 2023-09-04
 * 
 * @copyright Copyright (c) 2023
 * 
 */
#ifndef LIST_H_
#define LIST_H_

#include <stdlib.h>
#include <stdio.h>

// Declare a list
#define ListDeclare(LISTNAME, T) LISTNAME { T value; LISTNAME * next; }

// New List Definition
#define NewObj(objName) (objName *) malloc(sizeof(objName *))

// Name a list type & newList declaration
#define List struct list_t
#define ListF struct list_float_t

// Cannot redefine (yet)
ListDeclare(List, int);
// ListDeclare(ListF, float);

// New list declaration as a statement
#define NewList NewObj(List)

// Definition of printing a list
#define PrintList(IteratorListObj) size_t CurrentListObjNumber = 1;\
    puts("\nList Items:");\
    while (IteratorListObj != NULL) {\
        printf("\t%ld\tValue: %d\n", CurrentListObjNumber, IteratorListObj->value);\
        IteratorListObj = IteratorListObj->next;\
        CurrentListObjNumber++;\
    }\
    CurrentListObjNumber = 0;

// Custom Garbage collection
#ifdef LIST_GBC_TYPE
void ListGarbageCollector(void * list) {
    IteratorListObj = list;
    void * PreviousListObj = NULL;
    // Garbage collection
    while (1) {
        // Free the previous object
        if (PreviousListObj != NULL) {
            free((List *) PreviousListObj);
        }

        // Previous List Object Definition
        PreviousListObj = ((void *) IteratorListObj);

        // Iterator list object continuation
        if (IteratorListObj->next != NULL) {
            IteratorListObj = IteratorListObj->next;
        }
        else {
            free(IteratorListObj);
            IteratorListObj = NULL;
            PreviousListObj = NULL;
        }

        // Break Iterator
        if ((void *)IteratorListObj != NULL && (void *)PreviousListObj != NULL) {
            continue;
        }
        else if (IteratorListObj == NULL && PreviousListObj != NULL) {
            continue;
        }
        else {
            break;
        }
    }
}
#else
void ListGarbageCollector(void * list) {
    free(list);
    return;
}
#endif

// Test function that demonstrates the use of the datastructure
void TestList() {
    // 
    List * list = NewList;

    // Setup list int
    int i = 1;
    int targetValue = 10;
    List * currentNode = list;
    while (i <= targetValue) {
        currentNode->value = i;
        i++;
        if (i <= targetValue) {
            currentNode->next = NewList;
            currentNode = currentNode->next;
        }
    }

    PrintList(list)

    // Garbage collection
    ListGarbageCollector(list);
    puts("Cleaned up List objects");
}

#endif