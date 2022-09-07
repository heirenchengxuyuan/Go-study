package Hello

import "fmt"

func main() {
	//大括号应该紧接在前而不应该重新换行
	fmt.Println("hello\tjakc") // \t是一个制表位 实现对齐功能
	//go语言推荐用行注释 不用块注释
	fmt.Println("hello\rjakc")
	// \r 是回车 从当前行最前面输出，覆盖掉以前内容 \n是换行
	//如果要输入路径则要用“\\”作为转义符号输出斜杠
	//可在cmd中使用gofmt命令进行格式化  gofmt -w xxx.go  将新格式重写进文件中
	//运算符两边习惯加一个空格
	fmt.Println("hellohellohellohellohellohellohellohellohello",
		"hellohellohellohellohellohellohellohellohello",
		"hellohellohellohellohellohellohellohellohello")
	//一行字符太多则在行尾加上逗号，每一行结束都要用”引着
}
