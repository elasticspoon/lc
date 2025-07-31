require 'minitest/autorun'

# @param {Integer[]} nums
# @return {Integer}
def count_hill_valley(nums)
  prev_slope = 0
  hill_valleys = 0
  nums.each_cons(2) do |prev, curr|
    curr_slope = prev <=> curr
    next if curr_slope.zero?

    hill_valleys += 1 if (prev_slope != curr_slope) && !prev_slope.zero?
    prev_slope = curr_slope
  end

  hill_valleys
end

class Test < Minitest::Test
  def test_one
    got = count_hill_valley([2, 4, 1, 1, 6, 5])
    want = 3
    assert_equal(want, got)
  end

  def test_two
    got = count_hill_valley([6, 6, 5, 5, 4, 1])
    want = 0
    assert_equal(want, got)
  end

  def test_three
    got = count_hill_valley([1, 1, 1])
    want = 0
    assert_equal(want, got)
  end
end
