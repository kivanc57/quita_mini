package calc

import "math"

func CalcTTR(lemmaFreqMap, tokenFreqMap map[string]int) float64 {
	totalTokens := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return freq
	})
	ttr := float64(len(lemmaFreqMap)) / float64(totalTokens)
	return ttr
}

func CalcEntropy(tokenFreqMap map[string]int) float64 {
	var entropy float64 = 0
	totalTokens := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return freq
	})

	for _, freq := range tokenFreqMap {
		p := float64(freq) / float64(totalTokens)
		entropy += p * math.Log2(p)
	}
	return -entropy
}

func CalcNormEntropy(tokenFreqMap map[string]int) float64 {
	entropy := CalcEntropy(tokenFreqMap)
	totalTypes := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return 1
	})

	normEntropy := entropy / math.Log2(float64(totalTypes))
	return normEntropy
}

func CalcAvgTokenLen(tokenFreqMap map[string]int) float64 {
	totalChars := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return len(word) * freq
	})
	totalTokens := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return freq
	})
	avgTokenLen := float64(totalChars) / float64(totalTokens)
	return avgTokenLen
}

func CalcAvgTypeLen(lemmaFreqMap map[string]int) float64 {
	totalTypes := calculateTotal(lemmaFreqMap, func(word string, freq int) int {
		return 1
	})
	totalChars := calculateTotal(lemmaFreqMap, func(word string, freq int) int {
		return freq
	})

	avgTypeLen := float64(totalChars) / float64(totalTypes)
	return avgTypeLen
}

func CalcTokenLenSD(tokenFreqMap map[string]int) float64 {
	var sd, sumVariance, variance float64

	avgTokenLen := CalcAvgTokenLen(tokenFreqMap)
	totalTokens := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return freq
	})

	for token, freq := range tokenFreqMap {
		dif := float64(len(token)) - float64(avgTokenLen)
		sumVariance += float64(freq) * (dif * dif)
	}
	variance = sumVariance / float64(totalTokens)
	sd = math.Sqrt(variance)
	return sd
}

func CalcTypeLenSD(lemmaFreqMap map[string]int) float64 {
	var sd, sumVariance, variance float64

	totalTypes := calculateTotal(lemmaFreqMap, func(word string, freq int) int {
		return 1
	})
	avgTypeLen := CalcAvgTypeLen(lemmaFreqMap)

	for lemma, freq := range lemmaFreqMap {
		if freq == 1 {
			dif := float64(len(lemma)) - float64(avgTypeLen)
			sumVariance += (dif * dif)
		}
	}
	variance = sumVariance / float64(totalTypes)
	sd = math.Sqrt(variance)
	return sd
}

func CalcYulesK(tokenFreqMap map[string]int) float64 {
	var sumFreqDif float64

	totalTokens := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return freq
	})

	for _, freq := range tokenFreqMap {
		sumFreqDif += float64(freq) * (float64(freq) - 1)
	}
	k := (10000 * sumFreqDif) / (float64(totalTokens) * float64(totalTokens))
	return k
}

func CalcAdjustedModulus(lemmaFreqMap, tokenFreqMap map[string]int) float64 {
	ttr := CalcTTR(lemmaFreqMap, tokenFreqMap)
	adjustudModulus := float64(ttr) / math.Log2(float64(len(tokenFreqMap)))
	return adjustudModulus
}

func CalcPercOfTextHapax(tokenFreqMap map[string]int) float64 {
	totalTokens := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return freq
	})

	totalHapax := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		if freq == 1 {
			return 1
		}
		return 0
	})

	percOfTheTextHapax := (float64(totalHapax) / float64(totalTokens)) * 100
	return percOfTheTextHapax		
}

func CalcPercOfTypeHapax(lemmaFreqMap map[string]int) float64 {
	totalHapax := calculateTotal(lemmaFreqMap, func(word string, freq int) int {
		if freq == 1 {
			return 1
		}
		return 0
	})

	totalTypes := calculateTotal(lemmaFreqMap, func(word string, freq int) int {
		return 1
	})
	percOfTheTypeHapax := (float64(totalHapax) / float64(totalTypes)) * 100
	return percOfTheTypeHapax
}

func CalcR1(tokenFreqMap  map[string]int) float64 {
	totalHapax := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		if freq == 1 {
			return 1
		}
		return 0
	})

	totalTypes := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return 1
	})
	r1 := float64(totalHapax) / float64(totalTypes)
	return r1
}

func CalcRIndex(lemmaFreqMap, tokenFreqMap map[string]int) float64 {
	totalHapax := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		if freq == 1 {
			return 1
		}
		return 0
	})

	totalTypes := calculateTotal(lemmaFreqMap, func(word string, freq int) int {
		return 1
	})
	rIndex := float64(totalHapax) / float64(totalTypes)
	return rIndex
}

func CalcRrIndex(tokenFreqMap map[string]int) float64 {
	totalHapax := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		if freq == 1 {
			return 1
		}
		return 0
	})

	totalTokens := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return freq
	})
	rrIndex := float64(totalHapax) / float64(totalTokens) * 100
	return rrIndex
}

func CalcLIndex(lemmaFreqMap, tokenFreqMap map[string]int) float64 {
	totalTypes := calculateTotal(lemmaFreqMap, func(word string, freq int) int {
		return 1
	})
	totalTokens := calculateTotal(tokenFreqMap, func(word string, freq int) int {
		return freq
	})
	l := float64(totalTypes) / math.Log(float64(totalTokens))
	return l
}

