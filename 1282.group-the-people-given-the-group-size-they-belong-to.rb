#
# @lc app=leetcode id=1282 lang=ruby
#
# [1282] Group the People Given the Group Size They Belong To
#

# @lc code=start
# @param {Integer[]} group_sizes
# @return {Integer[][]}
def group_the_people(group_sizes)
  hash = {}
  result = []

  group_sizes.each_with_index do |v, k|
    hash[v] ||= []
    hash[v].push(k)
    if hash[v].length == v
      result.push(hash[v])
      hash[v] = nil
    end
  end

  result
end
# @lc code=end
