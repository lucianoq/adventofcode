package main

import (
	"fmt"
	"math"
)

const input = 3005290

func main() {
	//biggest power of 3 less than number
	maxPower3 := int(math.Pow(3, math.Floor(math.Log(float64(input))/math.Log(float64(3)))))

	var result int

	switch {
	case maxPower3 == input:
		// powers of 3 have the number itself as solution
		result = maxPower3

		// while between a power of 3 and the next
	case input <= maxPower3*2:

		// first half increase by 1
		result = input - maxPower3

	default:

		// second half increase by 2
		result = 2*input - 3*maxPower3
	}

	fmt.Println(result)
}

//1 1
//2 1
//3 3
//4 1
//5 2
//6 3
//7 5
//8 7
//9 9
//10 1
//11 2
//12 3
//13 4
//14 5
//15 6
//16 7
//17 8
//18 9
//19 11
//20 13
//21 15
//22 17
//23 19
//24 21
//25 23
//26 25
//27 27
//28 1
//29 2
//30 3
//31 4
//32 5
//33 6
//34 7
//35 8
//36 9
//37 10
//38 11
//39 12
//40 13
//41 14
//42 15
//43 16
//44 17
//45 18
//46 19
//47 20
//48 21
//49 22
//50 23
//51 24
//52 25
//53 26
//54 27
//55 29
//56 31
//57 33
//58 35
//59 37
//60 39
//61 41
//62 43
//63 45
//64 47
//65 49
//66 51
//67 53
//68 55
//69 57
//70 59
//71 61
//72 63
//73 65
//74 67
//75 69
//76 71
//77 73
//78 75
//79 77
//80 79
//81 81
//82 1
//83 2
//84 3
//85 4
//86 5
//87 6
//88 7
//89 8
//90 9
//91 10
//92 11
//93 12
//94 13
//95 14
//96 15
//97 16
//98 17
//99 18
