package leetcode.Strings.Reverse_String_344;

class Solution {
    public String reverseString(String s) {
        char[] chars = s.toCharArray();
        int lastIdx = chars.length - 1;

        for (int i = 0; i < chars.length; i++) {

            if (i > lastIdx - i) {
                break;
            }

            char first = chars[i];
            char last = chars[lastIdx - i];

            chars[i] = last;
            chars[lastIdx - i] = first;
        }

        return String.valueOf(chars);
    }
}