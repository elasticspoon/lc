def validate(s, dict)
  dict = dict.map { [_1, true] }.to_h
  l, r = [0, 1]
  while r <= s.length
    l = r if dict[(s[l...r])]
    r += 1
  end
  l == r - 1
end

puts validate("leetcode", %w[leet code])
puts validate("applepenapple", %w[apple pen])
puts validate("practicemakesperfect", %w[practice makes perfect])
puts validate("practicemakesperfectx", %w[practice makes perfect])
