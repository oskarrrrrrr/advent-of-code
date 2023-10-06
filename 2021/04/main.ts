import _ from "lodash"

export { }

class Board {
    board: number[][]
    marked: boolean[][]

    constructor(board: number[][]) {
        this.board = board
        this.marked = []
        for (let i = 0; i < board.length; i++) {
            const row: boolean[] = []
            for (let j = 0; j < board[i].length; j++) {
                row.push(false)
            }
            this.marked.push(row)
        }
    }

    mark(num: number): void {
        for (let i = 0; i < this.board.length; i++) {
            for (let j = 0; j < this.board[i].length; j++) {
                if (this.board[i][j] == num) {
                    this.marked[i][j] = true
                }
            }
        }
    }

    done(): boolean {
        // rows
        for (let i = 0; i < this.board.length; i++) {
            let count = 0
            for (let j = 0; j < this.board[i].length; j++) {
                if (this.marked[i][j]) {
                    count++
                } else {
                    break
                }
            }
            if (count == this.board[i].length) {
                return true
            }
        }
        // columns
        for (let i = 0; i < this.board[0].length; i++) {
            let count = 0
            for (let j = 0; j < this.board.length; j++) {
                if (this.marked[j][i]) {
                    count++
                } else {
                    break
                }
            }
            if (count == this.board.length) {
                return true
            }
        }
        return false
    }

    score(lastNum: number): number {
        let sum = 0
        for (let i = 0; i < this.board.length; i++) {
            for (let j = 0; j < this.board[i].length; j++) {
                if (!this.marked[i][j]) {
                    sum += this.board[i][j]
                }
            }
        }
        return sum * lastNum
    }
}

async function readInput(): Promise<[number[], Board[]]> {
    const inputFile = Bun.file("input.txt")
    const rawInput: string = await inputFile.text()
    const input = rawInput.trim()
    const lines = input.split("\n")
    const header = lines[0].split(",").map(_.toInteger)
    let boards: Board[] = []
    let currBoard: number[][] = []
    for (let i = 2; i < lines.length; i++) {
        if (lines[i] == "") {
            const newBoard = new Board(currBoard)
            currBoard = []
            boards.push(newBoard)
        } else {
            const lineNums = lines[i]
                .split(" ")
                .map(_.trim)
                .filter(s => s != "")
                .map(_.toInteger)
            currBoard.push(lineNums)
        }
    }
    boards.push(new Board(currBoard))
    return [header, boards]
}

async function part1(): Promise<number> {
    const [nums, boards] = await readInput()
    for (const num of nums) {
        for (const board of boards) {
            board.mark(num)
            if (board.done()) {
                return board.score(num)
            }
        }
    }
    throw new Error("No board won in time.")
}

async function part2(): Promise<number> {
    const [nums, boards] = await readInput()
    let wins = 0
    for (const num of nums) {
        for (const board of boards) {
            if (board.done()) {
                continue
            }
            board.mark(num)
            if (board.done()) {
                wins++
                if (wins == boards.length) {
                    return board.score(num)
                }
            }
        }
    }
    throw new Error("Some boards didn't win in time.")
}

async function main() {
    console.log("[1]", await part1())
    console.log("[2]", await part2())
}

main()
