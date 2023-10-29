function scrollBotton(element) {
  if (element) {
    return element.scrollHeight - element.scrollTop - element.clientHeight
  }
}

function scrollBarWidth() {
  get

  let elem = document.createElement("div")
  elem.width = "300px"
  elem.height = "300px"
  elem.style.overflowY = "scroll"

  document.body.append(elem)

  let scrollWidth = elem.offsetWidth - elem.clientWidth

  elem.remove()
  return scrollWidth
}

function addNumDesc(tree) {
  let lis = tree.querySelectorAll("li")
  for (let li of lis) {
    if (li.childElementCount > 0) {
      li.childNodes[0].textContent += ` [${li.querySelectorAll("li").length}]`
    }
  }
}
function sortTable(table, prop) {
  let headers = [...table.querySelectorAll("thead th")]
  let index = headers.findIndex((v) => v.textContent == prop)

  if (index == -1) {
    return
  }

  let rows = [...table.querySelectorAll("tbody tr")]
  rows.sort((a, b) => {
    return a.children[index].textContent > b.children[index].textContent ? 1 : -1
  })

  table.lastElementChild.append(...rows)
}

function makeCalendar(elem, year, month) {
  let calendar = document.createElement("table")
  let header = document.createElement("tr")
  for (let day of ["MO", "TU", "WE", "TH", "FR", "SA", "SU"]) {
    let th = document.createElement("th")
    th.textContent = day
    header.append(th)
  }
  calendar.append(header)

  let offset = (new Date(year, month - 1)).getDay() - 1
  let lastDay = (new Date(year, month, 0)).getDate()

  for (let i = 0; i < 6; i++) {
    let row = document.createElement("tr")
    for (let j = 0; j < 7; j++) {
      let td = document.createElement("td")
      let date = i * 7 + j + 1 - offset
      if (date <= lastDay && date > 0) {
        td.textContent = date
      }
      row.append(td)
    }
    if (row.innerText.trim().length == 0) {
      break
    }
    calendar.append(row)
  }
  elem.append(calendar)
}



function createTreeRec(container, data) {
  for (let [k, v] of Object.entries(data)) {
    let li = document.createElement("li")
    li.textContent = k
    if (Object.entries(v).length > 0) {
      let ul = document.createElement("ul")
      createTreeRec(ul, v)
      li.append(ul)
    }
    container.append(li)
  }
}
function createTree(container, data) {
  let ul = document.createElement("ul")
  createTreeRec(ul, data)
  container.append(ul)
}


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

function copySorted(array) {
  let arr = [].concat(array);
  return arr.sort((a, b) => {
    return a - b;
  });
}

function Calculator() {
  this.calculate = function(string) {
    let args = string.split(" ");
    this[args[1]](args[0], args[2]);
  };
  this["+"] = function(a, b) {
    return a + b;
  };
  this["-"] = function(a, b) {
    return a - b;
  };
  this.addMethod = function(method, func) {
    this[method] = func;
  };
}

let range = {
  from: 1,
  to: 5,
  [Symbol.iterator]: function() {
    return {
      current: this.from,
      final: this.to,
      next: function() {
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
  return function(item) {
    return low <= item && item <= high;
  };
}

function inArray(array) {
  return function(item) {
    return array.includes(item);
  };
}

function byField(fieldName) {
  return function(a, b) {
    return a[fieldName] > b[fieldName] ? 1 : -1;
  };
}

function makeArmy() {
  let shooters = [];

  function makeShooter(i) {
    let shooter = function() {
      // create a shooter function,
      // alert(i); // that should show its number
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

function makeCounter() {
  let count = 0;

  counter.set = function(val) {
    count = val;
  };

  counter.decrease = function() {
    count--;
  };

  function counter() {
    return ++count;
  }

  return counter;
}

let counter = makeCounter();

function sum(a) {
  let sum = a;

  function nestedFunc(b) {
    sum += b;
    return nestedFunc;
  }

  nestedFunc[Symbol.toPrimitive] = function() {
    return sum;
  };

  return nestedFunc;
}

function throttle(func, ms) {
  let lastArgs, lastThis;
  let isThrottled = false;

  return function wrapper() {
    if (isThrottled) {
      lastArgs = arguments;
      lastThis = this;
      return;
    }

    func.apply(this, arguments);
    isThrottled = true;

    setTimeout(() => {
      isThrottled = false;
      if (lastArgs) {
        wrapper.apply(lastThis, lastArgs);
        lastArgs = lastThis = null;
      }
    }, ms);
  };
}

// loadJson("https://javascript.info/no-such-user.json").catch(console.log);
class HttpError extends Error {
  constructor(response) {
    super(`${response.status} for ${response.url}`);
    this.name = "HttpError";
    this.response = response;
  }
}

async function loadJson(url) {
  let response = await fetch(url);
  if (response.status == 200) {
    return await response.json();
  } else {
    throw new HttpError(response.status);
  }
}

// Ask for a user name until github returns a valid user
async function demoGithubUser() {
  while (true) {
    try {
      let name = prompt("Enter a name?", "iliakan");
      let user = await loadJson(`https://api.github.com/users/${name}`);
      console.log(`Full name: ${user.name}.`);
      return user;
    } catch (err) {
      if (err instanceof HttpError && err.response.status == 404) {
        console.log("No such user, please reenter.");
      } else {
        throw err;
      }
    }
  }
}

// demoGithubUser();
async function wait() {
  await new Promise((resolve) => setTimeout(resolve, 1000));

  return 10;
}

function f() {
  // ...what should you write here?
  // we need to call async wait() and wait to get 10
  // remember, we can't use "await"
  new Promise((res) => {
    res(wait());
  }).then((v) => console.log(v));
}
