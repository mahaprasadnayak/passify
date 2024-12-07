package Passify

import (
	"unicode"
)

const (
	MIN_LENGTH      = 8
)

func Check_strength(password string)string{
	length:=len(password)
	var (has_upper ,has_lower, has_spl_char ,has_digit bool
	score int)

	for _,ch:= range password{
		switch {
		case unicode.IsUpper(ch):
			has_upper = true
		case unicode.IsLower(ch):
			has_lower = true
		case unicode.IsDigit(ch):
			has_digit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch)|| unicode.IsSpace(ch):
			has_spl_char=true
		}
	}
	
	if length >= MIN_LENGTH && has_upper && has_lower&& has_spl_char && has_digit{
		score++
	}
	switch score {
    case 5:
        return "Very strong"
    case 4:
        return "Strong"
    case 3:
        return "Moderate"
    case 2:
        return "Weak"
    default:
        return "Very weak, please retry"
    }

}


