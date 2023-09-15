#
# @lc app=leetcode id=332 lang=ruby
#
# [332] Reconstruct Itinerary
#

# @lc code=start
# @param {String[][]} tickets
# @return {String[]}
def find_itinerary(tickets)
  hash = {}
  tickets.each do |s, d|
    (hash[s] ||= []).push(d)
  end

  hash.each_value do |v|
    v.sort!.reverse!
  end

  itinerary = []
  stack = ["JFK"]

  until stack.empty?
    while hash[stack.last]&.any?
      stack.push(hash[stack.last].pop)
    end

    itinerary.push(stack.pop)
  end

  itinerary.reverse
end
# @lc code=end
