package caesar_cipher

import sentence_score "github.com/cryptography-research-lab/go-sentence-score"

// CrackForOne 返回可能性最大的一个结果
func CrackForOne(encryptText string) string {

	// 先穷举所有结果
	resultSlice := make([]string, 26)
	for i := 0; i < 26; i++ {
		resultSlice[i] = Decrypt(encryptText, i)
	}

	// 然后计算每个的可能性
	sort := sentence_score.CalculateScoreAndDescSort(resultSlice)

	// 选出一个可能性最大的返回
	return sort[0].Sentence
}

// CrackForAll 返回所有的结果，按照可能性排序
func CrackForAll(encryptText string) []string {

	// 先穷举所有结果
	resultSlice := make([]string, 26)
	for i := 0; i < 26; i++ {
		resultSlice[i] = Decrypt(encryptText, i)
	}

	// 然后计算每个的可能性
	sort := sentence_score.CalculateScoreAndDescSort(resultSlice)

	// 选出一个可能性最大的返回
	resultSlice = make([]string, 26)
	for index, x := range sort {
		resultSlice[index] = x.Sentence
	}
	return resultSlice
}
