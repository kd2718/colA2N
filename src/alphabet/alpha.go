package alphabet

type AlphaRune []rune

func (alphabet *AlphaRune) FirstIdx(beta rune) int64 {
	for idx, v := range *alphabet {
		if beta == v {
			return int64(idx)
		}
	}
	return -1
}
