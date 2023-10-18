package main

import (
	"fmt"
	"math/big"
)

type MyBigInt struct {
	value *big.Int
}

func MyNewBigInt() *MyBigInt {
	return &MyBigInt{value: new(big.Int)}
}

func (n *MyBigInt) SetHex(value string) error {
	_, success := n.value.SetString(value, 16)
	if !success {
		return fmt.Errorf("Недійсне шістнадцяткове число")
	}
	return nil
}

func (n *MyBigInt) GetHex() string {
	return fmt.Sprintf("%x", n.value)
}

func INV(a *MyBigInt) *MyBigInt {
	result := MyNewBigInt()
	result.value.Not(a.value)
	return result
}

func XOR(a, b *MyBigInt) *MyBigInt {
	result := MyNewBigInt()
	result.value.Xor(a.value, b.value)
	return result
}

func OR(a, b *MyBigInt) *MyBigInt {
	result := MyNewBigInt()
	result.value.Or(a.value, b.value)
	return result
}

func AND(a, b *MyBigInt) *MyBigInt {
	result := MyNewBigInt()
	result.value.And(a.value, b.value)
	return result
}

func shiftR(a *MyBigInt, n uint) *MyBigInt {
	result := MyNewBigInt()
	result.value.Rsh(a.value, n)
	return result
}

func shiftL(a *MyBigInt, n uint) *MyBigInt {
	result := MyNewBigInt()
	result.value.Lsh(a.value, n)
	return result
}

func main() {
	numberA := MyNewBigInt()
	numberB := MyNewBigInt()

	if err := numberA.SetHex("e035c6cfa42609b998b883bc1699df885cef74e2b2cc372eb8fa7e7"); err != nil {
		fmt.Println("Помилка встановлення числа A:", err)
	}
	if err := numberB.SetHex("5072f028943e0fd5fab3273782de14b1011741bd0c5cd6ba6474330"); err != nil {
		fmt.Println("Помилка встановлення числа B:", err)
	}

	numberC := INV(numberA)
	fmt.Println("INV (побітова інверсія):", numberC.GetHex())

	numberC = XOR(numberA, numberB)
	fmt.Println("XOR (побітове виключне або):", numberC.GetHex())

	numberC = OR(numberA, numberB)
	fmt.Println("OR (побітове або):", numberC.GetHex())

	numberC = AND(numberA, numberB)
	fmt.Println("AND (побітове і):", numberC.GetHex())

	numberC = shiftR(numberA, 5)
	fmt.Println("shiftR (зсув праворуч на n бітів):", numberC.GetHex())

	numberC = shiftL(numberA, 5)
	fmt.Println("shiftL (зсув ліворуч на n бітів):", numberC.GetHex())
}
