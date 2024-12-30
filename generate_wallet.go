package main

import (
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
)

func main() {
	// 指定后缀
	targetSuffix := "666"
	fmt.Printf("正在寻找后缀为 %s 的钱包地址...\n", targetSuffix)

	// 不断生成新的钱包，直到找到符合条件的地址
	var account *solana.Wallet
	attempts := 0
	for {
		account = solana.NewWallet()
		pubKeyStr := account.PublicKey().String()
		attempts++
		if strings.HasSuffix(pubKeyStr, targetSuffix) {
			fmt.Printf("尝试次数: %d\n", attempts)
			break
		}
	}

	fmt.Println("\n找到符合条件的账户:")
	fmt.Println("账户私钥:", account.PrivateKey)
	fmt.Println("账户公钥:", account.PublicKey())
	fmt.Println("\n请保存好您的私钥！")
}
