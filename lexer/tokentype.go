package lexer // Package lexer 词法分析器包
// TokenType 定义词法单元类型的枚举
type TokenType int

// 词法单元类型常量定义
const (
	// PUBLIC 关键字类型 (0-19)
	PUBLIC   TokenType = iota // 0: public关键字
	CLASS                     // 1: class关键字
	STATIC                    // 2: static关键字
	VOID                      // 3: void关键字
	MAIN                      // 4: main关键字
	CHAR                      // 5: char关键字
	INT                       // 6: int关键字
	PRINTF                    // 7: printf关键字
	SCANF                     // 8: scanf关键字
	SWITCH                    // 9: switch关键字
	CASE                      // 10: case关键字
	DEFAULT                   // 11: default关键字
	FOR                       // 12: for关键字
	IF                        // 13: if关键字
	ELSE                      // 14: else关键字
	WHILE                     // 15: while关键字
	DO                        // 16: do关键字
	RETURN                    // 17: return关键字
	BREAK                     // 18: break关键字
	CONTINUE                  // 19: continue关键字

	// IDENTIFIER 标识符和字面量类型 (20-23)
	IDENTIFIER // 20: 标识符（变量名/函数名等）
	STRING     // 21: 字符串常量
	CharConst  // 22: 字符常量
	NUMBER     // 23: 整型常量

	// ASSIGN 运算符类型 (24-35)
	ASSIGN     // 24: 赋值运算符（=）
	PLUS       // 25: 加号（+）
	MINUS      // 26: 减号（-）
	MULTIPLY   // 27: 乘号（*）
	DIVIDE     // 28: 除号（/）
	LESS       // 29: 小于号（<）
	LessEqual  // 30: 小于等于（<=）
	GREAT      // 31: 大于号（>）
	GreatEqual // 32: 大于等于（>=）
	NOTEQ      // 33: 不等于（!=）
	NOT        // 34: 非运算符（!）
	EQUAL      // 35: 等于（==）

	// COMMA 分隔符类型 (36-44)
	COMMA      // 36: 逗号（,）
	SEMICOLON  // 37: 分号（;）
	COLON      // 38: 冒号（:）
	LeftParen  // 39: 左圆括号（(）
	RightParen // 40: 右圆括号（)）
	LeftBrace  // 41: 左花括号（{）
	RightBrace // 42: 右花括号（}）
	LeftBrack  // 43: 左方括号（[）
	RightBrack // 44: 右方括号（]）

	// EOF 特殊类型 (45)
	EOF // 45: 文件结束标志
)

// 词法单元类型到输出代码的映射表
var tokenTypeCodes = map[TokenType]string{
	PUBLIC:     "PUBLIC",     // 0: 输出代码保持原样
	CLASS:      "CLASS",      // 1
	STATIC:     "STATIC",     // 2
	VOID:       "VOIDTK",     // 3: 特殊处理避免与类型名冲突
	MAIN:       "MAINTK",     // 4
	CHAR:       "CHARTK",     // 5
	INT:        "INTTK",      // 6
	PRINTF:     "PRINTFTK",   // 7
	SCANF:      "SCANFTK",    // 8
	SWITCH:     "SWITCHTK",   // 9
	CASE:       "CASETK",     // 10
	DEFAULT:    "DEFAULTTK",  // 11
	FOR:        "FORTK",      // 12
	IF:         "IFTK",       // 13
	ELSE:       "ELSETK",     // 14
	WHILE:      "WHILETK",    // 15
	DO:         "DOTK",       // 16
	RETURN:     "RETURNTK",   // 17
	BREAK:      "BREAKTK",    // 18
	CONTINUE:   "CONTINUETK", // 19
	IDENTIFIER: "IDENFR",     // 20: 标识符特殊代码
	STRING:     "STRCON",     // 21: 字符串常量
	CharConst:  "CHARCON",    // 22: 字符常量
	NUMBER:     "INTCON",     // 23: 整型常量
	ASSIGN:     "ASSIGN",     // 24
	PLUS:       "PLUS",       // 25
	MINUS:      "MINU",       // 26: 特殊缩写（原MINUS）
	MULTIPLY:   "MULT",       // 27
	DIVIDE:     "DIV",        // 28
	LESS:       "LSS",        // 29: 特殊缩写（原LESS）
	LessEqual:  "LEQ",        // 30
	GREAT:      "GRE",        // 31: 特殊缩写（原GREAT）
	GreatEqual: "GEQ",        // 32
	NOTEQ:      "NEQ",        // 33
	NOT:        "NOT",        // 34
	EQUAL:      "EQL",        // 35: 特殊缩写（原EQUAL）
	COMMA:      "COMMA",      // 36
	SEMICOLON:  "SEMICN",     // 37: 特殊缩写（原SEMICOLON）
	COLON:      "COLON",      // 38
	LeftParen:  "LPARENT",    // 39: 左圆括号特殊代码
	RightParen: "RPARENT",    // 40: 右圆括号特殊代码
	LeftBrace:  "LBRACE",     // 41
	RightBrace: "RBRACE",     // 42
	LeftBrack:  "LBRACK",     // 43
	RightBrack: "RBRACK",     // 44
	EOF:        "EOF",        // 45: 文件结束标志
}

// Code 返回词法单元类型的标准输出代码
func (tt TokenType) Code() string {
	// 通过映射表返回对应字符串代码
	// 如果未定义则返回空字符串（但所有类型均已定义）
	return tokenTypeCodes[tt]
}
