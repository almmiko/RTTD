package leetcode.arrays.Move_Zeros;

class Solution {
    public void moveZeroes(int[] nums) {
        int idx = 0;
        int length = nums.length;

        for (int number : nums) {
            if (number != 0) {
                nums[idx++] = number;
            }
        }

        while (idx < length) {
            nums[idx++] = 0;
        }
    }
}
