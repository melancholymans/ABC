var file = process.platform === 'linux' ? '/dev/stdin' : './input.txt';
var _a = require('fs')
    .readFileSync(file, 'utf8')
    .trim()
    .split('\n')
    .map(function (r) {
    return r
        .trim()
        .split(/\s+/)
        .map(function (i) { return (isNaN(i) ? i : i - 0); });
}), n = _a[0][0], input = _a[1];
var s = [];
var sum1 = 0;
var sum2 = 0;
var ans = Infinity;
for (var i = 1; i <= n; i++) {
    s.push(input[i - 1]);
    if (input.slice(i).length > 0) {
        sum2 = input.slice(i).reduce(function (a, b) { return a + b; });
    }
    sum1 = s.reduce(function (a, b) { return a + b; });
    if (ans > Math.abs(sum1 - sum2))
        ans = Math.abs(sum1 - sum2);
}
console.log(ans);
