package utils

import(
	"time"
	"math/rand"
)

func Random(min, max float64) float64 {
    rand.Seed(time.Now().UnixNano())
    return rand.Float64()*(max-min) + min
}
