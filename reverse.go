// Пакет morestrings реализует дополнительные функции  
// для управления UTF-8 закодированными строками сверх того, 
// что предусмотрено в стандартном пакете "strings".
package morestrings

// ReverseRunes возвращает строку аргументов, 
// перевернутую по рунам слева направо.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}