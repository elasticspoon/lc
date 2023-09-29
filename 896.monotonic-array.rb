#
# @lc app=leetcode id=896 lang=ruby
#
# [896] Monotonic Array
#

# @lc code=start
# @param {Integer[]} nums
# @return {Boolean}
def is_monotonic(nums)
  return true if nums.length <= 2

  diff = nums.first <=> nums.last

  nums.each_with_index do |num, i|
    next if i == 0
    current_diff = nums[i - 1] <=> nums[i]

    return false if diff != current_diff && current_diff != 0
  end

  true
end
# @lc code=end

# puts is_monotonic([1, 2, 2, 3]) == true
# puts is_monotonic([6, 5, 4, 4]) == true
# puts is_monotonic([1, 3, 2]) == false
# puts is_monotonic([1]) == true
# puts is_monotonic([1, 2]) == true
# puts is_monotonic([3, 2]) == true
