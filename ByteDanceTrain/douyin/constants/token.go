package constants

const (
	lowerLetter              = "abcdefghijklmnopqrstuvwxyz"
	upperLetter              = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digit                    = "0123456789"
	TokenCharacterDictionary = digit + lowerLetter + upperLetter
	TokenDictionaryLength    = len(TokenCharacterDictionary)

	TokenLength = 32
)
