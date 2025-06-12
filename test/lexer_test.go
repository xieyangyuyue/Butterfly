// 定义测试包（使用_test后缀表示测试包）
package test_test

// 导入所需包
import (
	"Butterfly/lexer"                    // 自定义词法分析器包
	"github.com/stretchr/testify/assert" // 断言库用于测试验证
	"os"                                 // 操作系统功能包
	"path/filepath"                      // 文件路径处理包
	"strings"                            // 字符串处理包
	"testing"                            // Go测试框架包
)

// TestLexer 是词法分析器的主测试函数
// 使用标准测试签名(t *testing.T)
func TestLexer(t *testing.T) {
	// 定义测试用例结构体切片
	// 每个测试用例包含：
	//   name: 测试用例名称（用于标识）
	//   testFile: 输入测试文件名
	//   expectFile: 预期输出文件名
	testCases := []struct {
		name       string
		testFile   string
		expectFile string
	}{
		// 三个测试用例定义
		{"Testcase1", "testfile1.txt", "output1.txt"}, // 测试用例1
		{"Testcase2", "testfile2.txt", "output2.txt"}, // 测试用例2
		{"Testcase3", "testfile3.txt", "output3.txt"}, // 测试用例3
	}

	// 遍历所有测试用例
	for _, tc := range testCases {
		// 运行子测试（可单独执行）
		// 第一个参数：子测试名称
		// 第二个参数：测试函数
		t.Run(tc.name, func(t *testing.T) {
			// ====== 1. 准备输入数据 ======
			// 构建完整文件路径：testdata目录 + 测试文件名
			inputPath := filepath.Join("testdata", tc.testFile)
			// 读取文件内容
			input := readFile(t, inputPath)

			// ====== 2. 执行词法分析 ======
			// 创建词法分析器实例
			l := lexer.New(input)
			// 存储词法单元字符串表示的切片
			var tokens []string

			// 循环获取所有词法单元
			for {
				// 获取下一个词法单元
				token := l.NextToken()
				// 将词法单元转为字符串并添加到切片
				tokens = append(tokens, token.String())
				// 遇到EOF结束符时终止循环
				if token.Type == lexer.EOF {
					break
				}
			}

			// ====== 3. 准备预期输出 ======
			// 构建完整文件路径：testdata目录 + 预期输出文件名
			expectPath := filepath.Join("testdata", tc.expectFile)
			// 读取预期输出文件（按行分割）
			expected := readFileLines(t, expectPath)

			// ====== 4. 验证结果 ======
			// 规范化处理预期输出和实际输出
			normExpected := normalize(expected)
			normTokens := normalize(tokens)

			// 使用断言比较两者是否相同
			// 参数1: 测试对象t
			// 参数2: 期望值
			// 参数3: 实际值
			// 参数4: 失败时的自定义消息
			assert.Equal(t, normExpected, normTokens,
				"测试用例失败: "+tc.name)
		})
	}
}

// readFile 读取文件内容并返回字符串
// 参数:
//
//	t: 测试对象（用于错误报告）
//	path: 文件路径
//
// 返回值:
//
//	文件内容的字符串表示
func readFile(t *testing.T, path string) string {
	// 读取文件内容到字节数组
	data, err := os.ReadFile(path)
	// 错误处理：测试失败并终止
	if err != nil {
		// Fatalf会终止当前测试用例的执行
		t.Fatalf("文件读取错误: %v", err)
	}
	// 将字节数组转为字符串返回
	return string(data)
}

// readFileLines 按行读取文件内容
// 参数:
//
//	t: 测试对象
//	path: 文件路径
//
// 返回值:
//
//	按行分割的字符串切片
func readFileLines(t *testing.T, path string) []string {
	// 调用readFile获取文件内容字符串
	content := readFile(t, path)
	// 去除首尾空白字符（包括换行符）
	trimmedContent := strings.TrimSpace(content)
	// 按换行符分割为行切片
	return strings.Split(trimmedContent, "\n")
}

// normalize 规范化文本行
// 功能:
//  1. 移除空行
//  2. 压缩连续空白为单个空格
//
// 参数:
//
//	lines: 输入行切片
//
// 返回值:
//
//	规范化后的行切片
func normalize(lines []string) []string {
	// 创建容量足够的空切片（容量=输入行数）
	normalized := make([]string, 0, len(lines))

	// 遍历每行输入
	for _, line := range lines {
		// 去除首尾空白
		trimmed := strings.TrimSpace(line)
		// 跳过空行
		if trimmed != "" {
			// 分割连续空白（Fields函数）
			// 示例: "a   b" -> ["a", "b"]
			fields := strings.Fields(trimmed)
			// 用单个空格连接字段
			// 示例: ["a", "b"] -> "a b"
			cleanLine := strings.Join(fields, " ")
			// 添加到结果切片
			normalized = append(normalized, cleanLine)
		}
	}
	// 返回规范化后的行切片
	return normalized
}
