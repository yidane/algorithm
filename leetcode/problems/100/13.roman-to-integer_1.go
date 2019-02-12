package problems

var symbol = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XX":   20,
	"XXX":  30,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
	"CC":   200,
	"CCC":  300,
	"CD":   400,
	"D":    500,
	"DC":   600,
	"DCC":  700,
	"DCCC": 800,
	"CM":   900,
	"M":    1000,
	"MM":   2000,
	"MMM":  3000,
}

/*
I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900
*/

func romanToInt(s string) int {
	num := 0
	for i := 0; i < len(s); {
		n := 1
		ns := s[i : i+n]
		_, ok := symbol[ns]
		for ok {
			n++
			if n+i > len(s) {
				break
			}

			ns = s[i : i+n]
			_, ok = symbol[ns]
		}

		n = n - 1
		num += symbol[s[i:i+n]]
		i = i + n
	}

	return num
}

func romanToInt1(s string) int {
	roman := make(map[uint8]int)
	roman['M'] = 1000
	roman['D'] = 500
	roman['C'] = 100
	roman['L'] = 50
	roman['X'] = 10
	roman['V'] = 5
	roman['I'] = 1
	z := 0
	for i := 0; i < len(s)-1; i++ {
		if roman[s[i]] < roman[s[i+1]] {
			z -= roman[s[i]]
		} else {
			z += roman[s[i]]
		}
	}
	return z + roman[s[len(s)-1]]
}
