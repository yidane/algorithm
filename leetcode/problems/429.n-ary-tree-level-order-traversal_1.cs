/*
 * @lc app=leetcode id=429 lang=csharp
 *
 * [429] N-ary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/n-ary-tree-level-order-traversal/description/
 *
 * algorithms
 * Easy (58.48%)
 * Total Accepted:    25K
 * Total Submissions: 42.7K
 * Testcase Example:  '{"$id":"1","children":[{"$id":"2","children":[{"$id":"5","children":[],"val":5},{"$id":"6","children":[],"val":6}],"val":3},{"$id":"3","children":[],"val":2},{"$id":"4","children":[],"val":4}],"val":1}'
 *
 * Given an n-ary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
 * 
 * For example, given a 3-ary tree:
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * We should return its level order traversal:
 * 
 * 
 * [
 * ⁠    [1],
 * ⁠    [3,2,4],
 * ⁠    [5,6]
 * ]
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The depth of the tree is at most 1000.
 * The total number of nodes is at most 5000.
 * 
 * 
 */
/*
// Definition for a Node.
public class Node {
    public int val;
    public IList<Node> children;

    public Node(){}
    public Node(int _val,IList<Node> _children) {
        val = _val;
        children = _children;
}
*/
public class Solution {
    public IList<IList<int>> LevelOrder(Node root) {
        IList<IList<int>> result = new List<IList<int>>();

        if (root == null)
            return result;

        readNodeValue(new List<Node>(){root}, 0, ref result);

        return result;
    }

    private void readNodeValue(IList<Node> nodes, int level, ref IList<IList<int>> arr)
    {
        if (nodes == null || nodes.Count == 0)
            return;

        if (arr.Count <= level)
        {
            var newArr = new List<int>();
            arr.Add(newArr);
        }

        var currentArr = arr[level];
        foreach (var node in nodes)
        {
            currentArr.Add(node.val);
            readNodeValue(node.children, level + 1, ref arr);
        }
    }
}

