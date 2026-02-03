package main

import (
	"fmt"
	"regexp"
	"strings"
)

// ============================================================================
// EBNF NOTATION EXAMPLES IN GO
// ============================================================================
//
// This file demonstrates how EBNF grammar rules translate to Go code
// Based on: https://go.dev/ref/spec#Notation

// ============================================================================
// 1. ALTERNATION (|) - Choose ONE option
// ============================================================================
// EBNF: Boolean = "true" | "false" .

func isBoolean(s string) bool {
	return s == "true" || s == "false"
}

// Example usage:
// isBoolean("true")   // true
// isBoolean("false")  // true
// isBoolean("maybe")  // false

// ============================================================================
// 2. GROUPING () - Group expressions together
// ============================================================================
// EBNF: Sign = "+" | "-" .
//       SignedNumber = [ Sign ] Number .

type SignedNumber struct {
	Sign   string // optional: "+" or "-"
	Number int
}

func parseSignedNumber(s string) (SignedNumber, error) {
	s = strings.TrimSpace(s)
	sn := SignedNumber{}

	// Optional sign (grouping with alternation)
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		sn.Sign = string(s[0])
		s = s[1:]
	} else {
		sn.Sign = "+" // default positive
	}

	// Parse number
	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		return sn, err
	}
	sn.Number = num
	return sn, nil
}

// Example usage:
// parseSignedNumber("+42")   // {"+", 42}
// parseSignedNumber("-15")   // {"-", 15}
// parseSignedNumber("99")    // {"+", 99}

// ============================================================================
// 3. OPTION [] - Zero or one occurrence (optional)
// EBNF: FileExtension = [ "." identifier ] .

type File struct {
	Name      string
	Extension string // optional
}

func parseFilename(filename string) File {
	parts := strings.Split(filename, ".")

	file := File{
		Name:      parts[0],
		Extension: "", // no extension by default
	}

	if len(parts) > 1 {
		file.Extension = parts[1] // optional extension
	}

	return file
}

// Example usage:
// parseFilename("document.txt")  // {Name: "document", Extension: "txt"}
// parseFilename("README")        // {Name: "README", Extension: ""}

// ============================================================================
// 4. REPETITION {} - Zero or more occurrences
// EBNF: Digits = { Digit } .
//       Digit = "0" … "9" .

func isDigits(s string) bool {
	// Match zero or more digits
	matched, _ := regexp.MatchString(`^\d*$`, s)
	return matched
}

// Example usage:
// isDigits("")      // true (zero occurrences)
// isDigits("5")     // true (one digit)
// isDigits("12345") // true (many digits)
// isDigits("12a45") // false

// ============================================================================
// 5. RANGE … - Set of characters
// EBNF: Digit = "0" … "9" .
//       Letter = "a" … "z" | "A" … "Z" .

