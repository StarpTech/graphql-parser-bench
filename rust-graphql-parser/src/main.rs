use std::fs;
use std::time::Instant;

fn main() {
    let contents = fs::read_to_string("./../schema.graphql").unwrap();
    let start = Instant::now();
    let _ast = graphql_parser::parse_schema::<String>(&contents);
    let duration = start.elapsed();
    println!("graphql-parser: {:?}", duration);
}
