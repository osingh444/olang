package vm

import (
	"fmt"
)

//OPCODES
const (
	OP_RETURN = iota

	OP_CONSTANT = iota

	OP_ADD = iota
	OP_SUBTRACT = iota
	OP_MULTIPLY = iota
	OP_DIVIDE = iota
)

type Bytecode struct {
	Code         []byte
	Constants    []interface{}
	NumConstants int

}

func CreateBytecode() *Bytecode {
	var code []byte
	var consts []interface{}

	bytecode := Bytecode{Code: code, Constants: consts, NumConstants: 0}
	return &bytecode
}

func (bc *Bytecode) WriteByte(b byte) {
	bc.Code = append(bc.Code, b)
}

func (bc *Bytecode) Disassemble() {

	offset := 0

	for offset < len(bc.Code) {
		offset = disassembleInstruction(offset)
	}
}

func (bc *Bytecode) disassembleInstruction(offset int) int {

	opcode := bc.Code[offset]

	switch opcode {
	case OP_RETURN:
		return soloInstruction("OP_RETURN", offset)
	case OP_CONSTANT:
		return constantInstruction("OP_CONSTANT", offset, bc)
	case OP_ADD:
		break
	case OP_SUBTRACT:
		break
	case OP_MULTIPLY:
		break
	case OP_DIVIDE:
		break
	default:
		fmt.Println("unknown opcode")
		return offset + 1
	}
}

func soloInstruction(name string, offset int) int {
	fmt.Println(name)
	return offset + 1
}

func (bc *Bytecode) addConstant(val interface{}) int {
	bc.Constants = append(bc.Constants, val)
	bc.NumConstants += 1
	return bc.NumConstants - 1
}

func constantInstruction(string name, offset int, bc *Bytecode) int {
	fmt.Println("%v, %v", name, bc.Constants[bc.Code[offset]])
	return offset + 2
}
