#
# @lc app=leetcode id=39 lang=ruby
#
# [39] Combination Sum
#

# @lc code=start
# @param {Integer[]} candidates
# @param {Integer} target
# @return {Integer[][]}
# def combination_sum(candidates, target)
#   dp = [[[]]]
#
#   (1..target).each do |i|
#     candidates.each do |candidate|
#       if candidate <= i && !dp[i - candidate].nil?
#         dp[i] ||= []
#         dp[i - candidate].each do |v|
#           dp[i].push(v.clone.append(candidate).sort)
#         end
#         dp[i] = [[candidate]] if dp[i] == []
#         dp[i] = dp[i].uniq
#       end
#     end
#   end
#
#   dp[target] || []
# end
def combination_sum(candidates, target)
  dp = Array.new(target + 1) { [] }

  candidates.each do |candidate|
    (candidate..target).each do |i|
      dp[i].push([candidate]) if candidate == i

      dp[i - candidate].each do |v|
        dp[i].push(v + [candidate])
      end
    end
  end

  dp[target]
end
# @lc code=end
