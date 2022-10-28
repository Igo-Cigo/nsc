use std::io;
use std::process;
extern crate text_io;

fn main() {
    println!("A numeral system converter");
    converter();
}

fn converter() {

    let stdin = io::stdin();
    let mut number = String::new();
    let mut ns = String::new();

    println!("Please enter a number: ");
    stdin.read_line(&mut number).expect("Error reading");
    println!("To what numeral system do you want to convert? (binary - B, octal - O, hexadecimal - H; quit - q)");
    stdin.read_line(&mut ns).expect("Error reading");
    ns = ns.trim().to_string();
    let conv_number = number.trim();
    match conv_number.parse::<u32>() {
        Ok(i) => println!("{}", i),
        Err(..) => println!("{}", conv_number),
    };
    let i: i32 = number.parse().unwrap();
    if ns == "q" {
        process::exit(0);
    }
    else if ns == "b" {
        println!("{:b}", i);
    };
    return converter()
}