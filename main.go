// 主程序入口包
package main

// 导入依赖包
import (
	"Butterfly/compiler" // 自定义编译器实现
	"Butterfly/lexer"    // 自定义词法分析器
	"Butterfly/parser"   // 自定义语法分析器
	"Butterfly/vm"       // 自定义字节码虚拟机
	"fmt"                // 提供格式化输入输出功能
	"os"                 // 提供操作系统功能接口和文件操作
)

// 程序主函数
func main() {
	// 验证命令行参数数量
	if len(os.Args) != 2 {
		// 参数不足时显示使用说明
		fmt.Println("用法: go run . <文件名.calc>")
		return // 终止程序
	}

	// 获取输入文件路径（第一个命令行参数）
	filepath := os.Args[1]

	// ========== 文件读取阶段 ==========
	// 使用 os.ReadFile
	// 读取源文件内容到字节数组
	data, err := os.ReadFile(filepath)
	// 处理文件读取错误
	if err != nil {
		// 显示错误详情
		fmt.Printf("文件读取失败: %s\n", err)
		return // 终止程序
	}

	// ========== 词法分析阶段 ==========
	// 创建词法分析器实例（将文件内容转为字符串）
	l := lexer.New(string(data))

	// ========== 语法分析阶段 ==========
	// 创建语法分析器实例（基于词法分析器）
	p := parser.New(l)
	// 解析程序生成抽象语法树(AST)
	program := p.ParseProgram()

	// 检查语法错误集合
	if len(p.Errors()) != 0 {
		// 输出错误标题
		fmt.Println("语法分析错误:")
		// 遍历输出所有错误信息（缩进格式）
		for _, msg := range p.Errors() {
			fmt.Println("\t" + msg)
		}
		return // 终止程序
	}

	// ========== 编译阶段 ==========
	// 初始化编译器实例
	c := compiler.New()
	// 将抽象语法树编译为字节码
	err = c.Compile(program)
	// 处理编译错误
	if err != nil {
		// 输出编译错误详情
		fmt.Printf("编译失败: %s\n", err)
		return // 终止程序
	}

	// (调试选项) 打印字节码指令
	fmt.Println("字节码指令集:")
	// 输出字节码指令序列（字符串表示）
	fmt.Println(c.Bytecode().Instructions)

	// ========== 虚拟机执行阶段 ==========
	// 初始化虚拟机（传入编译后的字节码）
	machine := vm.New(c.Bytecode())
	// 执行字节码指令
	err = machine.Run()
	// 处理虚拟机执行错误
	if err != nil {
		// 输出执行错误详情
		fmt.Printf("虚拟机执行失败: %s\n", err)
		return // 终止程序
	}

	// ========== 结果输出阶段 ==========
	// 获取栈顶元素作为最终计算结果
	// (函数变更说明：从LastPoppedStackElem()改为StackTop())
	result := machine.StackTop()
	// 格式化输出计算结果
	fmt.Printf("计算结果: %v\n", result)
}
