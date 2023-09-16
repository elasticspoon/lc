function camelize(string) {
  return string
    .split("-")
    .map((item, index) => {
      return index == 0 ? item : item[0].toUpperCase() + item.slice(1);
    })
    .join("");
}

function filterRange(arr, a, b) {
  return arr.filter((item) => {
    return item <= b && a <= item;
  });
}

function filterRangeInPlace(arr, a, b) {
  let j = arr.filter((item) => {
    return item <= b && a <= item;
  });

  arr.length = 0;
  for (let v of j) {
    arr.push(v);
  }
}

function copySorted(arr) {
  let arr = [].concat(arr);
  return arr.sort((a, b) => {
    return a - b;
  });
}

function Calculator() {
  this.calculate = function (string) {
    let args = string.split(" ");
    this[args[1]](args[0], args[2]);
  };
  this["+"] = function (a, b) {
    return a + b;
  };
  this["-"] = function (a, b) {
    return a - b;
  };
  this.addMethod = function (method, func) {
    this[method] = func;
  };
}

let range = {
  from: 1,
  to: 5,
  [Symbol.iterator]: function () {
    return {
      current: this.from,
      final: this.to,
      next: function () {
        return { done: this.current == this.final, value: this.current++ };
      },
    };
  },
};

function sumSalaries(salaries) {
  let sum = 0;
  for (let salary of Object.values(salaries)) {
    sum += salary;
  }
  return sum;
}

function count(obj) {
  Reflect.ownKeys(obj);
  return Object.entries(obj).length;
}

function inBetween(low, high) {
  return function (item) {
    return low <= item && item <= high;
  };
}

function inArray(array) {
  return function (item) {
    return array.includes(item);
  };
}

function byField(fieldName) {
  return function (a, b) {
    return a[fieldName] > b[fieldName] ? 1 : -1;
  };
}

function makeArmy() {
  let shooters = [];

  function makeShooter(i) {
    let shooter = function () {
      // create a shooter function,
      alert(i); // that should show its number
    };
    shooters.push(shooter); // and add it to the array
  }

  let i = 0;
  while (i < 10) {
    makeShooter(i);
    i++;
  }

  // ...and return the array of shooters
  return shooters;
}

let army = makeArmy();
