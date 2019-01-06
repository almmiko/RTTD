package leetcode.linkedList.Reverse_Linked_List_206;

/*
*
* Reverse a singly linked list.
* Todo: Solve this recursively.
*
* */

import leetcode.linkedList.common.ListNode;

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode newHead = null;

        while (head != null) {
            ListNode next = head.next;
            head.next = newHead;
            newHead = head;

            head = next;
        }

        return newHead;
    }
}
