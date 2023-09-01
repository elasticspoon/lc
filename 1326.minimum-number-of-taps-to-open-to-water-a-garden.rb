#
# @lc app=leetcode id=1326 lang=ruby
#
# [1326] Minimum Number of Taps to Open to Water a Garden
#
# https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/description/
#
# algorithms
# Hard (47.34%)
# Likes:    2985
# Dislikes: 161
# Total Accepted:    109.1K
# Total Submissions: 212K
# Testcase Example:  '5\n[3,4,1,1,0,0]'
#
# There is a one-dimensional garden on the x-axis. The garden starts at the
# point 0 and ends at the point n. (i.e The length of the garden is n).
#
# There are n + 1 taps located at points [0, 1, ..., n] in the garden.
#
# Given an integer n and an integer array ranges of length n + 1 where
# ranges[i] (0-indexed) means the i-th tap can water the area [i - ranges[i], i
# + ranges[i]] if it was open.
#
# Return the minimum number of taps that should be open to water the whole
# garden, If the garden cannot be watered return -1.
#
#
# Example 1:
#
#
# Input: n = 5, ranges = [3,4,1,1,0,0]
# Output: 1
# Explanation: The tap at point 0 can cover the interval [-3,3]
# The tap at point 1 can cover the interval [-3,5]
# The tap at point 2 can cover the interval [1,3]
# The tap at point 3 can cover the interval [2,4]
# The tap at point 4 can cover the interval [4,4]
# The tap at point 5 can cover the interval [5,5]
# Opening Only the second tap will water the whole garden [0,5]
#
#
# Example 2:
#
#
# Input: n = 3, ranges = [0,0,0,0]
# Output: -1
# Explanation: Even if you activate all the four taps you cannot water the
# whole garden.
#
#
#
# Constraints:
#
#
# 1 <= n <= 10^4
# ranges.length == n + 1
# 0 <= ranges[i] <= 100
#
#
#

# @lc code=start
# @param {Integer} n
# @param {Integer[]} ranges
# @return {Integer}
# def min_taps(n, ranges)
#   intervals = ranges.filter_map.with_index do |r, i|
#     next if r.zero?
#     [i - r, i + r]
#   end.sort_by(&:first)
#
#   puts intervals.inspect
#
#   taps = 0
#   garden_spot = 0
#
#   while garden_spot < n
#     t = intervals.filter { |interval| cover(interval, garden_spot) && interval.last > garden_spot }
#     return -1 if t.empty?
#     t.sort_by!(&:last)
#
#     tap = t.last
#     taps += 1
#     garden_spot = tap.last
#   end
#   taps
# end

def min_taps(n, ranges)
  intervals = ranges.filter_map.with_index do |r, i|
    next if r.zero?
    (i - r..i + r)
  end.sort_by(&:begin)

  taps = 0
  garden_spot = 0

  while garden_spot < n
    t = intervals.filter { |interval| interval.cover?(garden_spot) && interval.last > garden_spot }
    return -1 if t.empty?
    t.sort_by!(&:end)

    tap = t.last
    taps += 1
    garden_spot = tap.last
  end
  taps
end
#
# def cover(tuple, val)
#   return true if tuple.first <= val && tuple.last >= val
#
#   false
# end

# @lc code=end
