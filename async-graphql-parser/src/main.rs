use std::env;
use std::fs;
use std::time::Instant;

fn main() {
    let args: Vec<String> = env::args().collect();
    let contents = fs::read_to_string(&args[1]).unwrap();
    let start = Instant::now();
    if args[2] == "schema" {
        async_graphql_parser::parse_schema(contents).unwrap();
    } else {
        async_graphql_parser::parse_query(contents).unwrap();
    }
    let duration = start.elapsed();
    println!("{:?}", duration);
}
