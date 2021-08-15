import { buildSchema, parse } from 'graphql'
import { readFile } from 'fs/promises'

const schema = await readFile(process.argv[2], 'utf-8')

// warmup (jit)

for (let index = 0; index < 10; index++) {
    const hrstart = process.hrtime()
    parse(schema)
    if (index === 0) {
        const hrend = process.hrtime(hrstart)
        console.info('graphqljs (cold): %dms', hrend[1] / 1000000)
    }

}

const hrstart = process.hrtime()

parse(schema)

const hrend = process.hrtime(hrstart)
console.info('graphqljs (warmed): %dms', hrend[1] / 1000000)