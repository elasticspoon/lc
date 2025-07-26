require 'minitest/autorun'

# @param {Integer[]} nums
# @return {Integer}
def max_sum(nums)
  top = nums.max
  nums.filter(&:positive?).push(top).uniq.sum
end

class Test < Minitest::Test
  def test_no_deletions
    got = max_sum([1, 2, 3, 4, 5])
    want = 15

    assert_equal(want, got)
  end

  def test_two
    got = max_sum([1, 1, 0, 1, 1])
    want = 1

    assert_equal(want, got)
  end

  def test_leaves_one
    got = max_sum([-5])
    want = -5

    assert_equal(want, got)
  end

  def test_three
    got = max_sum([1, 2, -1, -2, 1, 0, -1])
    want = 3

    assert_equal(want, got)
  end
end
