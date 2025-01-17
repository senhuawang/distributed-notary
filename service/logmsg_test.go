package service

import (
	"fmt"
	"testing"

	"math/big"

	"github.com/SmartMeshFoundation/distributed-notary/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestSessionLogMsg(t *testing.T) {
	fmt.Println(SessionLogMsg(utils.NewRandomHash(), "123... %s %s", "aaa", "bbbb"))
}

func TestSign(t *testing.T) {
	secp256k1N, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1halfN := new(big.Int).Div(secp256k1N, big.NewInt(2))
	privateKey, _ := crypto.GenerateKey()
	count := 0
	for i := 0; i < 10240; i++ {
		msgToSign := utils.NewRandomHash()
		sig, err := crypto.Sign(msgToSign[:], privateKey)
		fmt.Println(len(sig), err)
		s := sig[32:64]
		t := new(big.Int)
		t.SetBytes(s)
		if t.Cmp(secp256k1halfN) > 0 {
			fmt.Println(t, common.Bytes2Hex(s))
			count++
		}
	}
	fmt.Println(count)
}

func TestTemp(t *testing.T) {
	var tt []byte
	fmt.Println(tt == nil)
}
