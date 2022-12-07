use std::io;
use std::io::BufRead;

fn main() {
    let mut lines = io::stdin().lock().lines();

    let mut max_cal = 0;
    let mut sum = 0;

    while let Some(line_res) = lines.next() {
        let line = line_res.unwrap();

        if line == "" {
            if sum > max_cal {
                max_cal = sum;
            }
            sum = 0;
            continue;
        }

        sum += line.parse::<i32>().unwrap();
    }

    println!("{}", max_cal)
}
