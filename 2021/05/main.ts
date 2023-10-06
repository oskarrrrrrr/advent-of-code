import _ from "lodash"

export { }

class Line {
    x1: number
    y1: number
    x2: number
    y2: number

    constructor(x1: number, y1: number, x2: number, y2: number) {
        this.x1 = x1
        this.y1 = y1
        this.x2 = x2
        this.y2 = y2
    }
}

class Map {
    map: number[][]

    constructor(width: number, height: number) {
        this.map = []
        for (let y = 0; y < height; y++) {
            this.map.push([])
            for (let x = 0; x < width; x++) {
                this.map[y].push(0)
            }
        }
    }

    applyLine(line: Line) {
        if (line.x1 == line.x2 || line.y1 == line.y2) {
            // vertical and horizontal lines
            const [x1, x2] = [Math.min(line.x1, line.x2), Math.max(line.x1, line.x2)]
            const [y1, y2] = [Math.min(line.y1, line.y2), Math.max(line.y1, line.y2)]
            for (let y = y1; y <= y2; y++) {
                for (let x = x1; x <= x2; x++) {
                    this.map[y][x]++
                }
            }
        } else {
            // 45 degrees lines
            let [x, y] = [line.x1, line.y1]
            const dx = (line.x1 > line.x2) ? -1 : 1
            const dy = (line.y1 > line.y2) ? -1 : 1
            while (x != line.x2 && y != line.y2) {
                this.map[y][x]++
                x += dx
                y += dy
            }
            this.map[y][x]++
        }
    }

    str(): string {
        let s = ""
        for (const row of this.map) {
            for (const v of row) {
                if (v == 0) {
                    s += "."
                } else {
                    s += v.toString()
                }
            }
            s += "\n"
        }
        return s
    }
}

async function readInput(): Promise<Line[]> {
    const inputFile = Bun.file("input.txt")
    const rawInput: string = await inputFile.text()
    const input = rawInput.trim()
    const textLines = input.split("\n")
    let lines: Line[] = []
    for (const line of textLines) {
        if (line == "") {
            continue
        }
        const [from, to] = line.split(" -> ")
        const [x1, y1] = from.split(",").map(_.toInteger)
        const [x2, y2] = to.split(",").map(_.toInteger)
        const newLine = new Line(x1, y1, x2, y2)
        lines.push(newLine)
    }
    return lines
}

function dimsFromLines(lines: Line[]): [number, number] {
    let width = 0
    let height = 0
    for (const line of lines) {
        width = Math.max(width, line.x1, line.x2)
        height = Math.max(width, line.y1, line.y2)
    }
    return [width + 1, height + 1]
}

function countOverlaps(map: Map): number {
    let count = 0
    for (const row of map.map) {
        for (const v of row) {
            if (v > 1) {
                count++
            }
        }
    }
    return count
}

function part1(lines: Line[]): number {
    const [width, height] = dimsFromLines(lines)
    const map = new Map(width, height)
    for (const line of lines) {
        if (line.x1 != line.x2 && line.y1 != line.y2) {
            continue
        }
        map.applyLine(line)
    }
    return countOverlaps(map)
}

function part2(lines: Line[]): number {
    const [width, height] = dimsFromLines(lines)
    const map = new Map(width, height)
    for (const line of lines) {
        map.applyLine(line)
    }
    return countOverlaps(map)
}

async function main() {
    const lines = await readInput()
    console.log("[1]", part1(lines))
    console.log("[2]", part2(lines))
}

main()
