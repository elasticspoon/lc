#
# @lc app=leetcode id=1631 lang=ruby
#
# [1631] Path With Minimum Effort
#

# @lc code=start
# @param {Integer[][]} heights
# @return {Integer}require 'thread'

class PriorityQueue
  def initialize
    @data = []
    @mutex = Mutex.new
  end

  def push(item, priority)
    @mutex.synchronize do
      @data << [item, priority]
      heapify_up(@data.size - 1)
    end
  end

  def pop
    @mutex.synchronize do
      return nil if empty?
      swap(0, @data.size - 1)
      item, _ = @data.pop
      heapify_down(0)
      item
    end
  end

  def empty?
    @data.empty?
  end

  private

  def swap(i, j)
    @data[i], @data[j] = @data[j], @data[i]
  end

  def heapify_up(index)
    return if index.zero?

    parent_index = (index - 1) / 2
    return if @data[parent_index][1] <= @data[index][1]

    swap(parent_index, index)
    heapify_up(parent_index)
  end

  def heapify_down(index)
    left_child_index = 2 * index + 1
    right_child_index = 2 * index + 2
    smallest = index

    if left_child_index < @data.size && @data[left_child_index][1] < @data[smallest][1]
      smallest = left_child_index
    end

    if right_child_index < @data.size && @data[right_child_index][1] < @data[smallest][1]
      smallest = right_child_index
    end

    if smallest != index
      swap(index, smallest)
      heapify_down(smallest)
    end
  end
end

def minimum_effort_path(heights)
  dist = heights.map do |r|
    r.map { Float::INFINITY }
  end
  dist[0][0] = 0

  visited = Array.new(heights.length) { Array.new(heights.first.length) { false } }

  neighbors = [[0, 1], [0, -1], [1, 0], [-1, 0]]

  queue = PriorityQueue.new
  queue.push([0, 0], 0)

  until queue.empty?
    x, y = queue.pop
    current_dist = dist[x][y]
    visited[x][y] = true

    return current_dist if x == heights.length && y == heights.first.length

    neighbors.each do |cx, cy|
      nx, ny = cx + x, cy + y

      if nx < 0 || heights.length <= nx ||
          ny < 0 || heights.first.length <= ny ||
          visited[nx][ny]
        next
      end

      new_dist = [current_dist, (heights[x][y] - heights[nx][ny]).abs].max
      if new_dist < dist[nx][ny]
        dist[nx][ny] = new_dist
        queue.push([nx, ny], new_dist)
      end
    end
  end

  dist[heights.length - 1][heights.first.length - 1]
end
# @lc code=end

heights = [[1, 2, 1, 1, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 1, 1, 2, 1]]
puts minimum_effort_path(heights) == 0
heights = [[1, 2, 3], [3, 8, 4], [5, 3, 5]]
puts minimum_effort_path(heights) == 1
heights = [[1, 2, 2], [3, 8, 2], [5, 3, 5]]
puts minimum_effort_path(heights) == 2
