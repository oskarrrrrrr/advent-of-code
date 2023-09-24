package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	op_nop = iota
	op_acc
	op_jmp
)

type Instruction struct {
	op  int
	arg int
}

func NewInstruction(s string) Instruction {
	fields := strings.Fields(s)
	if len(fields) != 2 {
		log.Fatal("Invalid instruction format: ", s)
	}

	var op int
	switch op_str := fields[0]; op_str {
	case "nop":
		op = op_nop
	case "acc":
		op = op_acc
	case "jmp":
		op = op_jmp
	default:
		log.Fatal("Unrecognized operation: ", op_str)
	}

	arg_str := fields[1]
	arg, err := strconv.Atoi(arg_str)
	if err != nil {
		log.Fatal(err)
	}

	return Instruction{
		op:  op,
		arg: arg,
	}
}

type Program struct {
	instructions []Instruction
	pc           int // program counter
	acc          int // accumulator
}

func NewProgram(Instructions []Instruction) Program {
	return Program{instructions: Instructions}
}

func (p *Program) RunUntilLoop() (bool, bool) {
	visited := make([]bool, len(p.instructions))
	p.pc = 0
	p.acc = 0
	for {
        if p.pc == len(p.instructions) {
            return false, true
        }
        if p.pc > len(p.instructions) {
            return false, false
        }
		if visited[p.pc] {
			return true, false
		}
		visited[p.pc] = true
		i := p.instructions[p.pc]
		switch i.op {
		case op_nop:
			p.pc++
		case op_acc:
			p.acc += i.arg
			p.pc++
		case op_jmp:
			p.pc += i.arg
		default:
			log.Fatal("Unrecognized operation: ", i.op)
		}
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new_inst := NewInstruction(scanner.Text())
		instructions = append(instructions, new_inst)
	}

	program := NewProgram(instructions)
	program.RunUntilLoop()
	fmt.Println("[1]", program.acc)

    for i := 0; i < len(instructions); i++ {
        ins := &instructions[i]
        original_op := ins.op
        if ins.op == op_nop {
            ins.op = op_jmp
        } else if ins.op == op_jmp {
            ins.op = op_nop
        }
        _, success := program.RunUntilLoop()
        if success {
            fmt.Println("[2]", program.acc)
        }
        ins.op = original_op
    }
}
