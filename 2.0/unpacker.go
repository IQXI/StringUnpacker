package main

import (
	"fmt"
	"strconv"
)

type literal struct {
	value string
	count int
}

func string_to_slice(s string) []string {
	sl := make([]string, 0)
	for _, r := range s {
		sl = append(sl, string(r))
	}
	return sl
}

func find_digits(sl []string) (int, int) {
	number := ""
	for i := 0; i < len(sl); i++ {
		if _, err := strconv.Atoi(sl[i]); err == nil {
			number += sl[i]
		} else {
			if number == "" {
				return 0, 1
			} else {
				num, _ := strconv.Atoi(number)
				return i, num
			}
		}
	}
	if number == "" {
		return 0, 1
	} else {
		num, _ := strconv.Atoi(number)
		return 0, num
	}
}

func formatting_slice(sl []string) []literal {
	lt := make([]literal, 0)
	for i := 0; i < len(sl); i++ {
		if _, err := strconv.Atoi(sl[i]); err != nil {
			if sl[i] == "\\" {
				if i+1 < len(sl)-1 {
					j, count := find_digits(sl[i+2:])
					lt = append(lt, literal{value: sl[i+1], count: count})
					i += j
				} else {
					j, count := find_digits(sl[i+2:])

					if i > 0 {
						if sl[i-1] == "\\" && j == 0 {
							continue
						}
					}

					lt = append(lt, literal{value: sl[i+1], count: count})
					i += j
				}
			} else {
				if i+1 < len(sl)-1 {
					j, count := find_digits(sl[i+1:])
					lt = append(lt, literal{value: sl[i], count: count})
					i += j
				} else {
					lt = append(lt, literal{value: sl[i], count: 1})
				}

			}
		}
	}
	return lt

}

func Unpacker(s string) string {
	sl := string_to_slice(s)
	lt := formatting_slice(sl)

	result := ""
	for i := 0; i < len(lt); i++ {
		for j := 0; j < lt[i].count; j++ {
			result += lt[i].value
		}
	}

	return result

}

func main() {
	type pair struct {
		i string
		s string
	}
	test := []pair{
		{"", ""},
		{"a4bc2d5e", "aaaabccddddde"},
		{"a15b11", "aaaaaaaaaaaaaaabbbbbbbbbbb"},
		{"abcd", "abcd"},
		{"a10b20", "aaaaaaaaaabbbbbbbbbbbbbbbbbbbb"},
		{"45", ""},
		{"012", ""},

		{`qwe\412`, `qwe444444444444`},
		{`qwe\4\5`, `qwe45`},
		{`qwe\45`, `qwe44444`},
		{`qwe\\5`, `qwe\\\\\`},
		{`qwe\\2\3\\2`, `qwe\\3\\`},
		{`\45q2w3e10`, `44444qqwwweeeeeeeeee`},
	}
	for _, t := range test {
		unpacked_s := Unpacker(t.i)
		if t.s == unpacked_s {
			fmt.Printf("Original %s:%s - %s\n", t.s, unpacked_s, "OK")
		} else {
			fmt.Printf("Original %s:%s - %s\n", t.s, unpacked_s, "FAIL")
		}
	}
}
