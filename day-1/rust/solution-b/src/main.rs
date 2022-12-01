use std::fs;

fn main() {
    let contents =
        fs::read_to_string("../../input.txt").expect("Something went wrong reading the file");

    let mut increments = 0;
    let mut last_window = std::i32::MAX;

    let lines = contents.lines().collect::<Vec<&str>>();

    for i in 0..lines.len() {
        if lines.len() < i + 3 {
            break;
        }

        let window = lines[i..i + 3]
            .iter()
            .map(|line| line.parse::<i32>().unwrap())
            .reduce(|a, b| a + b)
            .unwrap();

        if last_window < window {
            increments += 1;
        }
        last_window = window;
    }

    println!("Solution: {}", increments);
}
