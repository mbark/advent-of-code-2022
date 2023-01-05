package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/maths"
	"github.com/mbark/advent-of-code-2022/util"
)

func main() {
	var sum int
	for _, n := range util.ReadInput(Input, "\n") {
		sum += toDecimal(n)
	}

	fmt.Printf("first: %s\n", ToSnafu(sum))
}

func toDecimal(snafu string) int {
	var sum int
	for i := 0; i < len(snafu); i++ {
		idx := len(snafu) - 1 - i
		var times int
		switch snafu[idx] {
		case '2':
			times = 2
		case '1':
			times = 1
		case '0':
			times = 0
		case '-':
			times = -1
		case '=':
			times = -2
		}

		sum += times * maths.PowInt(5, i)
	}

	return sum
}

var snafuChar = map[int]byte{
	0: '0', 1: '1', 2: '2', 3: '=', 4: '-',
}

func ToSnafu(decimal int) string {
	var snafu []byte

	for decimal > 0 {
		num := decimal / 5
		rem := decimal % 5
		decimal = num

		snafu = append(snafu, snafuChar[rem])
		if rem > 2 {
			decimal += 1
		}
	}

	if len(snafu) == 0 {
		return "0"
	}

	var reversed []byte
	for i := len(snafu) - 1; i >= 0; i-- {
		reversed = append(reversed, snafu[i])
	}

	return string(reversed)
}

const testInput = `
1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
`

const Input = `
2=1121==100=22
22-0==122-1-
10-10-2=012=2=1==1
1--000-=12-
1=2=
1==--2--22-=
11-=20==-1-211=1-21
1=2-0200=-012=
10=2--211
2=
1-=
1==2
21==-=202=2
2==1=0==12=11=
1=-00=201=2-==
1-0=0=2=-2-0=-1=
1=2==
2=1202211-
2-=11=
122=01-
20=000=12-210-00-0-
1020=-212002
110==21100=-
1-=1=11-2-1=
22-=00011=-01-1-
20-2=22
1=02=01002---=0
12-==-2==020
1-22
10
2
20121
1001--102=-
1=20==011=-=
2100=--1112
102=
1--=2-20100---1=2-
100=02-2=00010
2=212-01-1-200--=-
1001121
1222
1==-1-
1=1=1---112---2=22
2--02020--010
2=02-2-20====12
2-1=21-2=202=2=2
1-20=120120-102=
22=-110
1=0
212-2=00220102
1-012--1=
112==0=-
202
1=22--
1212=2--20
2020==0=00=00=-
1==010
1=-0-201=-
1===0
10=
122-0-=02
1=-=12-110--02=1
1==000=11=1-1121
12020020-2=0--1220
12=-2=102=--2-=0-
2==211111=2-=02=0
121021
22=-1=2
1-0-01=-1220==1
20=122
1110-0
2-210====
2---0=2=
101=1-1-
200=1-00
1-10--2==1-10
10=-00=2=0=010-2-
2=00-100=0=1-2-
21==-
101=-0202
21020-=2102
2=0--
1-0-0-0--1=
21220211-
1=01-221=22
1--=-11=
2212
1-00022--122=2001
111
1=-1-2-
2=21-212122=21=1
10---0120-220
11-000=0
1=220=2-200=20
1=2=-==0=102-0=-0212
100-
1--2=11=
20
1=--0-=20--2=--=-
1==12-10
1001-=11=-2=11-0
12--2
`
