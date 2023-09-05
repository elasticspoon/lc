/*
 * @lc app=leetcode id=138 lang=typescript
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for Node.
 * class Node {
 *     val: number
 *     next: Node | null
 *     random: Node | null
 *     constructor(val?: number, next?: Node, random?: Node) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *         this.random = (random===undefined ? null : random)
 *     }
 * }
 */

// had to change name of this so that typescript would be happy
class RandNode {
  val: number;
  next: RandNode | null;
  random: RandNode | null;
  constructor(val?: number, next?: RandNode, random?: RandNode) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
    this.random = random === undefined ? null : random;
  }
}
function copyRandomList(head: RandNode | null): RandNode | null {
  if (head === null) {
    return head;
  }

  // build interweaved list
  let current = head;
  while (current !== null) {
    current.next = new RandNode(current.val, current.next, null);
    current = current.next.next;
  }

  // set random
  current = head;
  while (current !== null) {
    if (current.random) {
      current.next.random = current.random.next;
    }
    current = current.next.next;
  }

  current = head;
  let newHead = head.next;
  while (current !== null) {
    let newNode = current.next;
    current.next = newNode.next;
    current = newNode.next;
    if (newNode.next !== null) {
      newNode.next = newNode.next.next;
    }
  }
  // fix lists
  return newHead;
}
// @lc code=end
