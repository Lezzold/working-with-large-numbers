package main

import (
	"encoding/hex"
	"fmt"
)

type BigInteger struct {
	value []byte
}

func (n *BigInteger) SetHex(value string) error {
	bytes, err := hex.DecodeString(value)
	if err != nil {
		return err
	}
	n.value = bytes
	return nil
}

func (n *BigInteger) String() string {
	return hex.EncodeToString(n.value)
}

func main() {
	num := BigInteger{}
	if err := num.SetHex("3A4D"); err != nil {
		fmt.Println("Помилка встановлення числа:", err)
	} else {
		fmt.Println("Число:", num.String())
	}
}
