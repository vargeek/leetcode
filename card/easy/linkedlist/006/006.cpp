#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

// Linked List Cycle
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/773/
// https://leetcode.com/problems/linked-list-cycle/description/

struct ListNode {
    int val;
    struct ListNode *next;
};
// bool hasCycle(struct ListNode *head) {
    
//     if (head == NULL) {
//         return false;
//     }
    
//     struct ListNode *node = head->next;
//     struct ListNode *prev = head;
//     struct ListNode *next = NULL;
//     while(node != NULL && node != head){
//         next = node->next;
//         node->next = prev;
//         prev = node;
//         node = next;
//     }
    
//     return node == head;
// }

bool hasCycle(struct ListNode *head) {

    if (head == NULL || head->next == NULL) {
        return false;
    }

    struct ListNode *slow = head;
    struct ListNode *fast = head->next;
    
    while(slow != fast ){
        
        if (fast == NULL || fast->next == NULL) {
            return false;
        }
        slow = slow->next;
        fast = fast->next->next;
    }    
        
    return true;
}

void test1() {
    struct ListNode n1 = {1, NULL};
    struct ListNode n2 = {2, &n1};
    struct ListNode n3 = {3, &n2};
    struct ListNode n4 = {4, &n3};
    n1.next = &n4;

    assert(hasCycle(&n4));
}

void test2() {
    struct ListNode n4 = {-4, NULL};
    struct ListNode n3 = {0, &n4};
    struct ListNode n2 = {2, &n3};
    struct ListNode n1 = {3, &n2};
    n4.next = &n2;

    assert(hasCycle(&n4));

}

int main(int argc, char const *argv[])
{
    test1();
    test2();

    return 0;
}
