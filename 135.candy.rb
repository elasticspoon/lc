#
# @lc app=leetcode id=135 lang=ruby
#
# [135] Candy
#

# @lc code=start
# @param {Integer[]} ratings
# @return {Integer}
def candy(ratings)
  sum = 1
  cons_inc = 0
  cons_dec = 0
  peak = 0

  (1...ratings.length).each do |i|
    if ratings[i] > ratings[i - 1]
      cons_dec = 0
      cons_inc += 1
      peak = cons_inc
      sum += cons_inc + 1
    elsif ratings[i] == ratings[i - 1]
      cons_dec = 0
      cons_inc = 0
      peak = 0
      sum += 1
    else
      cons_inc = 0
      cons_dec += 1
      sum += cons_dec + ((peak >= cons_dec) ? 0 : 1)
    end
  end

  sum
end
# @lc code=end
