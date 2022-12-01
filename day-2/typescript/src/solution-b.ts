export const solution = "b";

const input = await Bun.file("../input.txt").text();

let positionX = 0;
let positionY = 0;
let aim = 0;

for (const line of input.split("\n")) {
  const [action, value] = line.split(" ");

  if (action === "forward") {
    positionX += Number(value);
    positionY += aim * Number(value);
  }

  if (action === "down") {
    aim += Number(value);
  }

  if (action === "up") {
    aim -= Number(value);
  }
}

console.log('Solution:', positionX * positionY);