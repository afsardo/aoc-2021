export const solution = "a";

const input = await Bun.file("../input.txt").text();

let positionX = 0;
let positionY = 0;

for (const line of input.split("\n")) {
  const [action, value] = line.split(" ");

  if (action === "forward") {
    positionX += Number(value);
  }

  if (action === "down") {
    positionY += Number(value);
  }

  if (action === "up") {
    positionY -= Number(value);
  }
}

console.log('Solution:', positionX * positionY);