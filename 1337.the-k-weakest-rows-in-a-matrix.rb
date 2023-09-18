#
# @lc app=leetcode id=1337 lang=ruby
#
# [1337] The K Weakest Rows in a Matrix
#

# @lc code=start
# @param {Integer[][]} mat
# @param {Integer} k
# @return {Integer[]}
def k_weakest_rows(mat, k)
  strength = mat.map.with_index { |row, i| [row.sum, i] }
  # idk why this works
  strength.sort_by { |s| s[0] }.map { |s| s[1] }[0...k]
end
# @lc code=end
