#
# @lc app=leetcode id=1584 lang=ruby
#
# [1584] Min Cost to Connect All Points
#

# @lc code=start
# @param {Integer[][]} points
# @return {Integer}
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

def min_cost_connect_points(points)
  weight = 0
  distances = Array.new(points.length) { Float::INFINITY }
  distances[0] = 0

  visited = []
  queue = PriorityQueue.new
  queue.push(0, 0)

  until queue.empty?
    current = queue.pop
    current_dist = distances[current]

    next if visited[current]

    visited[current] = true
    weight += current_dist

    points.each_with_index do |point, index|
      next if visited[index]
      next_dist = manhattan_dist(points[current], point)

      if next_dist < distances[index]
        queue.push(index, next_dist)
        distances[index] = next_dist
      end
    end
  end
  weight
end

def manhattan_dist(a, b)
  x1, y1 = a
  x2, y2 = b

  (x1 - x2).abs + (y1 - y2).abs
end
# @lc code=end
