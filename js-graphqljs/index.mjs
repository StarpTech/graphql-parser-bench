import { buildSchema } from 'graphql'
import { join } from 'path'
import { readFile } from 'fs/promises'

const schema = await readFile(join('..', 'schema.graphql'), 'utf-8')

// warmup (jit)

for (let index = 0; index < 10; index++) {
    const hrstart = process.hrtime()
    buildSchema(schema)
    if (index === 0) {
        const hrend = process.hrtime(hrstart)
        console.info('graphqljs (cold): %dms', hrend[1] / 1000000)
    }

}

const hrstart = process.hrtime()

buildSchema(schema)

const hrend = process.hrtime(hrstart)
console.info('graphqljs (warmed): %dms', hrend[1] / 1000000)