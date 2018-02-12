package searchSubstring

func compareStr(str1 string, str2 string) bool{
	n:=len(str1)
	m:=len(str2)
	if n!=m {
		return false
	}

	for i:=0;i<n;i++{
		if str1[i] != str2[i] {
			return false
		}
	}

	return true
}


func Simple(text string, sample string) []int{//возвращает допустимые сдвиги
	res:=make([]int,0)
	n:= len(text)
	m:=len(sample)
	for i:=0; i <= n-m; i++{
		if compareStr(text[i:i+m],sample){
			res=append(res,i)
		}
	}
	return res
}