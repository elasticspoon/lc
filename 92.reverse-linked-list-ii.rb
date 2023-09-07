#
# @lc app=leetcode id=92 lang=ruby
#
# [92] Reverse Linked List II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val = 0, _next = nil)
#         @val = val
#         @next = _next
#     end
# end
# @param {ListNode} head
# @param {Integer} left
# @param {Integer} right
# @return {ListNode}
def reverse_between(head, left, right)
  return head if head.nil? || left - right == 0

  dummy = ListNode.new(nil, head)
  swap_head = dummy
  index = 1

  while index < left
    swap_head = swap_head.next
    index += 1
  end

  tail = swap_head.next

  while index < right
    tmp = swap_head.next
    swap_head.next = tail.next
    tail.next = tail.next.next
    swap_head.next.next = tmp
    index += 1
  end

  dummy.next
end
# @lc code=end
