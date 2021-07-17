package vm

type VM struct {
	Code  *Bytecode
	ip    int
	stack *Stack
	sp    int
}

func CreateVM() *VM {
	return nil
}

const (

)

func (vm *VM) Interpret(bytecode *Bytecode) int {

}

func (vm *VM) advance() int {
	vm.ip += 1
	return vm.Code[vm.ip]
}

func (vm *VM) Run() {

	instruction := vm.advance()

	switch instruction {
	case OP_RETURN:
		break
	case OP_CONSTANT:
		break
	case OP_ADD:
		break
	case OP_SUBTRACT:
		break
	case OP_MULTIPLY:
		break
	case OP_DIVIDE:
		break

	}
}
