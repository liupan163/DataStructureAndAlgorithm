package Recursion;

public class Sum {

    public static int sum(int[] arr) {
        return sum(arr, 0);
    }

    private static int sum(int[] arr, int index) {
        if (index != arr.length) {
            return arr[index] + sum(arr, index + 1);
        } else {
            return 0;
        }
    }

    public static void main(String[] args) {
        int[] nums = {1, 2, 3, 4, 5, 6, 7, 8, 9};
        System.out.println(sum(nums));
    }
}
