package types

type Real float64

type RealIntPair struct {
	R Real
	I int32
}

type Predictions []RealIntPair
