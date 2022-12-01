export const solution = "a";

const input = await Bun.file("../input.txt").text();

let increments = 0;
let lastNumber = Infinity;

for (const number of input.split("\n")) {
  if (lastNumber < Number(number)) {
    increments++;
  }
  lastNumber = Number(number);
}

console.log('Solution:', increments);