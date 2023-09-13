#
# @lc app=leetcode id=1647 lang=ruby
#
# [1647] Minimum Deletions to Make Character Frequencies Unique
#

# @lc code=start
# @param {String} s
# @return {Integer}
def min_deletions(s)
  histogram = {}
  s.each_char { |c| histogram[c] = (histogram[c] || 0) + 1 }

  freq_arr = []
  histogram.each_value { |v| freq_arr[v] = (freq_arr[v] || 0) + 1 }

  subtractions = 0

  i = freq_arr.length
  until i.zero?
    if freq_arr[i] && freq_arr[i] > 1
      freq_arr[i - 1] = (freq_arr[i - 1] || 0) + freq_arr[i] - 1
      subtractions += freq_arr[i] - 1
    end
    i -= 1
  end

  subtractions
end
# @lc code=end