func isDigit(c rune) bool {
	return c >= '0' && c <= '9' // range 0 to 9
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isLowerLetter(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func isUpperLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

// Example usage:
// isDigit('5')      // true
// isLetter('A')     // true
// isLetter('z')     // true
// isLetter('1')     // false

// ============================================================================
// 6. COMPLETE EXAMPLE - Identifier
// ============================================================================
// EBNF: Identifier = letter { letter | unicode_digit | "_" } .
//       letter = "a"…"z" | "A"…"Z" | "_" .

func isValidIdentifier(s string) bool {
	if len(s) == 0 {
		return false
	}

	// First character must be letter or underscore
	firstChar := rune(s[0])
	if !isLetter(firstChar) && firstChar != '_' {
		return false
	}

	// Remaining characters: letter, digit, or underscore
	for _, c := range s[1:] {
		if !isLetter(c) && !isDigit(c) && c != '_' {
			return false
		}
	}

	return true
}

// Example usage:
// isValidIdentifier("name")       // true
// isValidIdentifier("_private")   // true
// isValidIdentifier("var123")     // true
// isValidIdentifier("MY_CONST")   // true
// isValidIdentifier("123var")     // false (starts with digit)
// isValidIdentifier("my-var")     // false (contains hyphen)

// ============================================================================
// 7. COMPLETE EXAMPLE - Integer Literal
// ============================================================================
// EBNF: IntLit = DecimalLit | HexLit .
//       DecimalLit = ( "1"…"9" ) { DecimalDigit } | "0" .
//       HexLit = "0" ( "x" | "X" ) HexDigit { HexDigit } .

func isValidInteger(s string) bool {
	// Try decimal
	if isValidDecimal(s) {
		return true
	}
	// Try hex
	if isValidHex(s) {
		return true
	}
	return false
}

func isValidDecimal(s string) bool {
	if len(s) == 0 {
		return false
	}

	// Single "0" is valid
	if s == "0" {
		return true
	}

	// First digit must be 1-9
	if s[0] < '1' || s[0] > '9' {
		return false
	}

	// Remaining digits must be 0-9
	for _, c := range s[1:] {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func isValidHex(s string) bool {
	if len(s) < 3 {
		return false
	}

	// Must start with 0x or 0X
	if s[0] != '0' || (s[1] != 'x' && s[1] != 'X') {
		return false
	}

	// Rest must be hex digits (0-9, a-f, A-F)
	for _, c := range s[2:] {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// Example usage:
// isValidInteger("0")         // true
// isValidInteger("123")       // true
// isValidInteger("0xFF")      // true
// isValidInteger("0xDEADBEEF") // true
// isValidInteger("00")        // false

// ============================================================================
// 8. COMPLETE EXAMPLE - For Statement
// ============================================================================
// EBNF: ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .

type ForStatement struct {
	ConditionType string // "condition", "clause", "range", or "infinite"
	Content       string
}

func parseForStatement(stmt string) (ForStatement, error) {
	stmt = strings.TrimSpace(stmt)

	if !strings.HasPrefix(stmt, "for") {
		return ForStatement{}, fmt.Errorf("not a for statement")
	}

	// Remove "for" keyword
	content := strings.TrimPrefix(stmt, "for")
	content = strings.TrimSpace(content)

	// Determine which type of for loop
	fs := ForStatement{Content: content}

	if content == "" {
		// for { ... } - infinite loop
		fs.ConditionType = "infinite"
	} else if strings.Contains(content, ":=") || strings.Contains(content, ";") {
		// for i := 0; i < 10; i++ { ... } - C-style loop
		fs.ConditionType = "clause"
	} else if strings.HasPrefix(content, "range") {
		// for i, v := range list { ... } - range loop
		fs.ConditionType = "range"
	} else {
		// for x < 10 { ... } - condition-based loop
		fs.ConditionType = "condition"
	}

	return fs, nil
}

// Example usage:
// parseForStatement("for x < 10 { ... }")           // condition
// parseForStatement("for i := 0; i < 10; i++ { ... }") // clause
// parseForStatement("for i, v := range list { ... }") // range
// parseForStatement("for { ... }")                     // infinite

// ============================================================================
// 9. PRACTICAL EXAMPLE - Function Call
// ============================================================================
// EBNF: FunctionCall = identifier "(" [ ArgumentList ] ")" .
//       ArgumentList = Argument { "," Argument } .
//       Argument = Expression | identifier "=" Expression .

type FunctionCall struct {
	Name      string
	Arguments []string
}

func parseFunctionCall(call string) (FunctionCall, error) {
	// Find opening parenthesis
	parenIdx := strings.Index(call, "(")
	if parenIdx == -1 {
		return FunctionCall{}, fmt.Errorf("no opening parenthesis")
	}

	name := strings.TrimSpace(call[:parenIdx])

	// Find closing parenthesis
	closeIdx := strings.LastIndex(call, ")")
	if closeIdx == -1 {
		return FunctionCall{}, fmt.Errorf("no closing parenthesis")
	}

	// Parse arguments (comma-separated)
	argsStr := strings.TrimSpace(call[parenIdx+1 : closeIdx])
	args := []string{}

	if argsStr != "" { // optional arguments
		parts := strings.Split(argsStr, ",")
		for _, part := range parts {
			args = append(args, strings.TrimSpace(part))
		}
	}

	return FunctionCall{
		Name:      name,
		Arguments: args,
	}, nil
}

// Example usage:
// parseFunctionCall("fmt.Println()")              // {Name: "fmt.Println", Args: []}
// parseFunctionCall("fmt.Println(\"Hello\")")     // {Name: "fmt.Println", Args: ["Hello"]}
// parseFunctionCall("add(2, 3)")                  // {Name: "add", Args: ["2", "3"]}

// ============================================================================
// MAIN - Demonstrate all examples
// ============================================================================

func main() {
	fmt.Println("=" * 70)
	fmt.Println("EBNF NOTATION EXAMPLES IN GO")
	fmt.Println("=" * 70)

	// 1. Alternation
	fmt.Println("\n1. ALTERNATION (|) - Choose ONE option")
	fmt.Println("   isBoolean(\"true\"):", isBoolean("true"))
	fmt.Println("   isBoolean(\"false\"):", isBoolean("false"))
	fmt.Println("   isBoolean(\"maybe\"):", isBoolean("maybe"))

	// 2. Grouping & Alternation
	fmt.Println("\n2. GROUPING () - Group expressions")
	sn1, _ := parseSignedNumber("+42")
	fmt.Printf("   parseSignedNumber(\"+42\"): %+v\n", sn1)
	sn2, _ := parseSignedNumber("-15")
	fmt.Printf("   parseSignedNumber(\"-15\"): %+v\n", sn2)
	sn3, _ := parseSignedNumber("99")
	fmt.Printf("   parseSignedNumber(\"99\"): %+v\n", sn3)

	// 3. Option
	fmt.Println("\n3. OPTION [] - Zero or one occurrence")
	f1 := parseFilename("document.txt")
	fmt.Printf("   parseFilename(\"document.txt\"): %+v\n", f1)
	f2 := parseFilename("README")
	fmt.Printf("   parseFilename(\"README\"): %+v\n", f2)

	// 4. Repetition
	fmt.Println("\n4. REPETITION {} - Zero or more occurrences")
	fmt.Println("   isDigits(\"\"):", isDigits(""))
	fmt.Println("   isDigits(\"12345\"):", isDigits("12345"))
	fmt.Println("   isDigits(\"12a45\"):", isDigits("12a45"))

	// 5. Range
	fmt.Println("\n5. RANGE … - Set of characters")
	fmt.Println("   isDigit('5'):", isDigit('5'))
	fmt.Println("   isLetter('A'):", isLetter('A'))
	fmt.Println("   isLetter('1'):", isLetter('1'))

	// 6. Complete Example - Identifier
	fmt.Println("\n6. COMPLETE EXAMPLE - Identifier")
	fmt.Println("   Valid identifiers:")
	fmt.Println("      isValidIdentifier(\"name\"):", isValidIdentifier("name"))
	fmt.Println("      isValidIdentifier(\"_private\"):", isValidIdentifier("_private"))
	fmt.Println("      isValidIdentifier(\"var123\"):", isValidIdentifier("var123"))
	fmt.Println("   Invalid identifiers:")
	fmt.Println("      isValidIdentifier(\"123var\"):", isValidIdentifier("123var"))
	fmt.Println("      isValidIdentifier(\"my-var\"):", isValidIdentifier("my-var"))

	// 7. Complete Example - Integer
	fmt.Println("\n7. COMPLETE EXAMPLE - Integer Literal")
	fmt.Println("   isValidInteger(\"0\"):", isValidInteger("0"))
	fmt.Println("   isValidInteger(\"123\"):", isValidInteger("123"))
	fmt.Println("   isValidInteger(\"0xFF\"):", isValidInteger("0xFF"))
	fmt.Println("   isValidInteger(\"0xDEADBEEF\"):", isValidInteger("0xDEADBEEF"))

	// 8. Complete Example - For Statement
	fmt.Println("\n8. COMPLETE EXAMPLE - For Statement")
	for1, _ := parseForStatement("for x < 10 { }")
	fmt.Printf("   parseForStatement(\"for x < 10 {{ }}\").Type: %s\n", for1.ConditionType)
	for2, _ := parseForStatement("for i := 0; i < 10; i++ { }")
	fmt.Printf("   parseForStatement(\"for i := 0; i < 10; i++ {{ }}\").Type: %s\n", for2.ConditionType)
	for3, _ := parseForStatement("for { }")
	fmt.Printf("   parseForStatement(\"for {{ }}\").Type: %s\n", for3.ConditionType)

	// 9. Complete Example - Function Call
	fmt.Println("\n9. COMPLETE EXAMPLE - Function Call")
	fc1, _ := parseFunctionCall("fmt.Println()")
	fmt.Printf("   parseFunctionCall(\"fmt.Println()\"): %+v\n", fc1)
	fc2, _ := parseFunctionCall("add(2, 3)")
	fmt.Printf("   parseFunctionCall(\"add(2, 3)\"): %+v\n", fc2)

	fmt.Println("\n" + "="*70)
}
