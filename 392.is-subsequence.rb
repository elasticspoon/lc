#
# @lc app=leetcode id=392 lang=ruby
#
# [392] Is Subsequence
#

# @lc code=start
# @param {String} s
# @param {String} t
# @return {Boolean}
def is_subsequence(s, t)
  main_idx = 0

  s.each_char do |l|
    letter_found = false
    while main_idx < t.length
      if t[main_idx] != l
        main_idx += 1
      else
        letter_found = true
        main_idx += 1
        break
      end
    end
    return false unless letter_found
  end
  true
end

# @lc code=end
