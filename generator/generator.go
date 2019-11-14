package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	// LengthWeak weak length password
	LengthWeak = 6

	// LengthOK ok length password
	LengthOK = 12

	// LengthStrong strong length password
	LengthStrong = 24

	// LengthVeryStrong very strong length password
	LengthVeryStrong = 36

	// DefaultLetterSet is the letter set that is defaulted to - just the
	// alphabet
	DefaultLetterSet = "abcdefghijklmnopqrstuvwxyz"

	// DefaultLetterAmbiguousSet are letters which are removed from the
	// chosen character set if removing similar characters
	DefaultLetterAmbiguousSet = "ijlo"

	// DefaultNumberSet the default symbol set if character set hasn't been
	// selected
	DefaultNumberSet = "0123456789"

	// DefaultNumberAmbiguousSet are the numbers which are removed from the
	// chosen character set if removing similar characters
	DefaultNumberAmbiguousSet = "01"

	// DefaultSymbolSet the default symbol set if character set hasn't been
	// selected
	DefaultSymbolSet = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"

	// DefaultSymbolAmbiguousSet are the symbols which are removed from the
	// chosen character set if removing ambiguous characters
	DefaultSymbolAmbiguousSet = "<>[](){}:;'/|\\,"
)

var (
	// DefaultConfig is the default configuration, defaults to:
	//    - length = 24
	//    - Includes symbols, numbers, lowercase and uppercase letters.
	//    - Excludes similar and ambiguous characters
	DefaultConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true,
		IncludeNumbers:             true,
		IncludeLowercaseLetters:    true,
		IncludeUppercaseLetters:    true,
		ExcludeSimilarCharacters:   true,
		ExcludeAmbiguousCharacters: true,
	}

	// ErrConfigIsEmpty is the error if the config given is empty
	ErrConfigIsEmpty = errors.New("config is empty")
)

// Generator is what generates the password
type Generator struct {
	*Config
}

// Config is the config struct to hold the settings about
// what type of password to generate
type Config struct {
	// Length is the length of password to generate
	Length int

	// CharacterSet is the setting to manually set the
	// character set
	CharacterSet string

	// IncludeSymbols is the setting to include symbols in
	// the character set
	// i.e. !"£*
	IncludeSymbols bool

	// IncludeNumbers is the setting to include number in
	// the character set
	// i.e. 1234
	IncludeNumbers bool

	// IncludeLowercaseLetters is the setting to include
	// lowercase letters in the character set
	// i.e. abcde
	IncludeLowercaseLetters bool

	// IncludeUppercaseLetters is the setting to include
	// uppercase letters in the character set
	// i.e. ABCD
	IncludeUppercaseLetters bool

	// ExcludeSimilarCharacters is the setting to exclude
	// characters that look the same in the character set
	// i.e. i1jIo0
	ExcludeSimilarCharacters bool

	// ExcludeAmbiguousCharacters is the setting to exclude
	// characters that can be hard to remember or symbols
	// that are rarely used
	// i.e. <>{}[]()/|\`
	ExcludeAmbiguousCharacters bool
}

// New returns a new generator
func New(config *Config) (*Generator, error) {
	if config == nil {
		config = &DefaultConfig
	}

	if !config.IncludeSymbols &&
		!config.IncludeUppercaseLetters &&
		!config.IncludeLowercaseLetters &&
		!config.IncludeNumbers &&
		config.CharacterSet == "" {
		return nil, ErrConfigIsEmpty
	}

	if config.Length == 0 {
		config.Length = LengthStrong
	}

	if config.CharacterSet == "" {
		config.CharacterSet = buildCharacterSet(config)
	}

	return &Generator{Config: config}, nil
}

func buildCharacterSet(config *Config) string {
	var characterSet string
	if config.IncludeLowercaseLetters {
		characterSet += DefaultLetterSet
		if config.ExcludeSimilarCharacters {
			characterSet = removeCharacters(characterSet, DefaultLetterAmbiguousSet)
		}
	}

	if config.IncludeUppercaseLetters {
		characterSet += strings.ToUpper(DefaultLetterSet)
		if config.ExcludeSimilarCharacters {
			characterSet = removeCharacters(characterSet, strings.ToUpper(DefaultLetterAmbiguousSet))
		}
	}

	if config.IncludeNumbers {
		characterSet += DefaultNumberSet
		if config.ExcludeSimilarCharacters {
			characterSet = removeCharacters(characterSet, DefaultNumberAmbiguousSet)
		}
	}

	if config.IncludeSymbols {
		characterSet += DefaultSymbolSet
		if config.ExcludeAmbiguousCharacters {
			characterSet = removeCharacters(characterSet, DefaultSymbolAmbiguousSet)
		}
	}

	return characterSet
}

func removeCharacters(str, characters string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(characters, r) {
			return r
		}
		return -1
	}, str)
}

// NewWithDefault returns a new generator with the default
// config
func NewWithDefault() (*Generator, error) {
	return New(&DefaultConfig)
}

// Generate generates one password with length set in the
// config
func (g Generator) Generate() (*string, error) {
	var generated string
	characterSet := strings.Split(g.Config.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))

	for i := 0; i < g.Config.Length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}

// GenerateMany generates multiple passwords with length set
// in the config
func (g Generator) GenerateMany(amount int) ([]string, error) {
	var generated []string
	for i := 0; i < amount; i++ {
		str, err := g.Generate()
		if err != nil {
			return nil, err
		}

		generated = append(generated, *str)
	}
	return generated, nil
}

// GenerateWithLength generate one password with set length
func (g Generator) GenerateWithLength(length int) (*string, error) {
	var generated string
	characterSet := strings.Split(g.Config.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))
	for i := 0; i < length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}

// GenerateManyWithLength generates multiple passwords with set length
func (g Generator) GenerateManyWithLength(amount, length int) ([]string, error) {
	var generated []string
	for i := 0; i < amount; i++ {
		str, err := g.GenerateWithLength(length)
		if err != nil {
			return nil, err
		}
		generated = append(generated, *str)
	}
	return generated, nil
}
