# graphql-parser-bench

Parsing a schema or document can be a critical part of the application, especially if you have to care about latency. This is the case in proxies and CDN applications. This benchmark should outline good candidates. 

## Benchmark

This test requires Rust, Go and Node.js (>=14)

```
./bench.sh
```

Parse the [gitlab schema](./schema.graphql). The schema has a size of around `671KB`.

## Intermediate result

```
async-graphql-parser: 12.7027ms
graphql-parser: 7.4608ms
graphql-go-tools: 5.7179ms
graphqljs (cold): 111.5407ms
graphqljs (warmed): 47.7197ms
```

**Winner**: In speed `graphql-go-tools` wins but compared to different aspect like accessibility and maintainibility `async-graphql-parser` is the clear winner. It uses a PEG grammar as input.