# graphql-parser-bench

Parsing a schema or document can be a critical part of the application, especially if you have to care about latency. This is the case for proxies or caches where high load is expected. This benchmark should outline good candidates. We use the libraries as documented in tests or documentation. We don't tune a specific parser. We assume that a new query arrives the parser at any time.

## Benchmark

This test requires Rust, Go and Node.js (>=14)

## Intermediate result

### Parsing a schema

Parse the [gitlab schema](./schema.graphql). The schema has a size of around `671KB`.

```sh
./bench.sh ./../schema.graphql schema
```

- ✅ graphql-go-tools: 5.7179ms
- ✅ graphql-parser: 7.4608ms
- ✅ gqlparser: 7.7642ms
- ✅ async-graphql-parser: 12.7027ms
- ✅ graphqljs (warmed): 11.1135ms
- ✅ graphqljs (cold): 39.505ms
- ❌ go-graphql: [Error](https://github.com/graphql-go/graphql/issues/611)

**🏆 Winner**: In terms of speed `graphql-go-tools` wins!

Compared to different aspect like accessibility and maintainability `async-graphql-parser` is the overall winner. It uses a PEG grammar as input.

### Parsing a query

Parse the [kitchen-sink document](./kitchen-sink.graphql). The document has a size of around `1KB` but with a wide usage of the specification.

```sh
./bench.sh ./../kitchen-sink.graphql query
```
- ✅ graphql-parser: 20.9µs
- ✅ gqlparser: 38.8µs
- ✅ graphql-go-tools: 42.6µs
- ✅ graphqljs (warmed): 0.2104ms
- ✅ graphqljs (cold): 1.4962ms
- ❌ async-graphql-parser: [Error](https://github.com/async-graphql/async-graphql/issues/602)
- ❌ go-graphql: [Error](https://github.com/graphql-go/graphql/issues/612)

**🏆 Winner**: In terms of speed `graphql-parser` wins! Alarming that some popular parser can not even parse the query.

## Contribution

Feel free to improve or add another parser to the list. PR's are welcome!