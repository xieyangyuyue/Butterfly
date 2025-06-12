// Package lexer 定义了词法分析器相关的代码。
// 它的主要作用是读取源代码文本，并将其分解成一系列的词法单元（Tokens）。
package lexer

import (
	"fmt"     // 导入 fmt 包，用于格式化字符串，主要在错误处理中使用。
	"unicode" // 导入 unicode 包，提供了一系列函数来检查字符的属性（如是否是字母、数字、空白等）。
)

// Lexer 结构体代表一个词法分析器。
// 它持有完整的输入字符串，并跟踪当前解析的位置。
type Lexer struct {
	input       string // 要进行词法分析的完整源代码字符串。
	pos         int    // 当前字符在 input 字符串中的索引位置。
	line        int    // 当前所在的代码行号，用于错误定位。
	column      int    // 当前所在的代码列号，用于错误定位。
	currentChar rune   // 当前正在检查的字符。使用 rune 类型以支持 Unicode 字符。
}

// reverseKeywords 是一个从字符串到 TokenType 的反向映射表。
// 它用于在读取一个标识符后，快速判断该标识符是否是一个语言的关键字。
// 例如，当读取到 "if" 字符串时，可以通过这个 map 查到它对应的 TokenType 是 IF。
var reverseKeywords = map[string]TokenType{
	"public":   PUBLIC,
	"class":    CLASS,
	"static":   STATIC,
	"void":     VOID,
	"main":     MAIN,
	"char":     CHAR,
	"int":      INT,
	"printf":   PRINTF,
	"scanf":    SCANF,
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
	"for":      FOR,
	"if":       IF,
	"else":     ELSE,
	"while":    WHILE,
	"do":       DO,
	"return":   RETURN,
	"break":    BREAK,
	"continue": CONTINUE,
}

// New 是 Lexer 的构造函数，用于创建一个新的词法分析器实例。
// 它接收源代码字符串作为输入，并初始化 Lexer 的状态。
func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1, // 初始行号为 1。
		column: 1, // 初始列号为 1。
	}
	// 如果输入不为空，则读取第一个字符以初始化 currentChar。
	if len(input) > 0 {
		l.currentChar = rune(input[0])
	}
	return l
}

// advance 方法将词法分析器的位置向前移动一个字符。
// 这是词法分析器在输入流中移动的基本操作。
func (l *Lexer) advance() {
	// 如果当前字符是换行符，则行号加 1，列号重置为 0（因为 advance 后会立即加 1）。
	if l.currentChar == '\n' {
		l.line++
		l.column = 0
	}
	l.pos++ // 移动位置索引。
	// 检查是否已经到达输入字符串的末尾。
	if l.pos >= len(l.input) {
		l.currentChar = 0 // 0 通常表示 EOF (End of File)。
	} else {
		// 读取下一个字符。
		l.currentChar = rune(l.input[l.pos])
	}
	l.column++ // 不论如何，列号都加 1。
}

// skipWhitespace 方法会跳过所有连续的空白字符（如空格、制表符、换行符等）。
func (l *Lexer) skipWhitespace() {
	// 循环直到当前字符不再是空白字符。
	for unicode.IsSpace(l.currentChar) {
		l.advance()
	}
}

// readIdentifier 方法读取一个完整的标识符或关键字。
// 标识符通常由字母和数字组成，且以字母开头。
func (l *Lexer) readIdentifier() Token {
	startPos := l.pos    // 记录标识符的起始位置。
	startLine := l.line  // 记录起始行号。
	startCol := l.column // 记录起始列号。

	// 持续前进，直到当前字符不再是字母或数字。
	for unicode.IsLetter(l.currentChar) || unicode.IsDigit(l.currentChar) {
		l.advance()
	}

	// 从输入中提取出完整的标识符字符串。
	value := l.input[startPos:l.pos]
	// 默认类型为普通标识符。
	tokenType := IDENTIFIER
	// 检查这个标识符是否是预定义的关键字。
	if code, exists := reverseKeywords[value]; exists {
		tokenType = code // 如果是关键字，则更新 TokenType。
	}
	// 返回构造好的 Token。
	return Token{tokenType, value, startLine, startCol}
}

// readNumber 方法读取一个完整的整型数字字面量。
func (l *Lexer) readNumber() Token {
	startPos := l.pos    // 记录数字的起始位置。
	startLine := l.line  // 记录起始行号。
	startCol := l.column // 记录起始列号。

	// 持续前进，直到当前字符不再是数字。
	for unicode.IsDigit(l.currentChar) {
		l.advance()
	}

	// 提取数字字符串。
	value := l.input[startPos:l.pos]
	// 返回 NUMBER 类型的 Token。
	return Token{NUMBER, value, startLine, startCol}
}

// readChar 方法解析一个字符字面量（例如 'a', '\n'）。
func (l *Lexer) readChar() Token {
	startLine := l.line  // 记录起始行号。
	startCol := l.column // 记录起始列号。
	l.advance()          // 跳过开始的单引号。

	var value string
	// 检查是否是转义字符。
	if l.currentChar == '\\' {
		l.advance() // 跳过反斜杠。
		switch l.currentChar {
		case 'n':
			value = "\n"
		case 't':
			value = "\t"
		case 'r':
			value = "\r"
		case '\'':
			value = "'"
		case '\\':
			value = "\\"
		default:
			// 如果是未知的转义序列，则引发恐慌（panic）。
			panic(fmt.Sprintf("无效转义字符: \\%c 在行 %d:%d", l.currentChar, l.line, l.column))
		}
		l.advance() // 移过转义字符本身。
	} else {
		// 如果不是转义字符，则直接读取该字符。
		value = string(l.currentChar)
		l.advance()
	}

	// 检查字符字面量是否以单引号正确闭合。
	if l.currentChar != '\'' {
		panic(fmt.Sprintf("字符未闭合，起始于行 %d:%d", startLine, startCol))
	}
	l.advance() // 跳过闭合的单引号。
	// 返回 CharConst 类型的 Token。
	return Token{CharConst, value, startLine, startCol}
}

