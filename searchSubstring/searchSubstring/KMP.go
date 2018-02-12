package searchSubstring


func computePrefixFunction(P string) []int {
	m := len(P)
	prefix := make([]int, m)

	prefix[0] = 0
	k := 0 //длина макс префикса, который является суффиксом
	for q := 1; q < m; q++ {
		for ; (k > 0) && (P[k] != P[q]); {
			k = prefix[k-1]
		}
		if P[k] == P[q] {
			k++
		}
		prefix[q] = k
	}
	return prefix
}

func KMP(text string, sample string) []int{
	res:=make([]int,0)
	n:=len(text)
	m:=len(sample)
	prefix:=computePrefixFunction(sample)
	q:=0
	for i:=0;i<n;i++{
		for ;(sample[q]!=text[i]) && (q>0);{
			q=prefix[q-1]
		}
		if sample[q]==text[i]{
			q++
		}
		if q==m{
			res=append(res,i-m+1)
			q=prefix[q-1]
		}
	}
	return res
}

