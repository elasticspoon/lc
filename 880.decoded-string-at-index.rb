#
# @lc app=leetcode id=880 lang=ruby
#
# [880] Decoded String at Index
#

# @lc code=start
# @param {String} s
# @param {Integer} k
# @return {String}
Node = Struct.new(:node, :string, :size)
def decode_at_index(s, k)
  i = 0
  tree = Node.new(nil, "", 0)

  s.each_char do |c|
    break if k <= i
    if /[a-z]/.match?(c)
      tree.string += c
      tree.size += 1
      i += 1
    else
      v = c.to_i
      i *= v
      tree = Node.new(tree, "", i)
    end
  end
  k -= 1

  # puts tree.inspect
  # puts "loop #{s}"
  while k >= 0
    # puts "k: #{k}, tree.string.length: #{tree.string.length}, tree.size: #{tree.size} tree.string: #{tree.string}"
    k %= tree.size if !tree.size.zero? && k >= tree.size
    # puts "return #{k - tree.size + tree.string.length - 1}"
    if tree.string.length > 0 && k - tree.size + tree.string.length >= 0
      # puts k - tree.size + tree.string.length - 1
      return tree.string[k - tree.size + tree.string.length]
    else
      # tree.size -= tree.string.length
      # k -= tree.string.length
      # k %= (tree.size - tree.string.length)
      tree = tree.node
    end
  end

  nil
end
# @lc code=end

puts 1
res = decode_at_index("leet2code3", 10)
puts res.inspect unless res == "o"
puts 2
res = decode_at_index("leet", 2)
puts res.inspect unless res == "e"
puts 3
res = decode_at_index("ha22", 5)
puts res.inspect unless res == "h"
puts 4
res = decode_at_index("a2345678999999999999999", 1)
puts res.inspect unless res == "a"
puts 5
res = decode_at_index("a2b3c4d5e6f7g8h9", 10)
puts res.inspect unless res == "c"
puts 6
res = decode_at_index("vk6u5xhq9v", 554)
puts res.inspect unless res == "k"
puts 7
res = decode_at_index("y959q969u3hb22odq595", 222280369)
puts res.inspect unless res == "y"
puts 8
res = decode_at_index("ixm5xmgo78", 711)
puts res.inspect unless res == "x"
