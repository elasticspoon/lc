#
# @lc app=leetcode id=118 lang=ruby
#
# [118] Pascal's Triangle
#

# @lc code=start
# @param {Integer} num_rows
# @return {Integer[][]}
def generate(num_rows)
  list = [[1]]

  (1...num_rows).each do |i|
    row = [1]
    row[i] = 1
    prev_row = list[i - 1]

    (1...prev_row.length).each do |j|
      row[j] = prev_row[j - 1] + prev_row[j]
    end
    list[i] = row
  end

  list
end
# @lc code=end
