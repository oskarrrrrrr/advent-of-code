import _ from "lodash"

export {};

async function readInput(): Promise<number[]> {
    const inputFile = Bun.file("input.txt")
    const rawInput: string = await inputFile.text()
    const input = rawInput.trim()
    let nums: number[] = []
    for (const numStr of input.split("\n")) {
        nums.push(_.toInteger(numStr))
    }
    return nums
}

function part1(nums: number[]): number {
    let incCount = 0
    for (let i = 1; i < nums.length; i++) {
        if (nums[i] > nums[i-1]) {
            incCount++
        }
    }
    return incCount
}

function part2(nums: number[]): number {
    let incCount = 0
    let slidingWindow = nums[0] + nums[1] + nums[2]
    for (let i = 3; i < nums.length; i++) {
        const newSlidingWindow = slidingWindow - nums[i-3] + nums[i]
        if (newSlidingWindow > slidingWindow) {
            incCount++
        }
    }
    return incCount
}

async function main(): Promise<void> {
    const nums = await readInput()
    console.log("[1]", part1(nums))
    console.log("[2]", part2(nums))
}

main();
