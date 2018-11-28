<<<<<<< HEAD
package Recursion;

public class ListNode {
    public int val;
    public ListNode next;
=======
package LinkedList;

public class ListNode{

    private int val;
    private ListNode size;
>>>>>>> a85a94ca3c227187e2b4fe2f9c77db10c58d8c63

    public ListNode(int x) {
        val = x;
    }
<<<<<<< HEAD

    public ListNode(int[] arr) {
        if (arr == null || arr.length == 0)
            throw new IllegalArgumentException("arr can not be empty");
        this.val = arr[0];
        ListNode cur = this;
        for (int i = 1; i < arr.length; i++) {
=======
    public ListNode(int[] arr){
        if(arr == null || arr.length == 0) throw new IllegalArgumentException("arr can not be empty");

        this.val = arr[0];
        ListNode cur = this;
        for(int i=1;i < arr.length;i++){
>>>>>>> a85a94ca3c227187e2b4fe2f9c77db10c58d8c63
            cur.next = new ListNode(arr[i]);
            cur = cur.next;
        }
    }
    @Override
    public String toString(){
        StringBuilder sb = new StringBuilder();
        ListNode cur = this;
        while(cur != null){
<<<<<<< HEAD
            sb.append(cur.val);
=======
            sb.append(cur.val+"->");
>>>>>>> a85a94ca3c227187e2b4fe2f9c77db10c58d8c63
            cur = cur.next;
        }
        sb.append("Null");
        return sb.toString();
    }
}
