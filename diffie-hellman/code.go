package diffiehellman

import (
	"math/big"
	"math/rand"
)

var (
	randSource = rand.New(rand.NewSource(99))
	scratch    = big.NewInt(0)
	two        = big.NewInt(2)
)

func PrivateKey(p *big.Int) *big.Int {
	scratch.Sub(p, two)
	scratch.Rand(randSource, scratch)
	return scratch.Add(scratch, two)
}

func PublicKey(a, p *big.Int, g int64) *big.Int {
	// g**a mod p
	return scratch.Exp(big.NewInt(g), a, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	a := PrivateKey(p)
	A := PublicKey(a, p, g)
	return a, A
}

func SecretKey(a, B, p *big.Int) *big.Int {
	return scratch.Exp(B, a, p)
}
