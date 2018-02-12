package searchSubstring


const p = 13
const r = 4294967296 //pow(2,32)


func powInt(x int, k int) int{ // возведение x в степень k
	if k==0 {
		return 1
	}
	if k==1 {
		return x
	}

	res := 1
	for ; k != 0; {
		if (k & 1) != 0 {
			res *= x
		}
		x *= x
		k >>= 1
	}
	return res
}
func hash(s string, Pi []int) int {
	m:=len(s)
	res:=Pi[m-2]*int(s[0])//if m=1?
	for i:=1;i<m-1;i++{
		res+=Pi[m-i-2]*int(s[i])
	}
	res+=int(s[m-1])
	return res
}

func evaluatePi(m int) []int{//p^1,....,p^m
	res:=make([]int,m)
	pi:=p
	for i:=0;i<m;i++{
		res[i]=pi
		pi*=p
	}
	return res
}

func RabinKarp(text string, sample string) []int {
	res := make([]int, 0)
	n := len(text)
	m := len(sample)

	Pi:=evaluatePi(m)
	hashSubText := hash(text[0:m],Pi)
	hashSample := hash(sample,Pi)

	for i := 0; i <= n-m; i++ {
		if hashSubText == hashSample {
			res = append(res, i)
		}
		if i<(n-m){
			hashSubText = (p*hashSubText - Pi[m-1]*int(text[i]) + int(text[i+m])) % r
		}
	}
	return res
}


