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
var ans = Infinity;
for (var i = 1; i <= n; i++) {
    s.push(input[i - 1]);
}
console.log(s);
