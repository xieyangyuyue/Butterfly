package lexer // 定义词法分析器包

import "fmt" // 导入格式化输出包

// Token 结构体表示词法分析器生成的词法单元
type Token struct {
	Type   TokenType // 词法单元类型（如关键字、运算符等）
	Value  string    // 词法单元的实际字符串值
	Line   int       // 词法单元在源代码中的行号
	Column int       // 词法单元在源代码中的列号
}

// Token的字符串表示方法，用于格式化输出
func (t Token) String() string {
	// 格式化输出：类型代码（左对齐占8字符） + 实际值
	return fmt.Sprintf("%-8s %s", t.Type.Code(), t.Value)
}
