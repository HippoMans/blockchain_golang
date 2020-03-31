package blockFunction

import "fmt"

func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlock(data)
	fmt.Println("Success")
}
