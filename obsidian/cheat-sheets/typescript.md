---
title: 🎃 Typescript
tags: [typescript]
draft: false
summary: "Typescript is 😖"

# PaperMod
ShowToc: true
TocOpen: true

date: 2022-12-17T21:41:33+01:00
---

## General notes

`{ id }` is sugar for `{ id: id }` which is sugar for `{ "id": id }`

## Tests

### Limit tests to be executed

Keywords `describe` and `it` are globally provided by the mocha test suite.

```ts
// npm run test:watch

// Use "only" to run a limited set of tests
describe.only() {
  it.only() {}
}
```

## Snippets

### List comprehension

```ts
// List comprehension or map
profiles.map((profile) => profile.id)  // python code equivalent: profile.id for profile in profiles

```

### Debugging with print

```ts
console.log(someObject)

console.dir({ someObject }, { depth: 5 })

throw new Error(JSON.stringify({ profiles: response.profiles, users: response.users }))
// Use JSON.parse to "load"
```

### Fat arrow

```ts
function x(a: number, b: number): number {
    return a + b
}
console.log(x(1,2))

const y = (a: number, b:number) => {return a + b}
console.log(y(2,2))
```

### Type casting

`as` and `<>` works as as casting operators.

### JSON parsing

```ts
const myObject = JSON.parse("{foo: bar}")

const myString: string = JSON.stringify(myObject)

// 
const invalidStringifiedObj: string = "foo"
const default_to_empty_obj = JSON.parse(invalidStringifiedObj | {})
```

