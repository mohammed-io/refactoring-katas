class Calculator {
  constructor() {}

  normalize(a, b, c) {
    let d = [];
    let e = Infinity;
    let f = -Infinity;

    for (let g = 0; g < a.length; g++) {
      if (a[g] < e) e = a[g];
      if (a[g] > f) f = a[g];
    }

    let h = f - e;
    let i = c - b;

    for (let g = 0; g < a.length; g++) {
      let j = (a[g] - e) / h;
      let k = b + j * i;
      d.push(Math.round(k * 100) / 100);
    }

    return d;
  }
}

export { Calculator };
