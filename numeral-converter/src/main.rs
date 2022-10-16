use std::io;

fn main() {
    println!("A numeral system converter");
    converter();
    let mut number = String::new();
    io::stdin().read_line(&mut number).expect("Error reading");
    println!("{}",number);
}

fn converter() {
    let mut number = String::new();
    io::stdin().read_line(&mut number).expect("Error reading");
    println!("{}",number);
}