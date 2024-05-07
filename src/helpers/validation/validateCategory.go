package validation

import (
	"errors"
)

const (
	CLOTHING    string = "Clothing"
	ACCESSORIES string = "Accessories"
	FOOTWEAR    string = "Footwear"
	BEVERAGES   string = "Beverages"
)

func ValidateRace(status string) error {
	switch status {
	case ACCESSORIES,
		CLOTHING,
		FOOTWEAR,
		BEVERAGES:

		return nil
	default:
		return errors.New("invalid Category Type, please check your Category")
	}
}
