/*
 * @lc app=leetcode id=332 lang=typescript
 *
 * [332] Reconstruct Itinerary
 */

// @lc code=start
function findItinerary(tickets: string[][]): string[] {
  const flights: Map<string, string[]> = new Map<string, string[]>();
  for (let ticket of tickets) {
    let list: string[] = flights.get(ticket[0]) ?? [];
    list.push(ticket[1]);
    flights.set(ticket[0], list);
  }

  for (let destinations of flights.values()) {
    destinations.sort((a, b) => (a < b ? 1 : -1));
  }

  const stack: string[] = ["JFK"];
  const iten: string[] = [];

  while (stack.length > 0) {
    let current: string = stack.at(-1);
    let next: string | undefined = flights.get(current)?.pop();

    if (next) {
      stack.push(next);
    } else {
      iten.push(stack.pop());
    }
  }

  return iten.reverse();
}
// @lc code=end
