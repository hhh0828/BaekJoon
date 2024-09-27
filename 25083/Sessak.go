package main

import (
	"bufio"
	"os"
)

var Writer *bufio.Writer

func init() {
	Writer = bufio.NewWriter(os.Stdout)
}
func main() {
	defer Writer.Flush()
	Writer.WriteString("         ,r'\"7\n")
	Writer.WriteString("r`-_   ,'  ,/\n")
	Writer.WriteString(" \\. \". L_r'\n")
	Writer.WriteString("   `~\\/\n")
	Writer.WriteString("      |\n")
	Writer.WriteString("      |")
}
