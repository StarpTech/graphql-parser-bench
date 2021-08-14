# graphql-parser-bench

Parsing a schema or document can be a critical part of the application, especially if you have to care about latency. This is the case for proxies or caches where high load is exptected. This benchmark should outline good candidates. 

## Benchmark

This test requires Rust, Go and Node.js (>=14)

```
./bench.sh
```

Parse the [gitlab schema](./schema.graphql). The schema has a size of around `671KB`.

## Intermediate result

- ✅ graphql-go-tools: 5.7179ms
- ✅ graphql-parser: 7.4608ms
- ✅ async-graphql-parser: 12.7027ms
- ✅ graphqljs (cold): 111.5407ms
- ✅ graphqljs (warmed): 47.7197ms
- ❌ go-graphql: [Error](https://github.com/graphql-go/graphql/issues/611)

**🏆 Winner**: In terms of speed `graphql-go-tools` wins!

Compared to different aspect like accessibility and maintainability `async-graphql-parser` is the overall winner. It uses a PEG grammar as input.