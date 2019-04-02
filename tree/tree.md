
# tree

## BST树(Binary Search Tree)

### 概念
Binary-search tree（BST）是一颗二叉树，每个树上的节点都有<=1个父亲节点，ROOT节点没有父亲节点。同时每个树上的节点都有[0,2]个孩子节点(left child AND right child)。每个节点都包含有各自的KEY值以及相应的satellite data。其中KEY是几种BST基本操作的主要操作对象。

### BST的几种基本操作
* SEARCH
> 从根节点开始查找，若该节点值等于查找值，则返回；若小于查找值，则在该节点左孩子查找；否则查找右孩子。
* MINIMUM
> 一个BST的最左叶子节点的key值就是BST所有key值中最小的
* MAXIMUM
> 一个BST的最右叶子节点的key值就是BST所有key值中最小的
* PREDECESSOR
* SUCCESSOR
* INSERT
* DELETE

## AVL平衡二叉搜索树

### 概念
定义：平衡二叉树或为空树,或为如下性质的二叉排序树:  
  （1）左右子树深度之差的绝对值不超过1;  
  （2）左右子树仍然为平衡二叉树.  
平衡因子BF = 左子树深度 － 右子树深度.  
平衡二叉树每个结点的平衡因子只能是1，0，-1。  
若其绝对值超过1，则该二叉排序树就是不平衡的。

## RBT 红黑树
AVL是严格平衡树，因此在增加或者删除节点的时候，根据不同情况，旋转的次数比红黑树要多；  
红黑是弱平衡的，用非严格的平衡来换取增删节点时候旋转次数的降低；  
所以简单说，搜索的次数远远大于插入和删除，那么选择AVL树，如果搜索，插入删除次数几乎差不多，应该选择RB树。

红黑树上每个结点内含五个域，color，key，left，right，p。如果相应的指针域没有，则设为NIL。
一般的，红黑树，满足以下性质，即只有满足以下全部性质的树，我们才称之为红黑树：  
1）每个结点要么是红的，要么是黑的。  
2）根结点是黑的。  
3）每个叶结点，即空结点（NIL）是黑的。  
4）如果一个结点是红的，那么它的俩个儿子都是黑的。  
5）对每个结点，从该结点到其子孙结点的所有路径上包含相同数目的黑结点。

## B-树