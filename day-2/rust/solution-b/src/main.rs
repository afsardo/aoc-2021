use std::fs;

fn main() {
    let contents =
        fs::read_to_string("../../input.txt").expect("Something went wrong reading the file");

    let mut position_y = 0;
    let mut position_x = 0;
    let mut aim = 0;

    for line in contents.lines() {
        let line_split = line.split(' ');
        let action = line_split.clone().nth(0).unwrap();
        let value = line_split.clone().nth(1).unwrap().parse::<i32>().unwrap();

        if action.to_owned().eq("forward") {
            position_x += value;
            position_y += aim * value;
        }

        if action.to_owned().eq("down") {
            aim += value;
        }

        if action.to_owned().eq("up") {
            aim -= value;
        }
    }

    println!(
        "Position: ({}, {}) = {}",
        position_x,
        position_y,
        position_x * position_y
    );
}
