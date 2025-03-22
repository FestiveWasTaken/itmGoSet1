package main

// Shift letter
//
// Shift a letter right by the given number.
// Wrap the letter around if it reaches the end of the alphabet.
//
// Examples
// shiftLetter("A", 0) -> "A"
// shiftLetter("A", 2) -> "C"
// shiftLetter("Z", 1) -> "A"
// shiftLetter("X", 5) -> "C"
// shiftLetter(" ", _) -> " "
//
// Note: we use the underscore `_` to denote the presence of a value that is present but irrelevant.
//
// Params:
// - letter, a single uppercase English letter, or a space
// - shift, the number by which to shift the letter
//
// Returns:
// - the letter, shifted appropriately, if a letter; otherwise a space.
func shiftLetter(letter string, shift int) string {
	engAlphabet := [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	if letter == " " {
		return " "
	}

	//convert to range to 0-25 range
	num := int(letter[0] - 'A')

	// apply shift and wrap
	shifted := (num + shift) % 26
	if shifted < 0 {
		shifted += 26
	}

	// convert with array as reference
	return string(engAlphabet[shifted])
}

// Caesar cipher
//
// Apply a shift number to a string of uppercase English letters and spaces.
//
// Params:
// - message, a string of uppercase English letters and spaces
// - shift, the number by which to shift the letters
//
// Returns
// - the message, shifted appropriately
func caesarCipher(message string, shift int) string {
	result := ""
	for _, char := range message {
		result += shiftLetter(string(char), shift)
	}
	return result
}

// Vigenere cipher https://www.wikiwand.com/en/articles/Vigen%C3%A8re_cipher
//
// Encrypt a message using a keyphrase instead of a static number.
// Every letter in the message is shifted by the number represented by the respective letter in the key.
// Spaces are ignored.
//
// Example
// vigenereCipher("A C", "KEY") -> "K A"
//
// For example, if the plaintext is attacking tonight and the key is oculorhinolaryngology, then
// the first letter of the plaintext, a, is shifted by 14 positions in the alphabet (because the first letter of the key, o, is the 14th letter of the alphabet, counting from zero), yielding o;
// the second letter, t, is shifted by 2 (because the second letter of the key, c, is the 2nd letter of the alphabet, counting from zero) yielding v;
// the third letter, t, is shifted by 20 (u), yielding n, with wrap-around;
//
// If needed, the keyphrase is extended to match the length of the key.
// If the key is "KEY" and the message is "LONGTEXT", the key will be extended to "KEYKEYKE".
//
// Params:
// - message, a string of uppercase english letters and/or spaces
// - key, a string of uppercase English letters with no spaces. Will not exceed the length of the message.
//
// Returns:
// - the message, shifted appropriately
func vigenereCipher(message string, key string) string {
	result := ""
	keyIndex := 0
	for _, char := range message {
		if char == ' ' {
			result += " "
		} else {
			// Convert key character to shift value (0-25)
			shift := int(key[keyIndex%len(key)] - 'A')
			result += shiftLetter(string(char), shift)
			keyIndex++
		}
	}
	return result
}

// Scytale cipher
//
// Encrypts a message by simulating a scytale cipher.
//
// A scytale is a cylinder around which you can wrap a long strip of
//
//	parchment that contained a string of apparent gibberish. The parchment,
//	when read using the scytale, would reveal a message due to every nth
//	letter now appearing next to each other, revealing a proper message.
//
// This encryption method is obsolete and should never be used to encrypt
//
//	data in production settings.
//
// You may read more about the method here:
//
//	https://en.wikipedia.org/wiki/Scytale
//
// You may follow this algorithm to implement a scytale-style cipher:
//  1. Take a message to be encoded and a "shift" number.
//     For this example, we will use "INFORMATION_AGE" as
//     the message and 3 as the shift.
//  2. Check if the length of the message is a multiple of
//     the shift. If it is not, add additional underscores
//     to the end of the message until it is.
//     In this example, "INFORMATION_AGE" is already a multiple of 3,
//     so we will leave it alone.
//  3. This is the tricky part. Construct the encoded message.
//     For each index i in the encoded message, use the character at the index
//     (i // shift) + (len(message) // shift) * (i % shift) of the raw message.
//     If this number doesn't make sense, you can play around with the cipher at
//     https://dencode.com/en/cipher/scytale to try to understand it.
//  4. Return the encoded message. In this case,
//     the output should be "IMNNA_FTAOIGROE".
//
// Example
// scytaleCipher("INFORMATION_AGE", 3) -> "IMNNA_FTAOIGROE"
// scytaleCipher("INFORMATION_AGE", 4) -> "IRIANMOGFANEOT__"
// scytaleCipher("ALGORITHMS_ARE_IMPORTANT", 8) -> "AOTSRIOALRH_EMRNGIMA_PTT"
//
// Params:
// - message, a string of uppercase English letters and underscores. Underscores represent spaces.
// - shift, a positive integer that does not exceed the length of the message.
//
// Returns:
// - the message, encoded appropriately.
func scytaleCipher(message string, shift int) string {
	// Pad message with underscores if needed
	for len(message)%shift != 0 {
		message += "_"
	}

	result := make([]rune, len(message))
	rows := len(message) / shift

	for i := 0; i < len(message); i++ {
		oldIndex := (i / shift) + (rows * (i % shift))
		result[i] = rune(message[oldIndex])
	}

	return string(result)
}
