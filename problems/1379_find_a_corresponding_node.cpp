/**
 * Definition for a binary tree node.
 */
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
    TreeNode* ans;
    TreeNode* target;
public:
    void inorder(TreeNode* o, TreeNode* c) {
        if (o != NULL) {
            inorder(o->left, c->left);
            if (o == this->target) {
                ans = c;
            }
            inorder(o->right, c->right);
        }
    }
    
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
        this->target = target;
        inorder(original, cloned);
        return this->ans;
    }
};
