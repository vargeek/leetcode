#include <stdio.h>
#include <stdlib.h>

// 删除链表中的节点
// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/6/linked-list/41/
// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/description/


/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode {
    int val;
    struct ListNode *next;
};

void deleteNode(struct ListNode* node) {
    node->val = node->next->val;
    struct ListNode* next = node->next;
    node->next = node->next->next;
    free(next);
}

int main(int argc, char const *argv[])
{
    return 0;
}
