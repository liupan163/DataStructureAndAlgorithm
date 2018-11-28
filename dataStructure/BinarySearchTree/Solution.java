package BinarySearchTree;

import java.util.LinkedList;
import java.util.List;
import java.util.Stack;

public class Solution {

    private class BST<E extends Comparable<E>> {
        private class Node {
            public E e;
            public Node left, right;

            public Node(E e) {
                this.e = e;
                left = null;
                right = null;
            }
        }

        private Node root;
        private int size;

        public BST() {
            root = null;
            size = 0;
        }

        public int size() {
            return size;
        }

        public boolean isEmpty() {
            return size == 0;
        }

        public void add(E e) {
            root = add(root, e);
        }

        public Node add(Node node, E e) {
            if (node == null) {
                size++;
                return new Node(e);
            }

            if (e.compareTo(node.e) < 0) {
                node.left = add(node.left, e);
            } else if (e.compareTo(node.e) > 0) {
                node.right = add(node.right, e);
            }
            return node;
        }
    }

    public int uniqueMorseRepresentations(String[] words) {
        String[] codes = {".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."};
        BST<String> bst = new BST<>();
        for (String word : words) {
            StringBuilder sb = new StringBuilder();
            for (int i = 0; i < word.length(); i++) {
                sb.append(codes[word.charAt(i) - 'a']);
            }
            bst.add(sb.toString());
        }
        return bst.size();
    }

    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }

    //Stack栈的方式遍历
    public List<Integer> preorderTraversal(TreeNode root) {
        List<Integer> res = new LinkedList<>();
        if (root == null) return res;

        Stack<TreeNode> stack = new Stack<>();
        stack.push(root);
        while (!stack.isEmpty()) {
            TreeNode cur = stack.pop();
            res.add(cur.val);

            if (cur.right != null) {
                stack.push(cur.right);
            }
            if (cur.left != null) {
                stack.push(cur.left);
            }
        }
        return res;
    }
}
