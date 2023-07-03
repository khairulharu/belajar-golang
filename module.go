package belajargolang

func Mltplction(aType int, bType int) int {
	return aType * bType
}

func Equal(zType int, xType int) bool {
	if zType == xType {
		return true
	} else {
		return false
	}
}

func Substrac(aType int, bType int) int {
	return aType - bType
}

func Addction(aType int, bType int) int {
	return aType + bType
}

func Sameas(zType int) string {
	if zType%2 == 0 {
		return "genap"
	} else {
		return "ganjil"
	}
}
