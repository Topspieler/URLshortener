package controllers

func idToShortUrl(id int) string {

	base := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]int, 0)

	for id > 0 {
		s = append(s, id%62)
		id /= 62
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	runeslice := make([]rune, len(s))

	for k, n := range s {
		runeslice[k] = base[n]
	}
	return string(runeslice)
}

func shortUrlToId(s string) int {

	base := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	m := make(map[rune]int)

	for n, ba := range base {
		m[ba] = n
	}

	srunes := []rune(s)

	id := 0

	for _, k := range srunes {
		id = id*62 + m[k]
	}

	return id
}
