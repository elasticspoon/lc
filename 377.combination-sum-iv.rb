#
# @lc app=leetcode id=377 lang=ruby
#
# [377] Combination Sum IV
#

# @lc code=start
# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer}
# def combination_sum4(nums, target)
#   puts "nums: #{nums}, target: #{target}"
#   return 1 if target == 0
#
#   res = 0
#   (0...nums.size).each do |i|
#     res += combination_sum4(nums, target - nums[i]) if target >= nums[i]
#   end
#
#   res
# end
def combination_sum4(nums, target)
  dp = [1]

  (1..target).each do |i|
    dp[i] = 0
    nums.each { |num| dp[i] += dp[i - num] if i >= num }
  end

  puts dp.inspect
  dp[target]
end
# @lc code=end
