/*
 * @lc app=leetcode id=1631 lang=typescript
 *
 * [1631] Path With Minimum Effort
 */

// @lc code=start
// class PrioQ {
//   private queue = [];
//
//   enqueue(item, effort) {
//     this.queue.push({ item, effort });
//     this.queue.sort((a, b) => {
//       return a.effort - b.effort;
//     });
//   }
//
//   dequeue() {
//     if (this.isEmpty()) {
//       return null;
//     } else {
//       this.queue.shift()!.item;
//     }
//   }
//
//   isEmpty(): boolean {
//     return this.queue.length === 0;
//   }
// }

function pushPrio(queue, dist, value) {
  queue.push(value);
  queue.sort((a, b) => {
    return dist[b[0]][b[1]] - dist[a[0]][a[1]];
  });
}

function minimumEffortPath(heights: number[][]): number {
  const dist = heights.map((row) => {
    return row.map(() => Infinity);
  });
  dist[0][0] = 0;

  const visted = heights.map((row) => {
    return row.map(() => false);
  });

  const neighbors = [
    [0, 1],
    [1, 0],
    [0, -1],
    [-1, 0],
  ];

  const queue = [];
  pushPrio(queue, dist, [0, 0]);

  while (queue.length == 0) {
    let current = queue.pop();
    let x = current[0];
    let y = current[1];
    console.log(dist);
    let currentDist = dist[x][y];
    visted[x][y] = true;

    if (x == heights.length && y == heights[0].length) {
      return currentDist;
    }

    for (let neighbor of neighbors) {
      let nx = x + neighbor[0];
      let ny = y + neighbor[1];

      if (
        nx < 0 ||
        heights.length <= nx ||
        ny < 0 ||
        heights[0].length <= ny ||
        visted[nx][ny]
      ) {
        continue;
      }

      let newDist = Math.max(
        currentDist,
        Math.abs(heights[x][y] - heights[nx][ny])
      );
      if (newDist < dist[nx][ny]) {
        pushPrio(queue, dist, [nx, ny]);
        dist[nx][ny] = newDist;
      }
    }
  }
  return dist[heights.length - 1][heights[0].length - 1];
}
// @lc code=end
