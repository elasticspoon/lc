#
# @lc app=leetcode id=287 lang=ruby
#
# [287] Find the Duplicate Number
#

# @lc code=start
# @param {Integer[]} nums
# @return {Integer}
def find_duplicate(nums)
  array = Array.new(nums.size) { false }
  nums.each do |num|
    return num if array[num]
    array[num] = true
  end
  -1
end
# @lc code=end
