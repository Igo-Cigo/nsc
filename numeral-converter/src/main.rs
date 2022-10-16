use std::io;
use std::process;

fn main() {
    println!("A numeral system converter");
    converter();
}

fn converter() {
    let mut number = String::new();
    let mut ns = String::new();
    println!("Please enter a number: ");
    io::stdin().read_line(&mut number).expect("Error reading");
    println!("To what numeral system do you want to convert? (binary - B, octal - O, hexadecimal - H; quit - q)");
    io::stdin().read_line(&mut ns).expect("Error reading");
    if ns == "q" {
        process::exit(0);
    }
    else if ns == "b" {
        print!("HI");
    };
    return converter()
}