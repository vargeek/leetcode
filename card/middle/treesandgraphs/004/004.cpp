#include <stdlib.h>
// Populating Next Right Pointers in Each Node
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/789/
// https://leetcode.com/submissions/detail/172428072/


/**
 * Definition for binary tree with next pointer.
 */
struct TreeLinkNode {
    int val;
    struct TreeLinkNode *left, *right, *next;
};

void connect(struct TreeLinkNode *root) {
    
    if (root == NULL) {
        return;
    }

    struct TreeLinkNode *head = root;
    
    while(head){
        
        if (head->left) {
            head->left->next = head->right;
        }
        
        if (head->right && head->next) {
            head->right->next = head->next->left;
        }
        head = head->next;
    }
    
    connect(root->left);
}
int main(int argc, char const *argv[])
{
    return 0;
}
