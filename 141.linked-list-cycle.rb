#
# @lc app=leetcode id=141 lang=ruby
#
# [141] Linked List Cycle
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val)
#         @val = val
#         @next = nil
#     end
# end
def _hasCycle(head)
  until head.nil?
    return true if head.val.nil?
    head.val = nil
    head = head.next
  end
  false
end

# @param {ListNode} head
# @return {Boolean}
def hasCycle(head)
  fast, slow = head, head

  until fast.nil? || fast.next.nil?
    fast = fast.next.next
    slow = slow.next

    return true if fast == slow

  end

  false
end
# @lc code=end
