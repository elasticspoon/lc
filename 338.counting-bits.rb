#
# @lc app=leetcode id=338 lang=ruby
#
# [338] Counting Bits
#
# https://leetcode.com/problems/counting-bits/description/
#
# algorithms
# Easy (76.32%)
# Likes:    10027
# Dislikes: 454
# Total Accepted:    889.8K
# Total Submissions: 1.2M
# Testcase Example:  '2'
#
# Given an integer n, return an array ans of length n + 1 such that for each i
# (0 <= i <= n), ans[i] is the number of 1's in the binary representation of
# i.
#
#
# Example 1:
#
#
# Input: n = 2
# Output: [0,1,1]
# Explanation:
# 0 --> 0
# 1 --> 1
# 2 --> 10
#
#
# Example 2:
#
#
# Input: n = 5
# Output: [0,1,1,2,1,2]
# Explanation:
# 0 --> 0
# 1 --> 1
# 2 --> 10
# 3 --> 11
# 4 --> 100
# 5 --> 101
#
#
#
# Constraints:
#
#
# 0 <= n <= 10^5
#
#
#
# Follow up:
#
#
# It is very easy to come up with a solution with a runtime of O(n log n). Can
# you do it in linear time O(n) and possibly in a single pass?
# Can you do it without using any built-in function (i.e., like
# __builtin_popcount in C++)?
#
#
#

# @lc code=start
# @param {Integer} n
# @return {Integer[]}
def count_bits(n)
  # bits = Array.new(n + 1) { 0 }
  bits = [0]

  (1..n).each do |i|
    bits[i] = bits[i / 2] + i % 2
  end

  bits
end
# @lc code=end
