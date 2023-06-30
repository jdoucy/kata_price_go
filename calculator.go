package calculator

import (
	"fmt"
	"math"
)

type thresholdDiscount struct {
	thresholdInCents int
	discount         float64
}

func ComputePrice(quantity, priceInCents int, tax float64) string {
	return formatPrice(
		computeTax(
			computeDiscount(
				float64(quantity*priceInCents),
			),
			tax,
		),
	)
}

func computeTax(totalPriceWithDiscountInCents float64, tax float64) float64 {
	return totalPriceWithDiscountInCents * (1 + tax)
}

func computeDiscount(totalPriceInCents float64) float64 {
	for _, thresholdDiscount := range thresholdDiscounts {
		if totalPriceInCents >= float64(thresholdDiscount.thresholdInCents) {
			return totalPriceInCents * (1 - thresholdDiscount.discount)
		}
	}
	return totalPriceInCents
}

func formatPrice(priceInCents float64) string {
	return fmt.Sprintf("%.2f €", math.RoundToEven(priceInCents)/100)
}
