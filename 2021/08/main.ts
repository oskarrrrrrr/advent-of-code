import _ from "lodash"

export { }

async function readInput(): Promise<[string[][], string[][]]> {
    const text = await Bun.file("input.txt").text()
    const rawInput: string = text.trim()
    let [signals, outputs] = [[], []]
    for (const line of rawInput.split("\n")) {
        let [signalStr, outputStr] = line.split(" | ")
        signals.push(signalStr.split(" "))
        outputs.push(outputStr.split(" "))
    }
    return [signals, outputs]
}

function part1(outputs: string[][]): number {
    let count = 0;
    for (const out of outputs) {
        for (const digit of out) {
            if ([2, 3, 4, 7].includes(digit.length)) {
                count++;
            }
        }
    }
    return count
}

function allDistinct(num: string) {
    for (let i = 0; i < num.length - 1; i++) {
        for (let j = i + 1; j < num.length; j++) {
            if (num[i] == num[j]) {
                return false
            }
        }
    }
    return true
}

function segmentsInList(num: string, order: string, expectedSegments: number[]): boolean {
    for (const n of num) {
        if (!expectedSegments.includes(order.indexOf(n))) {
            return false
        }
    }
    return true
}

function getNumber(num: string, order: string = "abcdefg"): number {
    if (num.length < 2 || num.length > 7) {
        throw new Error("Invalid number of digits: " + num.length)
    }
    if (!allDistinct(num)) {
        throw new Error("Duplicate segments in the digit: " + num)
    }
    if (num.length == 2) {
        if (segmentsInList(num, order, [2, 5])) {
            return 1
        }
        return -1
    }
    if (num.length == 3) {
        if (segmentsInList(num, order, [0, 2, 5])) {
            return 7
        }
        return -1
    }
    if (num.length == 4) {
        if (segmentsInList(num, order, [1, 2, 3, 5])) {
            return 4
        }
        return -1

    }
    if (num.length == 5) {
        if (segmentsInList(num, order, [0, 2, 3, 4, 6])) {
            return 2
        }
        if (segmentsInList(num, order, [0, 2, 3, 5, 6])) {
            return 3
        }
        if (segmentsInList(num, order, [0, 1, 3, 5, 6])) {
            return 5
        }
        return -1
    }
    if (num.length == 6) {
        if (segmentsInList(num, order, [0, 1, 2, 4, 5, 6])) {
            return 0
        }
        if (segmentsInList(num, order, [0, 1, 3, 4, 5, 6])) {
            return 6
        }
        if (segmentsInList(num, order, [0, 1, 2, 3, 5, 6])) {
            return 9
        }
        return -1
    }
    if (num.length == 7) {
        return 8
    }
    throw new Error()
}

function* stringPerms(s: string): Generator<string> {
    if (s.length <= 1) {
        yield s
        return
    }
    for (let i = 0; i < s.length; i++) {
        if (s.indexOf(s[i]) != i) { continue } // in case of duplicates in chars
        const rest = s.slice(0, i) + s.slice(i + 1, s.length);
        for (const subPerm of stringPerms(rest)) {
            yield s[i] + subPerm;
        }
    }
}

function part2(signals: string[][], outputs: string[][]): number {
    let sum = 0
    for (let i = 0; i < signals.length; i++) {
        for (const order of stringPerms("abcdefg")) {
            let allGood = true;
            for (const num of signals[i]) {
                if (getNumber(num, order) == -1) {
                    allGood = false
                    break
                }
            }
            if (!allGood) {
                continue
            }
            for (const num of outputs[i]) {
                if (getNumber(num, order) == -1) {
                    allGood = false
                    break
                }
            }
            if (allGood) {
                let result = 0
                for (const d of outputs[i]) {
                    result *= 10
                    result += getNumber(d, order)
                }
                sum += result
                break
            }
        }
    }
    return sum
}

async function main() {
    const [signals, outputs] = await readInput()
    console.log("[1]", part1(outputs))
    console.log("[2]", part2(signals, outputs))
}

main()
