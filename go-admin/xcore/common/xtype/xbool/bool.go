package xbool

func BooleanTo[T interface{}](p bool, trueValue T, falseValue T) T {
	if p {
		return trueValue
	} else {
		return falseValue
	}
}
