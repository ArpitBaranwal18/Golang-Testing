package functions

// Compare 2 strings,
// Return true if both are equal
// Return false otherwise
func CompareStr(str1, str2 string) bool {
	// implement this
	if str1 == str2{
		return true
	}
	return false
}

// Check if a string is a palindrome
// e.g: level 
func CheckPalindrome(str string) bool {
	// implement this
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

// If we have 2 packages in order for 1 package
// to see functions from the other package you
// have to "export" these functions.
// Just start the functions with an Uppercase.
// sum of first "n" integers
func Sumn(n int) int {
	s := 0
	for i:=1; i<=n; i++ {
		s += i
	}
	return s
}

// factorial (seen only locally - within the package)
func Factorial(n int) int {
	f := 1

	for i:=1; i<=n; i++ {
		f *= i
	}

	return f
}

