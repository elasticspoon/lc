#
# @lc app=leetcode id=1359 lang=ruby
#
# [1359] Count All Valid Pickup and Delivery Options
#

# @lc code=start
# @param {Integer} n
# @return {Integer}
MOD = 1_000_000_007
def count_orders(n)
  accum = 1

  (1..n).each do |i|
    accum = (accum * (i * 2 - 1) * i) % MOD
  end

  accum
end
# @lc code=end
