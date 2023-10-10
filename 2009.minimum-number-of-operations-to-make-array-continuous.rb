#
# @lc app=leetcode id=2009 lang=ruby
#
# [2009] Minimum Number of Operations to Make Array Continuous
#

# @lc code=start
# @param {Integer[]} nums
# @return {Integer}
def min_operations(nums)
  n = nums.size
  i, j = [0, 0]
  nums = nums.sort.uniq
  m = nums.size

  puts "nums: #{nums}"
  while j < m
    puts "i: #{i}, j: #{j} nums[i] + n: #{nums[i] + n}, nums[j]: #{nums[j]}"
    i += 1 if nums[i] + n <= nums[j]
    j += 1
  end

  puts "res i: #{i}, j: #{j}"
  n - j + i
end
# @lc code=end

# puts min_operations([1, 100, 3, 4, 2, 2]) == 2
puts min_operations([3, 4, 8, 9, 10, 11, 12]) == 2
# puts min_operations([4, 2, 5, 3]) == 0
# puts min_operations([1, 2, 3, 5, 6]) == 1
# puts min_operations([1, 10, 100, 1000]) == 3
