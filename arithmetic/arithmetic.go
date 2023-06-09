package arithmetic

// Checks if a number is prime or not
func IsPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func IsEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
