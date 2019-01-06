package leetcode.linkedList.Delete_Node_in_a_LinkedList_237;

/*
*
* Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.
  Supposed the linked list is 1 -> 2 -> 3 -> 4 and you are given the third node with value 3, the linked list should become 1 -> 2 -> 4 after calling your function.
*
* https://leetcode.com/problems/delete-node-in-a-linked-list/description/
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
    public void deleteNode(ListNode node) {
        //Todo: note! this solution does not check if node.next.next not null.
        node.val = node.next.val;
        node.next = node.next.next;
    }
}
