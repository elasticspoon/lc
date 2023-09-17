#
# @lc app=leetcode id=847 lang=ruby
#
# [847] Shortest Path Visiting All Nodes
#

# @lc code=start
# @param {Integer[][]} graph
# @return {Integer}
def hash_mask(mask, index)
  (mask << 4) + index
end

def shortest_path_length(graph)
  node_state = Struct.new(:index, :mask, :steps)

  visited = {}
  target_state = (1 << graph.length) - 1

  queue = graph.each_with_index.map do |neighbors, index|
    node_state.new(index, 1 << index, 0)
  end

  until queue.empty?
    s = queue.shift
    index = s.index
    mask = s.mask
    steps = s.steps

    return steps if mask == target_state

    graph[index].each do |neighbor|
      new_mask = mask | (1 << neighbor)
      if !visited[hash_mask(new_mask, neighbor)]
        queue.push(node_state.new(neighbor, new_mask, steps + 1))
        visited[hash_mask(new_mask, neighbor)] = true
      end
    end

  end

  -1
end
# @lc code=end
