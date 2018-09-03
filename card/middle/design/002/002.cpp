#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

// Serialize and Deserialize Binary Tree
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/112/design/812/
// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/

typedef struct {
    unsigned char *data;
    size_t length;
    size_t capacity;
} Bytes;
Bytes newBytes(size_t capacity) {
    assert(capacity >= 0);
    return (Bytes){.data=(unsigned char*)malloc(capacity), .length=0, .capacity=capacity};
}
void growBytesIfNeed(Bytes *bytes, size_t addition) {
    assert(bytes != NULL);
    size_t capacity = bytes->capacity;
    while(bytes->length + addition > capacity){
        capacity = capacity * 2;
    }
    if (capacity > bytes->capacity) {
        unsigned char *data = (unsigned char*)malloc(capacity);
        assert(data != NULL);
        memcpy(data, bytes->data, bytes->length);
        bytes->capacity = capacity;
        free(bytes->data);
        bytes->data = data;        
    }
}
void appendBytes(Bytes *bytes, unsigned char*data, size_t len) {
    assert(bytes != NULL);
    
    if (data == NULL || len <= 0) {
        return;
    }
    growBytesIfNeed(bytes, len);
    memcpy(bytes->data+bytes->length, data, len);
    bytes->length += len;
}

void appendInt(Bytes *bytes,int intVal) {
    unsigned char data[sizeof(int)];
    size_t size = sizeof(int);
    for(size_t i = 0; i < size; i++)
    {
        data[i] = (unsigned char)(intVal & 0xff);
        intVal >>= 8;
    }
    appendBytes(bytes, data, size);
}

int getInt(unsigned char *data) {
    int val = 0;
    size_t size = sizeof(int);
    for(size_t i = 0; i <size; i++)
    {
        val <<= 8;
        val |= data[size-i-1];
    }
    return val;
}
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};
void _serialize(struct TreeNode* root, Bytes *bytes) {
    if (root == NULL) {
        return;
    }
    unsigned char flag = (root->left ? 0x01:0x00) | (root->right ? 0x02:0x00);
    appendInt(bytes, root->val);
    appendBytes(bytes, &flag, sizeof(flag));
    _serialize(root->left, bytes);
    _serialize(root->right, bytes);
}

/** Encodes a tree to a single string. */
char* serialize(struct TreeNode* root) {
    Bytes bytes = newBytes(1024);
    struct TreeNode head = {0, root, NULL};
    _serialize(&head, &bytes);
    return (char *)bytes.data;
}
struct TreeNode* _deserialize(unsigned char **data_ptr) {
    unsigned char *data = *data_ptr;
    int val = getInt(data);
    char flag = *((char*)data + sizeof(int));

    struct TreeNode* root = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    memset((void*)root, 0, sizeof(struct TreeNode));
    root->val = val;
    *data_ptr = (unsigned char*)data + sizeof(int) + 1;

    if (flag & 0x01) {
        root->left = _deserialize(data_ptr);
    }
    
    if (flag & 0x02) {
        root->right = _deserialize(data_ptr);
    }
    
    return root;
}
/** Decodes your encoded data to tree. */
struct TreeNode* deserialize(char* data) {
    struct TreeNode* head = _deserialize((unsigned char **)&data);
    return head != NULL ? head->left : NULL;
}

// Your functions will be called as such:
// char* data = serialize(root);
// deserialize(data);
void _printTreeNode(struct TreeNode *node) {
    
    if (node == NULL) {
        return;
    }
    printf("|%d", node->val);
    _printTreeNode(node->left);
    _printTreeNode(node->right);
}

void printTreeNode(struct TreeNode *node) {
    if (node == NULL) {
        printf("NULL\n");
    } else {
        _printTreeNode(node);
        printf("|\n");
    }
}

bool areTreeNodesEqual(struct TreeNode *node1, struct TreeNode *node2) {
    
    if ((node1 == NULL) != (node2 == NULL)) {
        return false;
    }
    
    if (node1 == NULL) {
        return true;
    }
    
    return areTreeNodesEqual(node1->left, node2->left) && areTreeNodesEqual(node1->right, node2->right);
}
void demo1() {
    struct TreeNode n1 = {1001,NULL, NULL};
    struct TreeNode n2 = {1002,NULL, NULL};
    struct TreeNode n3 = {1003,NULL, NULL};
    struct TreeNode n4 = {1004,NULL, NULL};
    struct TreeNode n5 = {1005,NULL, NULL};
    n1.left = &n2;
    n1.right = &n3;
    n3.left = &n4;
    n3.right = &n5;

    char *data = serialize(&n1);
    struct TreeNode* des = deserialize(data);

    printTreeNode(&n1);
    printTreeNode(des);
    assert(areTreeNodesEqual(&n1, des));
}

void demo2() {
    struct TreeNode *root = NULL;
    char *data = serialize(root);
    struct TreeNode *des = deserialize(data);
    assert(areTreeNodesEqual(root, des));
}

int main(int argc, char const *argv[])
{
    demo1();
    demo2();
    return 0;
}