// readString 方法解析一个字符串字面量（例如 "hello world"）。
func (l *Lexer) readString() Token {
	startLine := l.line  // 记录起始行号。
	startCol := l.column // 记录起始列号。
	l.advance()          // 跳过开始的双引号。

	var value string
	// 循环读取字符，直到遇到闭合的双引号或文件末尾。
	for l.currentChar != '"' && l.currentChar != 0 {
		// 处理转义字符。
		if l.currentChar == '\\' {
			l.advance() // 跳过反斜杠。
			switch l.currentChar {
			case 'n':
				value += "\n"
			case 't':
				value += "\t"
			case 'r':
				value += "\r"
			case '"':
				value += "\""
			case '\\':
				value += "\\"
			default:
				// 抛出无效转义字符的错误。
				panic(fmt.Sprintf("Invalid escape: \\%c at %d:%d", l.currentChar, l.line, l.column))
			}
			l.advance() // 移过转义字符本身。
		} else {
			// 将普通字符追加到结果字符串中。
			value += string(l.currentChar)
			l.advance()
		}
	}

	// 检查字符串是否以双引号正确闭合。
	if l.currentChar != '"' {
		panic(fmt.Sprintf("Unclosed string at %d:%d", startLine, startCol))
	}
	l.advance() // 跳过闭合的双引号。
	// 返回 STRING 类型的 Token。
	return Token{STRING, value, startLine, startCol}
}

// NextToken 是词法分析器的核心方法。
// 每次调用，它都会从源代码中解析并返回下一个 Token。
func (l *Lexer) NextToken() Token {
	// 主循环，只要没到文件末尾就一直运行。
	for l.currentChar != 0 {
		// 如果是空白字符，则跳过。
		if unicode.IsSpace(l.currentChar) {
			l.skipWhitespace()
			continue // 继续下一次循环，获取下一个非空白字符。
		}
		// 如果是单引号，开始解析字符字面量。
		if l.currentChar == '\'' {
			return l.readChar()
		}

		// 如果是双引号，开始解析字符串字面量。
		if l.currentChar == '"' {
			return l.readString()
		}

		// 如果是字母，开始解析标识符或关键字。
		if unicode.IsLetter(l.currentChar) {
			return l.readIdentifier()
		}

		// 如果是数字，开始解析数字字面量。
		if unicode.IsDigit(l.currentChar) {
			return l.readNumber()
		}

		// ---- 处理各种单字符和多字符的运算符及分隔符 ----
		currentChar := l.currentChar // 保存当前字符，因为 l.advance() 会改变它。
		currentLine := l.line        // 保存当前行号。
		currentCol := l.column       // 保存当前列号。
		l.advance()                  // 向前移动，以便检查可能的双字符运算符。

		switch currentChar {
		case '=':
			if l.currentChar == '=' { // 检查是否是 "=="
				l.advance()
				return Token{EQUAL, "==", currentLine, currentCol}
			}
			return Token{ASSIGN, "=", currentLine, currentCol} // 否则是 "="
		case '!':
			if l.currentChar == '=' { // 检查是否是 "!="
				l.advance()
				return Token{NOTEQ, "!=", currentLine, currentCol}
			}
			return Token{NOT, "!", currentLine, currentCol} // 否则是 "!"
		case '<':
			if l.currentChar == '=' { // 检查是否是 "<="
				l.advance()
				return Token{LessEqual, "<=", currentLine, currentCol}
			}
			return Token{LESS, "<", currentLine, currentCol} // 否则是 "<"
		case '>':
			if l.currentChar == '=' { // 检查是否是 ">="
				l.advance()
				return Token{GreatEqual, ">=", currentLine, currentCol}
			}
			return Token{GREAT, ">", currentLine, currentCol} // 否则是 ">"
		case '+':
			return Token{PLUS, "+", currentLine, currentCol}
		case '-':
			return Token{MINUS, "-", currentLine, currentCol}
		case '*':
			return Token{MULTIPLY, "*", currentLine, currentCol}
		case '/':
			return Token{DIVIDE, "/", currentLine, currentCol}
		case ';':
			return Token{SEMICOLON, ";", currentLine, currentCol}
		case ',':
			return Token{COMMA, ",", currentLine, currentCol}
		case '(':
			return Token{LeftParen, "(", currentLine, currentCol}
		case ')':
			return Token{RightParen, ")", currentLine, currentCol}
		case '{':
			return Token{LeftBrace, "{", currentLine, currentCol}
		case '}':
			return Token{RightBrace, "}", currentLine, currentCol}
		case '[':
			return Token{LeftBrack, "[", currentLine, currentCol}
		case ']':
			return Token{RightBrack, "]", currentLine, currentCol}
		case ':':
			return Token{COLON, ":", currentLine, currentCol}
		default:
			// 如果遇到无法识别的字符，则引发恐慌。
			panic(fmt.Sprintf("Unexpected character: %c at %d:%d", currentChar, currentLine, currentCol))
		}
	}

	// 如果循环结束（currentChar 为 0），说明到达了文件末尾，返回 EOF Token。
	return Token{EOF, "", l.line, l.column}
}
