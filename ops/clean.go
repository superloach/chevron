package ops

import "strings"

func Clean(s string) string {
	s = strings.ReplaceAll(s, ">", "^_r")
	s = strings.ReplaceAll(s, "<", "^_l")
	s = strings.ReplaceAll(s, "^", "^_u")
	s = strings.ReplaceAll(s, "?", "^_q")
	s = strings.ReplaceAll(s, "-", "^_d")
	s = strings.ReplaceAll(s, "_", "^_s")
	s = strings.ReplaceAll(s, "=", "^_e")
	s = strings.ReplaceAll(s, ":", "^_o")
	s = strings.ReplaceAll(s, "\n", "^_n")
	s = strings.ReplaceAll(s, "~", "^_t")
	s = strings.ReplaceAll(s, "`", "^_b")
	s = strings.ReplaceAll(s, ",", "^_m")
	s = strings.ReplaceAll(s, " ", "^_p")
	return s
}
