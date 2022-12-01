#!/bin/bash

echo "Advent of Code 2021 - Day 2"

echo "============================"

cd ./typescript
echo "Running TypeScript Solutions"
echo "Part 1"
time npm run solution-a
echo "Part 2"
time npm run solution-b
cd ..

echo "============================"

cd ./go
echo "Running Go Solutions"
echo "Part 1"
go build solution-a.go >> /dev/null
time ./solution-a
echo "Part 2"
go build solution-b.go >> /dev/null
time ./solution-b
cd ..

echo "============================"

echo "Running Rust Solutions"
echo "Part 1"
cd ./rust/solution-a
cargo build --release -q >> /dev/null
time ./target/release/solution-a
cd ../..
echo "Part 2"
cd ./rust/solution-b
cargo build --release -q >> /dev/null
time ./target/release/solution-b
cd ../..

echo "============================"