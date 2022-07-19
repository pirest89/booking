package validator

import "github.com/go-playground/validator/v10"

func FlightSlotValidator(maxSlot int64) validator.Func {
	return func(fl validator.FieldLevel) bool {
		slot, ok := fl.Field().Interface().(int64)
		if ok {
			return slot <= maxSlot
		}
		return false
	}
}
