package main

import "fmt"

// Go uses a sequence of bytes to represent strings. A string is a read-only slice of bytes.
// These bytes are not to be in any particular character encoding. For example, a string containing the character 'a' is represented by the byte 97 in ASCII encoding.
// Several built-in functions are assumed that a string is composed of a sequence of UTF-8 encoded Unicode code points.
// Go source code always written in UTF-8 encoding. Unless you have to use a different encoding, your string literals are always UTF-8 encoded.
// A code point is a numerical value that maps to a specific character in a character set, such as ASCII, Unicode, or UTF-8.
// UTF-8 is an encoding that represents Unicode code points using a variable number of bytes. It is the most common encoding for Unicode characters.

func main() {
	// we can extract a single byte from a string using the index expression
	s := "hello"
	c := s[1]      // c is byte
	fmt.Println(c) // 101 (ASCII code for 'e')
	// when we use the index expression on a string, we get a byte at that position.
	// and each byte is an 8-bit value that represents a single character in the ASCII character set.
	// a byte is an alias for uint8, so it can hold values from 0 to 255.

	// or we can extract a substring from a string using the slice expression
	ss := s[1:3]    // ss is string
	fmt.Println(ss) // el

	// The string above was composed entirely of code points that are one byte long, so everything worked as expected.
	// However, not all code points are 1 byte long. Some code points are two, three, or four bytes long, e.g., emoji, Chinese characters, Japanese characters, etc.
	// While a code point in UTF-8 can be anywhere from 1 to 4 bytes long, a code unit in UTF-8 is always 1 byte long.
	// If we try to extract a single byte from a string that contains multi-byte code points, we will get only the first byte of the code point.
	var s1 = "hello, 😈" // 😈 is a single code point, but it is represented by 4 bytes
	var s2 = s1[5:7]    // got � because the slice expression extracted only the first byte of the code point
	var s3 = s1[:5]     // got hello
	var s4 = s1[6:]     // got 😈 because the slice expression extracted the entire code point
	fmt.Printf("s1: %s\ns2: %s\ns3:%s\ns4: %s\n", s1, s2, s3, s4)

	// Given that string index and slice expressions operate on bytes, not code points, so built-in len function returns the number of bytes in a string, not the number of characters.
	fmt.Println(len("hello 😈")) // 10 (6 bytes for hello_ and 4 bytes for 😈)

	// We also can convert a single run or byte to a string
	var a rune = 'a'       // rune is an alias for int32 and represents a Unicode code point
	fmt.Println(string(a)) // a
	var b byte = 'b'
	fmt.Println(string(b)) // b

	// Do not try to make an int into a string by using string(42) or string(3.14159), because it will return the Unicode character with that code point.
	var i int = 42
	fmt.Println(string(i)) // *

	// A string can be converted back and forth to a slice of bytes or a slice of runes
	str := "Hello, 🌞"
	bs := []byte(str) // convert string to []byte
	fmt.Println(bs)   // [72 101 108 108 111 44 32 240 159 140 158]
	rs := []rune(str) // convert string to []rune
	fmt.Println(rs)   // [72 101 108 108 111 44 32 127774]
	/*
		- 240 159 140 158 is the UTF-8 encoding of the 🌞 emoji. When we convert the string to a slice of bytes, we get the UTF-8 encoding of the string. Each byte represents a single character.
		- 127774 is the Unicode code point of the 🌞 emoji. When we convert the string to a slice of runes, we get the Unicode code points of the string. Each code point represents a single character.
		- So if when we work with strings that contain multi-byte code points, we should use the slice of runes to get the correct code point.
	*/

}
