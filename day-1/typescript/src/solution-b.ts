export const solution = "b";

const input = await Bun.file("../input.txt").text();

const numbers = input.split("\n").map(Number);

let increments = 0;
let lastWindow = Infinity;

for (let i = 0; i < numbers.length; i++) {
  if (numbers.slice(i, i + 3).length < 3) break;
  const window = numbers.slice(i, i + 3).reduce((a, b) => a + b, 0);
  if (lastWindow < window) {
    increments++;
  }
  lastWindow = window;
}

console.log('Solution:', increments);