package main

import bf "bitcoin_golang/blockFunction"

func main() {
        Bchain := bf.NewBlockChain()
        defer Bchain.DB.Close()

        clii := bf.CLI{Bchain}
        clii.Run()
}
