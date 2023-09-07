/*
 * @lc app=leetcode id=92 lang=typescript
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

class ListNode92 {
  val: number;
  next: ListNode92 | null;
  constructor(val?: number, next?: ListNode92 | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}
function reverseBetween(
  head: ListNode92 | null,
  left: number,
  right: number
): ListNode92 | null {
  if (head === null || right - left == 0) {
    return head;
  }

  const dummy = new ListNode92(undefined, head);
  let swapHead = dummy;
  let index = 1;

  while (index < left) {
    swapHead = swapHead.next;
    index++;
  }

  let tail = swapHead.next;
  while (index < right) {
    let tmp = swapHead.next;
    swapHead.next = tail.next;
    tail.next = tail.next.next;
    swapHead.next.next = tmp;

    index++;
  }

  return dummy.next;
}
// @lc code=end
