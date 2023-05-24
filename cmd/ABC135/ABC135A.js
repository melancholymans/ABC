"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
function Main(input) {
    var list = input.split("\n");
    var a = parseInt(list[0]);
    var b = parseInt(list[1]);
    console.log("".concat(a, " ").concat(b));
}
Main(require("fs").readFileSync("/dev/stdin", "utf8"));
