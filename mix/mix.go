package mix

import "github.com/superloach/chevron/vars"

func Mix(text string, v *vars.Vars) (string, error) {
	for i := 0; i < len(text)-1; i++ {
		if text[i] == '^' {
			n := string(text[i+1])
			if (n == "_" || n == ":") && i < len(text)-2 {
				n += string(text[i+2])
			}
			a, err := v.Get(n)
			if err != nil {
				return "", err
			}
			text = text[:i] + a + text[i+1+len(n):]
			i += len(a) - 1
		}
	}
	return text, nil
}
