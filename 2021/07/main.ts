import _ from "lodash"

export { }

async function readInput(): Promise<number[]> {
    const inputFile = Bun.file("input.txt")
    const rawInput: string = await inputFile.text()
    const input = rawInput.trim()
    return input.split(",").map(_.toInteger)
}

function median(nums: number[]): number {
    if (nums.length % 2 == 0) {
        return (nums[(nums.length / 2) - 1] + nums[nums.length / 2]) / 2
    }
    return nums[(nums.length - 1) / 2]
}

function part1(nums: number[]): number {
    nums = nums.toSorted((n1, n2) => n1 - n2)
    const m = median(nums)
    let count = 0
    for (const n of nums) {
        count += Math.abs(n - m);
    }
    return count;
}

function part2(nums: number[]): number {
    let minCost = Number.MAX_SAFE_INTEGER;
    for (let n = Math.min(...nums); n < Math.max(...nums); n++) {
        let cost = 0
        for (const m of nums) {
            const diff = Math.abs(n - m)
            cost += ((diff + 1) * diff) / 2
        }
        minCost = Math.min(minCost, cost)
    }
    return minCost
}

async function main() {
    const nums = await readInput()
    console.log("[1]", part1(nums))
    console.log("[2]", part2(nums))
}

main()
