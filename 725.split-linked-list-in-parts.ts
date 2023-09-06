/*
 * @lc app=leetcode id=725 lang=typescript
 *
 * [725] Split Linked List in Parts
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

class ListNode725 {
  val: number;
  next: ListNode725 | null;
  constructor(val?: number, next?: ListNode725 | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}
function splitListToParts(
  head: ListNode725 | null,
  k: number
): Array<ListNode725 | null> {
  let len = 0;
  let current = head;

  while (current !== null) {
    len++;
    current = current.next;
  }

  const list: Array<ListNode725 | null> = [];

  for (let i = 0; i < k; i++) {
    list.push(null);
  }
  current = head;
  let prev = head;

  const fullS = Math.floor(len / k);
  let remS = len % k;
  let i = 0;

  while (current !== null) {
    list[i] = current;

    if (remS > 0) {
      prev = current;
      current = current.next;
      remS--;
    }

    for (let j = 0; j < fullS; j++) {
      prev = current;
      current = current.next;
    }

    prev.next = null;
    prev = current;
    i++;
  }

  return list;
}
// @lc code=end
