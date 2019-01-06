package leetcode.unstructured.Single_Number_136;

//Topic hashmap, bitmanipulation

import java.util.HashMap;
import java.util.Map;

class Solution {
    public int singleNumber(int[] nums) {

        //using xor
        /*
            int result = 0;
            for (int i = 0; i<n; i++)
            {
                result ^=A[i];
            }
        */

        Map<Integer, Integer> arrayMap = new HashMap<>();
        final int[] res = {0};

        for (int number : nums) {
            arrayMap.put(number, arrayMap.getOrDefault(number, 0) + 1);
        }

        arrayMap.forEach((k, v) -> {
            if (v == 1) {
                res[0] = k;
            }
        });

        return res[0];

    }
}
