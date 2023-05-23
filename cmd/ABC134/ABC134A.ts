const file = process.platform === 'linux' ? '/dev/stdin' : './input.txt';

let [[n], input] = require('fs')
  .readFileSync(file, 'utf8')
  .trim()
  .split('\n')
  .map(r =>
    r
      .trim()
      .split(/\s+/)
      .map(i => (isNaN(i) ? i : i - 0))
  );
 
let s = [];
let ans = Infinity;
for (let i = 1; i <= n; i++) {
  s.push(input[i - 1]);
}
console.log(s);
