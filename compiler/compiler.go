package compiler

import (
	"Butterfly/ast"
)

type Compiler struct {
	instructions Instructions
	constants    []interface{}
}

func New() *Compiler {
	return &Compiler{
		instructions: Instructions{},
		constants:    []interface{}{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}

	case *ast.InfixExpression:
		err := c.Compile(node.Left)
		if err != nil {
			return err
		}
		err = c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "+":
			c.emit(OpAdd)
		case "-":
			c.emit(OpSub)
		}

	case *ast.IntegerLiteral:
		c.constants = append(c.constants, node.Value)
		// 将常量的索引作为 OpConstant 的操作数
		c.emit(OpConstant, len(c.constants)-1)
	}

	return nil
}

// emit 发出指令和操作数
func (c *Compiler) emit(op Opcode, operands ...int) {
	ins := []byte{byte(op)}
	for _, o := range operands {
		// 这里我们假设操作数是2字节的
		ins = append(ins, byte(o>>8)) // 高位
		ins = append(ins, byte(o))    // 低位
	}
	c.instructions = append(c.instructions, ins...)
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}
