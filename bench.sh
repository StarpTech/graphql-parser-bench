cd rust-async-graphql-parser && cargo build --release && ./target/release/rust-async-graphql-parser
cd .. && cd go-graphql-go-tools && go build -ldflags "-s -w" && ./go-graphql-go-tools
cd .. && cd js-graphqljs && node ./index.mjs