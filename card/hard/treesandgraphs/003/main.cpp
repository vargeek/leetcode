#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

// Lowest Common Ancestor of a Binary Tree
// https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/844/


// Definition for a binary tree node.
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

typedef struct {
    struct TreeNode *node;
    int p;
    int q;
} Context;

int search(struct TreeNode *root, Context *ctx) {
    
    if (root == NULL) {
        return 0;
    }
    
    int left = search(root->left, ctx);
    if (ctx->node != NULL) {
        return 2;
    }

    int mid = (root->val == ctx->p || root->val == ctx->q) ? 1 : 0;
    if (left + mid == 2) {
        ctx->node = root;
        return 2;
    }

    int right = search(root->right, ctx);
    if (ctx->node != NULL) {
        return 2;
    }
    
    if (left + mid + right == 2) {
        ctx->node = root;
    }
    return left+mid+right;
}


struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {
    Context ctx = {NULL, p->val, q->val};
    search(root, &ctx);
    return ctx.node;
}


void test1() {
    struct TreeNode n1 = {3, NULL, NULL};
    struct TreeNode n2 = {5, NULL, NULL};
    struct TreeNode n3 = {1, NULL, NULL};
    struct TreeNode n4 = {6, NULL, NULL};
    struct TreeNode n5 = {2, NULL, NULL};
    struct TreeNode n6 = {0, NULL, NULL};
    struct TreeNode n7 = {8, NULL, NULL};
    struct TreeNode n8 = {7, NULL, NULL};
    struct TreeNode n9 = {4, NULL, NULL};

    n1.left = &n2;
    n1.right = &n3;
    n2.left = &n4;
    n2.right = &n5;
    n3.left = &n6;
    n3.right = &n7;
    n5.left = &n8;
    n5.right = &n9;

    struct TreeNode *got = lowestCommonAncestor(&n1, &n2, &n3);
    struct TreeNode *want = &n1;
    if (!(got != NULL && got == want)) {
        printf("got: %d, want: %d\n", got ? got->val : -1, want ? want->val : -1);
        assert(false);
    } else {
        printf("PASS!\n");
    }
}
void test2() {
    struct TreeNode n1 = {3, NULL, NULL};
    struct TreeNode n2 = {5, NULL, NULL};
    struct TreeNode n3 = {1, NULL, NULL};
    struct TreeNode n4 = {6, NULL, NULL};
    struct TreeNode n5 = {2, NULL, NULL};
    struct TreeNode n6 = {0, NULL, NULL};
    struct TreeNode n7 = {8, NULL, NULL};
    struct TreeNode n8 = {7, NULL, NULL};
    struct TreeNode n9 = {4, NULL, NULL};

    n1.left = &n2;
    n1.right = &n3;
    n2.left = &n4;
    n2.right = &n5;
    n3.left = &n6;
    n3.right = &n7;
    n5.left = &n8;
    n5.right = &n9;

    struct TreeNode *got = lowestCommonAncestor(&n1, &n2, &n9);
    struct TreeNode *want = &n2;
    
    if (!(got != NULL && got == want)) {
        printf("got: %d, want: %d\n", got ? got->val : -1, want ? want->val : -1);
        assert(false);
    } else {
        printf("PASS!\n");
    }
    
}
int main(int argc, char const *argv[])
{
    test1();
    test2();
    return 0;
}
