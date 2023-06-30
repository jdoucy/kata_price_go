package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputePrice(t *testing.T) {
	t.Run(`3 articles à 1,21 € et taxe 0 % → “3.63 €”`, func(t *testing.T) {
		assert.Equal(t, "3.63 €", ComputePrice(3, 121, 0.0))
	})
	t.Run(`3 articles à 1,21 € et taxe 5 % → “3.81 €”`, func(t *testing.T) {
		assert.Equal(t, "3.81 €", ComputePrice(3, 121, 0.05))
	})
	t.Run(`3 articles à 1,21 € et taxe 20 % → “4.36 €”`, func(t *testing.T) {
		assert.Equal(t, "4.36 €", ComputePrice(3, 121, 0.2))
	})
	t.Run(`1000 € → 3% remise 5 x 345,00 € + taxe 10% → “1840.58 €”`, func(t *testing.T) {
		assert.Equal(t, "6787.28 €", ComputePrice(5, 129900, 0.1))
	})
	t.Run(`5000 € → 5% remise 5 x 1299,00 € + taxe 10% → “6787.28 €”`, func(t *testing.T) {
		assert.Equal(t, "6787.28 €", ComputePrice(5, 129900, 0.1))
	})
}
