use std::fs;

fn main() {
    let contents =
        fs::read_to_string("../../input.txt").expect("Something went wrong reading the file");

    let mut increments = 0;
    let mut last_number = std::i32::MAX;

    for line in contents.lines() {
        let number = line.parse::<i32>().unwrap();

        if last_number < number {
            increments += 1;
        }
        last_number = number;
    }

    println!("Solution: {}", increments);
}
