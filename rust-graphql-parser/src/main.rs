use std::env;
use std::fs;
use std::time::Instant;

fn main() {
    let args: Vec<String> = env::args().collect();
    let contents = fs::read_to_string(&args[1]).unwrap();
    let start = Instant::now();
    let _ast = graphql_parser::parse_schema::<String>(&contents);
    let duration = start.elapsed();
    println!("graphql-parser: {:?}", duration);
}
