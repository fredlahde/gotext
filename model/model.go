package model

import (
	"github.com/fredlahde/gotext/types"
	"gonum.org/v1/gonum/mat"
)

type Loss struct{}

type Model struct {
	wi                *mat.Matrix
	wo                *mat.Matrix
	loss              *Loss
	normalizeGradient bool
}

type State struct {
	lossValue types.Real
	nexamples int64
	hidden    mat.Vector
	output    mat.Vector
	gradient  mat.Vector
}
