#include <stdio.h>
#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};
struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB) {
    struct ListNode *a = headA;
    struct ListNode *b = headB;
    
    while(a != b){
        a = a == NULL ? headB : a->next;
        b = b == NULL ? headA : b->next;
    }

    return a;
}

int main(int argc, char const *argv[])
{

    return 0;
}

