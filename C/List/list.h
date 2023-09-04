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

// Name a list
#define List struct list_t

// Declare a list
#define ListDeclare(T) \
    List {\
        T value;\
        List * next;\
    }

#define NewList() (List *) malloc(sizeof(List *))

// Cannot redefine (yet)
ListDeclare(int);
// ListDeclare(float);


void TestList() {
    // 
    List * listInt = NewList();
    // List * listFloat = NewList();

    // Setup list int
    listInt->next = NewList();
    listInt->value = (int) 1;

    // Define list float value
    // listFloat->value = 3.1415927f;

    // Print out values
    printf("List Value: %d\n", listInt->value);
    // printf("List Value: %0.2f\n", listFloat->value);

    // Garbage collection
    listInt->value = 0;
    free(listInt->next);
    free(listInt);
    // free(listFloat);
}

#endif