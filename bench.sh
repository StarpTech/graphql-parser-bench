echo -n "async-graphql-parser: " && cd async-graphql-parser && cargo build --release -q && ./target/release/async-graphql-parser $@
echo -n "graphql-parser: " && cd .. && cd graphql-parser && cargo build --release -q && ./target/release/graphql-parser $@
echo -n "graphql-go-tools: " && cd .. && cd graphql-go-tools && go build -ldflags "-s -w" -o test && ./test $@
echo -n "go-graphql: " && cd .. && cd go-graphql && go build -ldflags "-s -w" -o test && ./test $@
echo -n "gqlparser: " && cd .. && cd gqlparser && go build -ldflags "-s -w" -o test && ./test $@
cd .. && cd graphqljs && node ./index.mjs $@