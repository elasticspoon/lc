/*
 * @lc app=leetcode id=847 lang=typescript
 *
 * [847] Shortest Path Visiting All Nodes
 */

// @lc code=start
class NodeState {
  public index: number;
  public mask: number;
  public steps: number;

  constructor(index: number, mask: number, steps: number) {
    this.index = index;
    this.mask = mask;
    this.steps = steps;
  }
}

function hashMask(mask: number, index: number): number {
  return (mask << 4) + index;
}

function shortestPathLength(graph: number[][]): number {
  const visited = {};
  const targetState = (1 << graph.length) - 1;

  const queue: NodeState[] = [];
  for (let i = 0; i < graph.length; i++) {
    queue.push(new NodeState(i, 1 << i, 0));
  }

  while (queue.length > 0) {
    let { index, mask, steps } = queue.shift();

    if (mask == targetState) {
      return steps;
    }

    for (let n of graph[index]) {
      let newMask = mask | (1 << n);
      if (!visited[hashMask(newMask, n)]) {
        queue.push(new NodeState(n, newMask, steps + 1));
        visited[hashMask(newMask, n)] = true;
      }
    }
  }

  return -1;
}
// @lc code=end
