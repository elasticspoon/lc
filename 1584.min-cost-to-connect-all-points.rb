#
# @lc app=leetcode id=1584 lang=ruby
#
# [1584] Min Cost to Connect All Points
#

# @lc code=start
# @param {Integer[][]} points
# @return {Integer}
# def min_cost_connect_points(points)
#   return 0 if points.length == 1
#   edge_list = []
#
#   points.each_with_index do |a, i|
#     points.each_with_index do |b, j|
#       next if a == b
#       dist = dist(a, b)
#       edge_list.push([b, dist, a])
#     end
#   end
#
#   edge_list.sort_by! { |e| e[1] }
#
#   sum = 0
#   puts edge_list.inspect
#   visited = []
#
#   loop do
#     if ()
#
#
#
#     break if visited.length == points.length
#   end
#
#   sum
# end

def min_cost_connect_points(points)
  return 0 if points.length == 1
  hash = {}

  points.each_with_index do |a, i|
    points.each_with_index do |b, j|
      next if a == b
      dist = dist(a, b)
      hash[a] = (hash[a] || []).push([b, dist])
    end
  end

  visited = {}
  queue = []
  current = points.first
  sum = 0

  loop do
    visited[current] = true

    queue.concat(hash[current])
    queue = queue.reject { |v| visited[v[0]] }

    min_weight = 0
    min_edge = nil
    queue.each do |edge|
      if !visited[edge[1]] &&
          (min_edge.nil? || edge[1] < min_weight)
        min_weight = edge[1]
        min_edge = edge
      end
    end
    break if min_edge.nil?

    queue.delete(min_edge)
    sum += min_weight
    current = min_edge[0]

    # do while end
    break if visited.length == points.length
  end

  sum
end

def dist(a, b)
  x1, y1 = a
  x2, y2 = b

  (x1 - x2).abs + (y1 - y2).abs
end
# @lc code=end
