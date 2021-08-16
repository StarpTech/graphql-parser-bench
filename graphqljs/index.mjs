import { readFile } from 'fs/promises'
import { parse } from 'graphql'
import { performance } from 'perf_hooks'

const schema = await readFile(process.argv[2], 'utf-8')

// warmup (jit)

for (let index = 0; index < 10; index++) {
  const start = performance.now()
    parse(schema)
    if (index === 0) {
        const end = performance.now()
        console.info(`graphqljs (cold): ${end - start}ms`)
    }

}

const start = performance.now()

parse(schema)

const end = performance.now()
console.info(`graphqljs (warmed): ${end - start}ms`, )