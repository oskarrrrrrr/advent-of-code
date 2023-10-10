import _ from "lodash"

export { }

type Fish = Map<number, number>

function addFish(fish: Fish, n: number, c: number) {
    const curr = fish.get(n);
    if (curr === undefined) {
        fish.set(n, c)
    } else {
        fish.set(n, curr + c)
    }
}

function countFish(fish: Fish): number {
    let count = 0
    for (const c of fish.values()) {
        count += c
    }
    return count
}

function simDay(fish: Fish): Fish {
    const newFish = new Map()
    for (const [n, c] of fish) {
        if (n == 0) {
            addFish(newFish, 6, c)
            addFish(newFish, 8, c)
        } else {
            addFish(newFish, n - 1, c)
        }
    }
    return newFish
}

function simDays(fish: Fish, n: number): Fish {
    for (let i = 0; i < n; i++) {
        fish = simDay(fish)
    }
    return fish
}

async function readInput(): Promise<Fish> {
    const inputFile = Bun.file("input.txt")
    const rawInput: string = await inputFile.text()
    const input = rawInput.trim()
    const nums = input.split(",").map(_.toInteger)
    const fish = new Map()
    for (const n of nums) {
        addFish(fish, n, 1)
    }
    return fish
}

function part1(fish: Fish): number {
    fish = simDays(fish, 80)
    return countFish(fish)
}

function part2(fish: Fish): number {
    fish = simDays(fish, 256)
    return countFish(fish)
}

async function main() {
    const fish = await readInput()
    console.log("[1]", part1(fish))
    console.log("[2]", part2(fish))
}

main()
