#
# @lc app=leetcode id=2038 lang=ruby
#
# [2038] Remove Colored Pieces if Both Neighbors are the Same Color
#

# @lc code=start
# @param {String} colors
# @return {Boolean}
def winner_of_game(colors)
  num_alice = num_bob = 0
  cons_a = cons_b = 0

  colors.each_char do |char|
    if char == "A"
      cons_a += 1
      cons_b = 0
      num_alice += 1 if cons_a > 2
    elsif char == "B"
      cons_b += 1
      cons_a = 0
      num_bob += 1 if cons_b > 2
    end
  end

  num_alice - num_bob > 0
end
# @lc code=end
