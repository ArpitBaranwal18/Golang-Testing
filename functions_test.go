// In go's convention, any module ending with _test.go
// go will think its a unit test module.
package functions

import (
	"testing"
)

func TestSumn(t *testing.T) {
	// test cases
	// Sumn(1) => 1
	// Sumn(2) => 3
	// Sumn(0) => 0
	// Sumn(5) => 15
	// let me create a map
	vals := map[int]int{
		0: 0,
		1: 1,
		2: 3,
		5: 15,
	}
	
	for k,v := range vals {
		if Sumn(k) != v {
			t.Errorf("Sumn failed for %d", k)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := map[int]int{
		1:1,
		2:2,
		3:6,
		5:120,
		6:720,

	}

	for inp, exp := range tests {
		result := Factorial(inp)
		if result != exp {
			t.Errorf("Factorial(%d) = %d; want %d", inp, result, exp)
		}
		
	}

}

func TestPalindrome(t *testing.T){
	
	tests := map[string] bool{
		"radar": true,
		"arpit": false,
		"level": true,
		"madam": true,
		"abba":  true,
	}

	for input, expected := range tests {
		result := CheckPalindrome(input) 
		if result != expected {
			t.Errorf("CheckPalindrome(%q) = %v; want %v", input, result, expected)
			
		}
	}
}

func TestCompare(t *testing.T){
	tests := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"arpit", "arpit", true},       
		{"akhil", "joges", false},  
		{"", "", true},    
		     
	}

	
	for _, test := range tests {
		
			result := CompareStr(test.str1, test.str2)
			if result != test.expected {
				t.Errorf("CompareStrings(%q, %q) = %v; want %v", test.str1, test.str2, result, test.expected)
			}
		
	}
}