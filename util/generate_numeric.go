package util

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class utils generator numeric
*/

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type GeneratorNumeric struct {
}

var generatorNun *GeneratorNumeric

func NewGeneratorNumeric() *GeneratorNumeric {
	lock := &sync.Mutex{}
	if generatorNun == nil {
		lock.Lock()
		defer lock.Unlock()
		generatorNun = &GeneratorNumeric{}
	}
	return generatorNun
}

func (g GeneratorNumeric) Generate(amountDigits uint64) int {
	rand.Seed(time.Now().UnixNano())
	num1 := rand.Intn(time.Now().Hour())
	num2 := rand.Intn(time.Now().Minute())
	num3 := rand.Intn(time.Now().Second())
	return rand.Intn(int(amountDigits)) + int(amountDigits) + num1 + num2 + num3
}

func (g GeneratorNumeric) ToString(amountDigits int) string {
	return strconv.FormatInt(int64(amountDigits), 10)
}

func (g GeneratorNumeric) GenerateToString(amountDigits uint64) string {
	return g.ToString(g.Generate(amountDigits))
}
