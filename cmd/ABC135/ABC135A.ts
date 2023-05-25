export{}

function Main(input:string){
    const list=input.split("\n");
    const a:number=parseInt(list[0]);
    const b:number=parseInt(list[1]);
    console.log(`${a} ${b}`)
}
//Main(require("fs").readFileSync("/dev/stdin","utf8"));
