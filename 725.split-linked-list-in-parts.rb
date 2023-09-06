#
# @lc app=leetcode id=725 lang=ruby
#
# [725] Split Linked List in Parts
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
# @param {Integer} k
# @return {ListNode[]}
def split_list_to_parts(head, k)
  list = Array.new(k) { nil }

  current = head
  len = 0

  until current.nil?
    current = current.next
    len += 1
  end

  whole = len / k
  rem = len % k

  current = head
  prev = head
  i = 0

  until current.nil?
    list[i] = current

    (0...whole).each do
      prev = current
      current = current.next
    end

    if rem > 0
      rem -= 1
      prev = current
      current = current.next
    end

    prev.next = nil
    prev = current
    i += 1
  end

  list
end
# @lc code=end
