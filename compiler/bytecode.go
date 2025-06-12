package compiler

import "fmt"

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota // 加载常量
	OpAdd                    // 加
	OpSub                    // 减
)

// 用于调试
func (ins Instructions) String() string {
	var out string
	i := 0
	for i < len(ins) {
		op := Opcode(ins[i])
		switch op {
		case OpConstant:
			// OpConstant 指令后面跟着一个2字节的操作数 (常量池索引)
			operand := int(ins[i+1])<<8 | int(ins[i+2])
			out += fmt.Sprintf("OpConstant %d\n", operand)
			i += 3
		case OpAdd:
			out += "OpAdd\n"
			i++
		case OpSub:
			out += "OpSub\n"
			i++
		}
	}
	return out
}

type Bytecode struct {
	Instructions Instructions
	Constants    []interface{}
}
