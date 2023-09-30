/*
 * @lc app=leetcode id=456 lang=typescript
 *
 * [456] 132 Pattern
 */

// @lc code=start
function find132pattern(nums: number[]): boolean {
  const stack: number[] = [];
  let second: number = Number.MIN_SAFE_INTEGER;

  for (let i = nums.length - 1; i >= 0; i--) {
    if (nums[i] < second) {
      return true;
    }

    while (stack.length > 0 && nums[i] > stack[stack.length - 1]) {
      second = stack.pop()!;
    }
    stack.push(nums[i]);
  }

  return false;
}
// @lc code=end

// O(n^2)
function find132patternO2(nums: number[]): boolean {
  let numI = nums[0];
  for (let j = 1; j < nums.length - 1; j++) {
    if (nums[j - 1] <= numI) {
      numI = nums[j - 1];
    }
    for (let k = j + 1; k < nums.length; k++) {
      if (nums[j] > nums[k] && nums[k] > numI) {
        return true;
      }
    }
  }

  return false;
}
// O(n^3)
function find132patternO3(nums: number[]): boolean {
  for (let i = 0; i < nums.length; i++) {
    for (let k = nums.length - 1; k > i + 1; k--) {
      if (nums[i] > nums[k]) {
        continue;
      }
      for (let j = k - 1; j > i; j--) {
        if (nums[i] < nums[k] && nums[k] < nums[j]) {
          return true;
        }
      }
    }
  }

  return false;
}
function find132pattern1Pass(nums: number[]): boolean {
  let n = nums.length;
  let top = n;
  let third = -Infinity;

  for (let i = n - 1; i >= 0; i--) {
    if (nums[i] < third) {
      return true;
    }
    while (top < n && nums[i] > nums[top]) {
      third = nums[top];
      top++;
    }
    top--;
    nums[top] = nums[i];
  }

  return false;
}
console.log(find132pattern([1, 2, 3, 4]) === false);
console.log(find132pattern([3, 1, 4, 2]) === true);
console.log(find132pattern([-1, 3, 2, 0]) === true);
