package leetcode.arrays.Contains_Duplicate_217;

import java.util.Arrays;

class Solution {
    public boolean containsDuplicate(int[] nums) {

        if (nums.length < 2) return false;

        Arrays.sort(nums);

        for (int i = 0; i < nums.length - 1; i++) {
            if (nums[i] == nums[i + 1]) {
                return true;
            }
        }

        return false;
    }
}