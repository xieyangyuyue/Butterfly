package vm

import (
	"Butterfly/compiler"
	"fmt"
)

const StackSize = 2048

type VM struct {
	bytecode *compiler.Bytecode
	stack    []interface{}
	sp       int // 栈顶指针 (Stack Pointer)
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		bytecode: bytecode,
		stack:    make([]interface{}, StackSize),
		sp:       0,
	}
}

func (vm *VM) pop() interface{} {
	if vm.sp == 0 {
		return nil
	}
	vm.sp--
	return vm.stack[vm.sp]
}

func (vm *VM) push(o interface{}) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = o
	vm.sp++
	return nil
}

// StackTop 返回栈顶元素但不弹出
func (vm *VM) StackTop() interface{} {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
	ip := 0 // 指令指针
	for ip < len(vm.bytecode.Instructions) {
		op := compiler.Opcode(vm.bytecode.Instructions[ip])
		ip++

		switch op {
		case compiler.OpConstant:
			constIndex := int(vm.bytecode.Instructions[ip])<<8 | int(vm.bytecode.Instructions[ip+1])
			ip += 2
			err := vm.push(vm.bytecode.Constants[constIndex])
			if err != nil {
				return err
			}

		case compiler.OpAdd, compiler.OpSub:
			right := vm.pop().(int64)
			left := vm.pop().(int64)
			var result int64
			if op == compiler.OpAdd {
				result = left + right
			} else {
				result = left - right
			}
			err := vm.push(result)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
