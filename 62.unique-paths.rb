#
# @lc app=leetcode id=62 lang=ruby
#
# [62] Unique Paths
#
# https://leetcode.com/problems/unique-paths/description/
#
# algorithms
# Medium (63.02%)
# Likes:    15443
# Dislikes: 410
# Total Accepted:    1.6M
# Total Submissions: 2.5M
# Testcase Example:  '3\n7'
#
# There is a robot on an m x n grid. The robot is initially located at the
# top-left corner (i.e., grid[0][0]). The robot tries to move to the
# bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move
# either down or right at any point in time.
#
# Given the two integers m and n, return the number of possible unique paths
# that the robot can take to reach the bottom-right corner.
#
# The test cases are generated so that the answer will be less than or equal to
# 2 * 10^9.
#
#
# Example 1:
#
#
# Input: m = 3, n = 7
# Output: 28
#
#
# Example 2:
#
#
# Input: m = 3, n = 2
# Output: 3
# Explanation: From the top-left corner, there are a total of 3 ways to reach
# the bottom-right corner:
# 1. Right -> Down -> Down
# 2. Down -> Down -> Right
# 3. Down -> Right -> Down
#
#
#
# Constraints:
#
#
# 1 <= m, n <= 100
#
#
#

def unique_paths_slower(m, n)
  return 1 if m == 1 || n == 1

  fact(m + n - 2) / (fact(m - 1) * fact(n - 1))
end

def fact(n)
  (1..n).inject(:*)
end

# @lc code=start
# @param {Integer} m
# @param {Integer} n
# @return {Integer}
def unique_paths(m, n)
  return 1 if m == 1 || n == 1

  if m > n
    (m..(m + n - 2)).inject(:*) / (1..n - 1).inject(:*)
  else
    (n..(m + n - 2)).inject(:*) / (1..m - 1).inject(:*)
  end
end

# @lc code=end
