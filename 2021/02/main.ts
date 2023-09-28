import _ from "lodash"

export { };

class Cmd {
    direction: string;
    value: number;

    constructor(direction: string, value: number) {
        this.direction = direction
        this.value = value
    }
}

async function readInput(): Promise<Cmd[]> {
    const inputFile = Bun.file("input.txt")
    const rawInput: string = await inputFile.text()
    const input = rawInput.trim()
    let cmds: Cmd[] = []
    for (const line of input.split("\n")) {
        const parts = line.split(" ")
        cmds.push(new Cmd(parts[0], _.toInteger(parts[1])))
    }
    return cmds
}

function part1(cmds: Cmd[]): number {
    let horizontal = 0
    let depth = 0
    for (const cmd of cmds) {
        switch (cmd.direction) {
            case "forward":
                horizontal += cmd.value
                break
            case "down":
                depth += cmd.value
                break
            case "up":
                depth -= cmd.value
                break
            default:
                throw new Error("Unexpected direction name.")
        }
    }
    return horizontal * depth
}

function part2(cmds: Cmd[]): number {
    let horizontal = 0
    let depth = 0
    let aim = 0
    for (const cmd of cmds) {
        switch (cmd.direction) {
            case "forward":
                horizontal += cmd.value
                depth += aim * cmd.value
                break
            case "down":
                aim += cmd.value
                break
            case "up":
                aim -= cmd.value
                break
            default:
                throw new Error("Unexpected direction name.")
        }
    }
    return horizontal * depth
}

async function main(): Promise<void> {
    const cmds = await readInput()
    console.log("[1]", part1(cmds))
    console.log("[2]", part2(cmds))
}

main();
