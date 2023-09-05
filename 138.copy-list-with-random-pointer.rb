#
# @lc app=leetcode id=138 lang=ruby
#
# [138] Copy List with Random Pointer
#

# @lc code=start
# Definition for Node.
# class Node
#     attr_accessor :val, :next, :random
#     def initialize(val = 0)
#         @val = val
#		  @next = nil
#		  @random = nil
#     end
# end

# @param {Node} node
# @return {Node}
def copyRandomList(head)
  current = head
  until current.nil?
    node = Node.new(current.val)
    node.next = current.next
    current.next = node
    current = node.next
  end

  # current = head
  # until current.nil?
  #
  #   puts "current: #{current.val} next: #{current.next&.val} random: #{current.random&.val}"
  #   current = current.next&.next
  # end

  # puts "----------------------------\n"
  current = head
  until current.nil?
    new_node = current.next
    new_node.random = current.random&.next
    current = new_node.next
  end

  current = head
  new_head = head&.next
  until current.nil?
    new_node = current.next
    current.next = new_node.next
    new_node.next = new_node.next&.next

    current = current.next
  end

  new_head
end

# def copyRandomList(head)
#   list = []
#   current = head
#   index = 0
#   until current.nil?
#     current.random.val = [*current.random.val, index] unless current.random.nil?
#
#     current = current.next
#     index += 1
#   end
#
#   puts "loop 1"
#   current = head
#   index = 0
#   until current.nil?
#     node = Node.new(9)
#     node.val = current.val
#     list.push(node)
#
#     list[index - 1]&.next = node if index.positive?
#
#     current = current.next
#     index += 1
#   end
#
#   list.each do |node|
#     node.val, *pointers = node.val
#
#     pointers.each do |pointer|
#       list[pointer].random = node
#     end
#   end
#
#   list.first
# end
# def copyRandomList(head)
#   list = []
#   current = head
#   until current.nil?
#     list.push([current.val, node_index(current.random, head)])
#     current = current.next
#   end
#
#   new_list = []
#   list.each_with_index do |val, index|
#     node = Node.new(val.first)
#     new_list[index - 1]&.next = node
#     new_list.push(node)
#   end
#
#   list.each.with_index do |val, index|
#     new_list[index].random = new_list[val.last] unless val.last.negative?
#   end
#
#   new_list.first
# end
#
# def node_index(node, head)
#   index = 0
#   until head.nil?
#     return index if node == head
#
#     head = head.next
#     index += 1
#   end
#
#   -1
# end
# @lc code=end
