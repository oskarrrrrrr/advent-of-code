import _ from "lodash"

export { }

async function readInput(): Promise<string[]> {
    const inputFile = Bun.file("input.txt")
    const rawInput: string = await inputFile.text()
    const input = rawInput.trim()
    return input.split("\n")
}

function arrayToInteger(bits: number[]): number {
    let result = 0
    for (let i = 0; i < bits.length; i++) {
        if (bits[i] == 1) {
            result |= 1 << (bits.length - i - 1)
        }
    }
    return result
}

function stringToInteger(bits: string): number {
    let result = 0
    for (let i = 0; i < bits.length; i++) {
        if (bits[i] == '1') {
            result |= 1 << (bits.length - i - 1)
        }
    }
    return result
}

function part1(input: string[]): number {
    let bitsCounts: number[] = new Array(input[0].length)
    for (let i = 0; i < input[0].length; i++) {
        bitsCounts[i] = 0
    }

    for (const num of input) {
        for (let di = 0; di < num.length; di++) {
            switch (num[di]) {
                case '1':
                    bitsCounts[di]++
                    break
                case '0':
                    bitsCounts[di]--
                    break
                default:
                    throw new Error("Unexpected digit value.")
            }
        }
    }

    const gammaBits = bitsCounts.map(x => x > 0 ? 1 : 0)
    const gamma = arrayToInteger(gammaBits)
    const epsilonBits = bitsCounts.map(x => x < 0 ? 1 : 0)
    const epsilon = arrayToInteger(epsilonBits)

    return gamma * epsilon
}

function countNthDigit(nums: Set<string>, n: number): number {
    let count = 0
    for (const num of nums) {
        if (num[n] == '1') {
            count++
        } else {
            count--
        }
    }
    return count
}

function getOxygenRating(input: string[]): number {
    let oxygenNums: Set<string> = new Set()
    for (const num of input) {
        oxygenNums.add(num)
    }
    for (let i = 0; i < input[0].length; i++) {
        const count = countNthDigit(oxygenNums, i)
        const newOxygenNums: Set<string> = new Set()
        if (count >= 0) {
            for (const num of oxygenNums) {
                if (num[i] == '1') {
                    newOxygenNums.add(num)
                }
            }
        } else {
            for (const num of oxygenNums) {
                if (num[i] == '0') {
                    newOxygenNums.add(num)
                }
            }
        }
        oxygenNums = newOxygenNums
        if (oxygenNums.size == 1) {
            break
        }
    }
    const [oxygenNum] = oxygenNums
    return stringToInteger(oxygenNum)
}

function getCo2Rating(input: string[]): number {
    let co2Nums: Set<string> = new Set()
    for (const num of input) {
        co2Nums.add(num)
    }
    for (let i = 0; i < input[0].length; i++) {
        const count = countNthDigit(co2Nums, i)
        const newCo2Nums: Set<string> = new Set()
        if (count >= 0) {
            for (const num of co2Nums) {
                if (num[i] == '0') {
                    newCo2Nums.add(num)
                }
            }
        } else {
            for (const num of co2Nums) {
                if (num[i] == '1') {
                    newCo2Nums.add(num)
                }
            }
        }
        co2Nums = newCo2Nums
        if (co2Nums.size == 1) {
            break
        }
    }
    const [co2Num] = co2Nums
    return stringToInteger(co2Num)
}

function part2(input: string[]): number {
    return getOxygenRating(input) * getCo2Rating(input)
}

async function main() {
    const input = await readInput()
    console.log("[1]", part1(input))
    console.log("[2]", part2(input))
}

main()
