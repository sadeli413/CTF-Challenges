use std::{cmp::Ordering, io};

use rand::Rng;

mod guess;

fn main() {
    // Generate a random number
    println!("Guess the number!");
    let secret_number = rand::thread_rng().gen_range(1..=100);
    println!("The secret number is: {}", secret_number);

    // Infinite loop
    loop {
        println!("Please input your guess.");

        // Get user input
        let mut ig = String::new();
        io::stdin().read_line(&mut ig).expect("Failed to read line");

        // Convert to i32
        let ig: i32 = match ig.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        let g = guess::Guess::new(ig);

        // Compare number
        println!("You guessed: {}", g.value());
        match g.value().cmp(&secret_number) {
            Ordering::Less => println!("To small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
