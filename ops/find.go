package ops

import "strings"

func Find(line string) Op {
	switch {
	case line[0] == '<':
		switch {
		case line[1] == '>':
			if line[2] == ':' {
				return LBL{line[3:]}
			}
			return COM{line[2:]}
		default:
			return nil
		}
	case line[0] == '>':
		switch {
		case strings.Contains(line[1:], ">>"):
			parts := strings.Split(line[1:], ">>")
			if len(parts) != 2 {
				return nil
			}
			if parts[1][0] == '^' {
				parts[1] = parts[1][1:]
			}
			return NIN{parts[0], parts[1]}
		case strings.Contains(line[1:], ">"):
			parts := strings.Split(line[1:], ">")
			if len(parts) != 2 {
				return nil
			}
			if parts[1][0] == '^' {
				parts[1] = parts[1][1:]
			}
			return TIN{parts[0], parts[1]}
		default:
			return OUT{line[1:]}
		}
	default:
		switch {
		case strings.Contains(line, "<<"):
			parts := strings.Split(line, "<<")
			if len(parts) != 2 {
				return nil
			}
			if parts[1][0] == '^' {
				parts[1] = parts[1][1:]
			}
			return NAS{parts[0], parts[1]}
		case strings.Contains(line, "<"):
			return nil
		default:
			return nil
		}
	}
}
