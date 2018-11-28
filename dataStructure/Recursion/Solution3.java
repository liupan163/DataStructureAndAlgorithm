package Recursion;

import java.util.List;

public class Solution3 {

    public ListNode removeElements(ListNode head, int val) {
        ListNode dummpyHead = new ListNode(-1);
        dummpyHead.next = head;

        ListNode prev = dummpyHead;
        while (prev.next != null) {
            if (prev.next.val == val) {
                ListNode delNode = prev.next;
                prev.next = delNode.next;
                delNode.next = null;
            } else {
                prev = prev.next;
            }
        }
        return dummpyHead.next;
    }

    //LinkedList
    public ListNode removeRElements(ListNode head, int val) {
        if (head == null) {
            return head;
        }
        head.next = removeRElements(head.next, val);
        return head.val == val ? head.next : head;
    }

    public ListNode removeElementsWithDepth(ListNode head, int val, int depth) {
        String depthString = generateDepthString(depth);
        if (head == null) {
            return head;
        }
        ListNode res = removeElementsWithDepth(head.next, val, depth + 1);
        ListNode ret;
        if (head.val == val) {
            ret = res;
        } else {
            head.next = res;
            ret = head;
        }
        return ret;
    }

    private String generateDepthString(int depth) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < depth; i++) {
            sb.append("--");
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        int[] nums = {1, 2, 3, 4, 5, 6, 7, 8};
        ListNode head = new ListNode(nums);
        System.out.println(head);

        ListNode res = new Solution3().removeElements(head, 6);
        System.out.println(res);
    }
}
