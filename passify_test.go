package Passify

import "testing"

func TestCheckStrength(t *testing.T) {
    tests := []struct {
        password string
        expected string
    }{
        {"Password123!", "Very strong"}, // Meets all criteria
        {"Password123", "Very weak, please retry"}, // Missing special character
        {"password123!", "Very weak, please retry"}, // Missing uppercase character
        {"PASSWORD123!", "Very weak, please retry"}, // Missing lowercase character
        {"Password!", "Very weak, please retry"}, // Missing digit
        {"Pass1!", "Very weak, please retry"}, // Too short
        {"", "Very weak, please retry"}, // Empty password
    }

    for _, tt := range tests {
        t.Run(tt.password, func(t *testing.T) {
            result := Check_strength(tt.password)
            if result != tt.expected {
                t.Errorf("Check_strength(%q) = %q; want %q", tt.password, result, tt.expected)
            }
        })
    }
}
