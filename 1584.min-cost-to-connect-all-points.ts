/*
 * @lc app=leetcode id=1584 lang=typescript
 *
 * [1584] Min Cost to Connect All Points
 */

// @lc code=start
function minCostConnectPoints(points: number[][]): number {
  function prioPop(queue: number[]) {
    let minIndex = 0;
    let minDist = Infinity;

    for (let i = 0; i < queue.length; i++) {
      if (minDist > distance[queue[i]]) {
        minIndex = i;
        minDist = distance[queue[i]];
      }
    }
    let tmp = queue[minIndex];
    queue[minIndex] = queue[queue.length - 1];
    queue[queue.length - 1] = tmp;

    return queue.pop();
  }
  const len = points.length;
  let visited = Array(len).fill(false);
  let visitsCount = 0;
  let weight = 0;
  const queue = [0];

  const distance = Array(len).fill(Infinity);
  distance[0] = 0;

  while (queue.length) {
    let current = prioPop(queue);
    let currentDist = distance[current];
    if (visited[current]) {
      continue;
    }
    visited[current] = true;
    visitsCount++;
    weight += currentDist;

    for (let i = 0; i < len; i++) {
      if (visited[i]) {
        continue;
      }

      let manhattanDist =
        Math.abs(points[current][0] - points[i][0]) +
        Math.abs(points[current][1] - points[i][1]);
      if (manhattanDist < distance[i]) {
        distance[i] = manhattanDist;
        queue.push(i);
      }
    }
  }

  return weight;
}

// @lc code=end
