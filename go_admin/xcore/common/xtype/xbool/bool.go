package xbool

func BooleanTo[T string | int32](p bool, trueValue T, falseValue T) T {
	if p {
		return trueValue
	} else {
		return falseValue
	}
}
