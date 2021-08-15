echo -n "async-graphql-parser: " && cd rust-async-graphql-parser && cargo build --release -q && ./target/release/rust-async-graphql-parser $@
echo -n "graphql-parser: " && cd .. && cd rust-graphql-parser && cargo build --release -q && ./target/release/rust-graphql-parser $@
echo -n "graphql-go-tools: " && cd .. && cd graphql-go-tools && go build -ldflags "-s -w" -o test && ./test $@
echo -n "go-graphql: " && cd .. && cd go-graphql && go build -ldflags "-s -w" -o test && ./test $@
echo -n "gqlparser: " && cd .. && cd gqlparser && go build -ldflags "-s -w" -o test && ./test $@
cd .. && cd graphqljs && node ./index.mjs $@