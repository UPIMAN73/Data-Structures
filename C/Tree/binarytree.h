/**
 * @file binarytree.h
 * @author Joshua Calzadillas (jcalzadillas.job@gmail.com)
 * @brief 
 * @version 0.1
 * @date 2023-09-04
 * 
 * @copyright Copyright (c) 2023
 * 
 */
#ifndef BINARYTREE_H_
#define BINARYTREE_H_

#include <stdlib.h>
#include <stdio.h>

// Makes life soo much easier when defining the environment
#define BTNode struct node

/**
 * @brief Binary Tree Structure
 * Used to define a binary tree structure based on an input type T.
 * That type is the type that will be presented when returning the value
 * of the binary tree based on it's datatype.
 */
#define BinaryTree(T) \
    BTNode {\
        BTNode* left;\
        BTNode* right;\
        T value;\
    }

// Define Binary Tree Size Type
#define BTSize_T unsigned long
#define BTMax(x, y) (y * (x < y) + x * (x > y)) // branchless way of saying (which one is greatest)

// Used to define any type based on generics
/**
 * @brief Construct a new Binary Tree object
 * The object listed bellow is required for all code base, to use this definition.
 * It essentially defines a struct around a type and the type is an input argument.
 * This makes it soo much easier for me to define whatever datatype I want to.
 */
BinaryTree(int);    // This is required 

// This is not required but it makes it a whole lot easier when defining new objects
#define BTNew (BTNode *) malloc(sizeof(BTNode *))

/**
 * @brief Get the Depth of the Binary Tree
 * 
 * @param node 
 * @return BTSize_T 
 */
BTSize_T GetDepth(BTNode * node) {
    if (node != NULL) {
        // Depth first search
        BTSize_T left = GetDepth(node->left);
        BTSize_T right = GetDepth(node->right);

        // Recursive iteration of the max value
        return BTMax(left, right) + 1;
    }
    else {
        return 0;
    }
}

// Test function that demonstrates the use of the datastructure
void TestBinaryTree() {
    // Defining the root node
    BTNode * root = BTNew;
    root->left = BTNew;

    // setting values for each of the nodes
    root->value = 1;
    root->left->value = 2;

    // Printing out the values
    printf("Root Value: %d\n", root->value);
    printf("Left Value: %d\n", root->left->value);

    // Get the depth of the tree
    printf("Depth of the tree: %ld\n", GetDepth(root));

    // Garbage collection (Will be a method for this later on)
    root->value = 0;
    root->left->value = 0;
    free(root->left);
    free(root);

    // Get out of here
    return;
}
#endif