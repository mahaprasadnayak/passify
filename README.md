# Passify: A Golang Package for Checking Password Strength

**Passify** is a lightweight and efficient Golang package designed to evaluate the strength of passwords. With Passify, you can easily determine if a password meets security standards by analyzing its complexity and providing feedback.

## Why Passify?

- **Simple API**: Easily integrate password strength checks into your applications.
- **Customizable Rules**: Define your own password complexity requirements.
- **Lightweight**: Designed to be fast and efficient, perfect for high-performance systems.
- **Informative Feedback**: Helps users create stronger passwords by providing actionable feedback.

## Features

- **Strength Scoring**: Assigns a strength score based on password length, character diversity, and patterns.
- **Customizable Policies**: Supports custom rules for minimum length, required character sets, etc.
- **Feedback Messages**: Provides recommendations for improving weak passwords.
- **No External Dependencies**: Built purely in Go, ensuring portability and simplicity.

## Installation

To install Passify, use the `go get` command:

```bash
go get github.com/mahaprasadnayak/passify
```

## Example in Go that demonstrates how to use Passify in your Go application
```
	package main

import (
	"fmt"
	"github.com/mahaprasadnayak/passify"
)

func main() {
	password := "Passw0rd!"
	result := passify.CheckStrength(password)
	fmt.Printf("Strength Score: %s\n", result)
}
```
## Key Features
- Custom Rules: Adjust password policies, such as required length or specific character sets.
- Actionable Feedback: Provides insights to help users improve weak passwords.

## Use Case
- Authentication Systems: Enforce secure password creation during user registration.
- Password Validation: Check password strength during login or account updates.