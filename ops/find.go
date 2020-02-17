package ops

import "strings"

func Find(line string) Op {
	switch {
	case len(line) >= 1 && line[:1] == ":":
		name := strings.Trim(line[1:], " ")
		if len(name) != 1 {
			return BAD{line, "label more than 1 character"}
		}
		return LBL{name}
	case len(line) >= 2 && line[:2] == "<>":
		cont := strings.Trim(line[2:], " ")
		return COM{cont}
	case len(line) >= 2 && line[:2] == "><":
		msg := strings.Trim(line[2:], " ")
		return DIE{msg}
	case len(line) >= 2 && line[:2] == "->":
		switch {
		case strings.Contains(line[2:], "??"):
			parts := strings.Split(line[2:], "??")
			if len(parts) != 2 {
				return BAD{line, "wrong number of parts"}
			}

			opi := strings.IndexAny(parts[1], AllCmps())
			if opi == -1 {
				return BAD{line, "no cmp"}
			}

			to := strings.Trim(parts[0], " ")
			lh := strings.Trim(parts[1][:opi], " ")
			op := strings.Trim(parts[1][opi:opi+1], " ")
			rh := strings.Trim(parts[1][opi+1:], " ")

			return JMP{to, lh, op, rh}
		case strings.Contains(line[2:], "?"):
			parts := strings.Split(line[2:], "?")
			if len(parts) != 2 {
				return BAD{line, "wrong number of parts"}
			}

			to := strings.Trim(parts[0], " ")
			ifs := strings.Trim(parts[1], " ")

			return SKP{to, ifs}
		default:
			to := strings.Trim(line[2:], " ")

			return HOP{to}
		}
	case len(line) >= 1 && line[:1] == ">":
		switch {
		case strings.Contains(line[1:], ">>"):
			parts := strings.Split(line[1:], ">>")
			if len(parts) != 2 {
				return BAD{line, "wrong number of parts"}
			}

			prompt := parts[0]
			varn := strings.Trim(parts[1], " ")

			return NIN{prompt, varn}
		case strings.Contains(line[1:], ">"):
			parts := strings.Split(line[1:], ">")
			if len(parts) != 2 {
				return BAD{line, "wrong number of parts"}
			}

			prompt := parts[0]
			varn := strings.Trim(parts[1], " ")

			return TIN{prompt, varn}
		default:
			text := line[1:]

			return OUT{text}
		}
	default:
		switch {
		case strings.Contains(line, ">>"):
			parts := strings.Split(line, ">>")
			if len(parts) != 2 {
				return BAD{line, "wrong number of parts"}
			}

			expr := strings.Trim(parts[0], " ")
			varn := strings.Trim(parts[1], " ")

			return NAS{expr, varn}
		case strings.Contains(line, ">"):
			parts := strings.Split(line, ">")
			if len(parts) != 2 {
				return BAD{line, "wrong number of parts"}
			}

			expr := strings.Trim(parts[0], " ")
			varn := strings.Trim(parts[1], " ")

			return TAS{expr, varn}
		default:
			if strings.Trim(line, " ") == "" {
				return EMP{}
			} else {
				return BAD{line, "unknown"}
			}
		}
	}
}
