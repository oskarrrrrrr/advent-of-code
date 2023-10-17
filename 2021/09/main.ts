import _ from "lodash"

export { }

async function readInput(): Promise<number[][]> {
    const text = await Bun.file("input.txt").text()
    const rawInput: string = text.trim()
    let nums: number[][] = []
    for (const line of rawInput.split("\n")) {
        const row: number[] = []
        for (const d of line) {
            row.push(_.toInteger(d))
        }
        nums.push(row)
    }
    return nums
}

function part1(nums: number[][]): number {
    let result = 0
    for (let y = 0; y < nums.length; y++) {
        for (let x = 0; x < nums[y].length; x++) {
            const curr = nums[y][x]
            if (
                (y > 0 && nums[y - 1][x] <= curr)
                || (x > 0 && nums[y][x - 1] <= curr)
                || (y < nums.length - 1 && nums[y + 1][x] <= curr)
                || (x < nums[y].length - 1 && nums[y][x + 1] <= curr)
            ) {
                // not a low point
            } else {
                // q low point
                result += curr + 1
            }
        }
    }
    return result
}

function printBasins(basins: number[][]): void {
    let s = ""
    for (const row of basins) {
        for (const d of row) {
            s += " "
            if (d < 10) {
                s += "0"
            }
            if (d < 100) {
                s += "0"
            }
            s += d
        }
        s += "\n"
    }
    console.log(s)
}

function part2(nums: number[][]): number {
    // calculate rough basin numbers
    let basins: number[][] = []
    let basinNum = 0
    for (let y = 0; y < nums.length; y++) {
        basins.push([])
        for (let x = 0; x < nums[y].length; x++) {
            if (nums[y][x] == 9) {
                basins[y].push(0)
            } else if (y > 0 && basins[y - 1][x] != 0) {
                basins[y].push(basins[y - 1][x])
            } else if (x > 0 && basins[y][x - 1] != 0) {
                basins[y].push(basins[y][x - 1])
            } else {
                basinNum++
                basins[y].push(basinNum)
            }
        }
    }

    // unify neighbouring basins
    let neighbours: [number, number][] = []
    do {
        neighbours = []
        for (let y = 0; y < nums.length; y++) {
            for (let x = 0; x < nums[y].length; x++) {
                if (basins[y][x] == 0) {
                    continue
                }
                if (y > 0 && basins[y - 1][x] > 0 && basins[y - 1][x] != basins[y][x]) {
                    const min = Math.min(basins[y - 1][x], basins[y][x])
                    const max = Math.max(basins[y - 1][x], basins[y][x])
                    neighbours.push([min, max])
                }
                if (x > 0 && basins[y][x - 1] > 0 && basins[y][x - 1] != basins[y][x]) {
                    const min = Math.min(basins[y][x - 1], basins[y][x])
                    const max = Math.max(basins[y][x - 1], basins[y][x])
                    neighbours.push([min, max])
                }
            }
        }
        neighbours.sort(function([min1, max1], [min2, max2]) {
            if (min1 == min2) {
                return max2 - max1
            }
            return min2 - min1
        })

        // unify basins (some basins got assigned multiple numbers)
        for (let y = 0; y < basins.length; y++) {
            for (let x = 0; x < basins[y].length; x++) {
                for (const [min, max] of neighbours) {
                    if (max == basins[y][x]) {
                        basins[y][x] = min
                    }
                }
            }
        }
    } while (neighbours.length > 0)

    // count result
    let basinSizes: [number, number][] = []
    for (let y = 0; y < basins.length; y++) {
        for (let x = 0; x < basins[y].length; x++) {
            if (basins[y][x] == 0) {
                continue
            }
            let assigned = false
            for (let i = 0; i < basinSizes.length; i++) {
                const [basinNo, count] = basinSizes[i]
                if (basinNo == basins[y][x]) {
                    basinSizes[i] = [basinNo, count + 1]
                    assigned = true
                    break
                }
            }
            if (!assigned) {
                basinSizes.push([basins[y][x], 1])
            }
        }
    }
    basinSizes.sort(([n1, c1], [n2, c2]) => c2 - c1)

    let result = 1
    for (let i = 0; i < 3 && i < basinSizes.length; i++) {
        const [_, size] = basinSizes[i]
        result *= size
    }

    return result
}

async function main() {
    const nums = await readInput()
    console.log("[1]", part1(nums))
    console.log("[2]", part2(nums))
}

main()
