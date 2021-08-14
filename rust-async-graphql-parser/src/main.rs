use std::fs;
use std::time::Instant;

fn main() {
    let contents = fs::read_to_string("./../schema.graphql").unwrap();
    let start = Instant::now();
    async_graphql_parser::parse_schema(contents).unwrap();
    let duration = start.elapsed();
    println!("async-graphql-parser: {:?}", duration);
}
