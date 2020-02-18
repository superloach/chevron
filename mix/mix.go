package mix

import "github.com/superloach/chevron/vars"

func Mix(t string, v *vars.Vars) (string, error) {
	text := []rune(t)
	for i := 0; i < len(text)-1; i++ {
		if text[i] == '^' {
			n := text[i+1:i+2]
			if (n[0] == '_' || n[0] == ':') && i < len(text)-2 {
				n = append(n, text[i+2])
			}
			a, err := v.Get(string(n))
			if err != nil {
				return "", err
			}
			newtxt := make([]rune, len(text) + len(a))
			newtxt = append(newtxt, text[:i]...)
			newtxt = append(newtxt, []rune(a)...)
			newtxt = append(newtxt, text[i+1+len(n):]...)
			text = newtxt
			i += len(a) - 1
		}
	}
	return string(text), nil
}
