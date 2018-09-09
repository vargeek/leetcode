// Copy List with Random Pointer
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/117/linked-list/841

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void printNode(struct RandomListNode *node);
/**
 * Definition for singly-linked list with a random pointer.
 * struct RandomListNode {
 *     int label;
 *     struct RandomListNode *next;
 *     struct RandomListNode *random;
 * };
 */
struct RandomListNode {
    int label;
    struct RandomListNode *next;
    struct RandomListNode *random;
};
struct RandomListNode *copyRandomList1(struct RandomListNode *head) {
    if (head == NULL) {
        return NULL;
    }
    struct RandomListNode node = {0, NULL, NULL};
    struct RandomListNode *ptr = &node;
    struct RandomListNode *tail = head;

    struct RandomListNode node2 = {0, NULL, NULL};
    struct RandomListNode *ptr2 = &node2;
    
    
    while(tail != NULL){
        size_t size = sizeof(struct RandomListNode);

        ptr2->next = (struct RandomListNode *)malloc(size);
        ptr2 = ptr2->next;
        memset(ptr2, 0, size);
        ptr2->random = tail->random;

        ptr->next = (struct RandomListNode *)malloc(size);
        ptr = ptr->next;
        memset(ptr, 0, size);
        // ptr->label = tail->label + 100;
        ptr->label = tail->label;
        ptr->random = tail->random;
        tail->random = ptr;
        tail = tail->next;
    }

    tail = head;
    ptr = node.next;
    while(tail != NULL){
        
        if (ptr->random != NULL) {
            ptr->random = ptr->random->random;
        }

        ptr = ptr->next;
        tail = tail->next;
    }

    tail = head;
    ptr2 = node2.next;
    while(tail != NULL){
        tail->random = ptr2->random;
        ptr2 = ptr2->next;
        tail = tail->next;
    }
    
    return node.next;
}

struct RandomListNode *copyRandomList(struct RandomListNode *head) {
    
    if (head == NULL) {
        return NULL;
    }
    
    struct RandomListNode *tail = head;
    struct RandomListNode *copy = NULL;
    
    // 将每个copy的节点插到原节点后面
    while(tail != NULL){
        copy = (struct RandomListNode*)malloc(sizeof(struct RandomListNode));
        memset(copy, 0, sizeof(struct RandomListNode));
        copy->label = tail->label;
        copy->next = tail->next;
        tail->next = copy;
        tail = copy->next;
    }
    
    tail = head;
    while(tail != NULL){
        copy = tail->next;        
        if (tail->random != NULL) {
            copy->random = tail->random->next;
        }
        tail = copy->next;
    }

    struct RandomListNode *result = head->next;
    tail = head;
    while(tail != NULL){
        copy = tail->next;
        tail->next = copy->next;
        tail = tail->next;
        copy->next = tail != NULL ? tail->next : NULL;
    }
    
    return result;
}

void printNode(struct RandomListNode *node) {
    
    while(node != NULL){
        
        if (node->random != NULL) {
            printf("[%d~>%d]->", node->label, node->random->label);
        } else {
            printf("%d->", node->label);
        }
        node = node->next;
    }
    printf("NULL\n");
}

void test1() {
    int count = 10;
    struct RandomListNode nodes[count];
    memset(nodes, 0, sizeof(nodes));
    
    for(size_t i = 0; i < count; i++)
    {
        nodes[i].label = i;
        
        if (i > 0) {
            nodes[i-1].next = &nodes[i];
        }
    }
    nodes[3].random = &nodes[1];
    nodes[5].random = &nodes[8];
    struct RandomListNode *copied = copyRandomList(&nodes[0]);
    printNode(&nodes[0]);
    printNode(copied);
}
int main(int argc, char const *argv[])
{
    
    test1();

    return 0;
}
