#
# @lc app=leetcode id=1658 lang=ruby
#
# [1658] Minimum Operations to Reduce X to Zero
#

# @lc code=start
# @param {Integer[]} nums
# @param {Integer} x
# @return {Integer}
def min_operations(nums, x)
  reverse_hash = {0 => 0}

  accum = 0
  i = nums.length - 1
  while i >= 0
    accum += nums[i]

    break if accum > x

    reverse_hash[accum] = nums.length - i
    i -= 1
  end

  min_ops = Float::INFINITY
  forward_ops = 0
  accum = 0

  while accum <= x && forward_ops < nums.length
    revese_ops = reverse_hash[x - accum]
    if revese_ops && forward_ops + revese_ops < min_ops && forward_ops + revese_ops <= nums.length
      min_ops = forward_ops + revese_ops
    end

    accum += nums[forward_ops]
    forward_ops += 1
  end

  (min_ops == Float::INFINITY) ? -1 : min_ops
end
# @lc code=end
