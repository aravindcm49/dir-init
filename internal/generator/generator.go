package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/aravindcm49/dir-init/internal/categories"
)

type SuffixType string

const (
	SuffixAlpha      SuffixType = "alpha"
	SuffixNumeric    SuffixType = "numeric"
	SuffixMixed      SuffixType = "mixed"
	SuffixTimestamp  SuffixType = "timestamp"
)

type Config struct {
	Category        string
	SuffixType      SuffixType
	SuffixLength    int
	IncludeNumbers  bool
	UseEmojis       bool
	Count           int
	Seed            int64
}

type Generator struct {
	config Config
	rand   *rand.Rand
}

func NewGenerator(config Config) *Generator {
	r := rand.New(rand.NewSource(config.Seed))
	if config.Seed == 0 {
		r.Seed(time.Now().UnixNano())
	}

	return &Generator{
		config: config,
		rand:   r,
	}
}

func (g *Generator) Generate() []string {
	if g.config.Count <= 0 {
		g.config.Count = 1
	}

	names := make([]string, 0, g.config.Count)

	for i := 0; i < g.config.Count; i++ {
		category := g.selectCategory()
		prefix := g.selectWordFromCategory(category)
		suffix := g.generateSuffix()
		name := fmt.Sprintf("%s%s", prefix, suffix)
		names = append(names, name)
	}

	return names
}

func (g *Generator) selectCategory() string {
	if g.config.Category == "all" || g.config.Category == "" {
		categories := []string{"tech", "food", "animals", "pop", "silly", "dev"}
		return categories[g.rand.Intn(len(categories))]
	}
	return g.config.Category
}

func (g *Generator) selectWordFromCategory(category string) string {
	var words []string

	switch category {
	case "tech":
		words = categories.TechWords
	case "food":
		words = categories.FoodWords
	case "animals":
		words = categories.AnimalWords
	case "pop":
		words = categories.PopWords
	case "silly":
		words = categories.SillyWords
	case "dev":
		words = categories.DevWords
	default:
		words = categories.TechWords
	}

	if len(words) == 0 {
		return "folder"
	}

	return words[g.rand.Intn(len(words))]
}

func (g *Generator) generateSuffix() string {
	switch g.config.SuffixType {
	case SuffixAlpha:
		return g.generateAlphaSuffix()
	case SuffixNumeric:
		return g.generateNumericSuffix()
	case SuffixMixed:
		return g.generateMixedSuffix()
	case SuffixTimestamp:
		return g.generateTimestampSuffix()
	default:
		return g.generateMixedSuffix()
	}
}

func (g *Generator) generateAlphaSuffix() string {
	length := g.config.SuffixLength
	if length < 3 {
		length = 3
	} else if length > 8 {
		length = 8
	}

	chars := "abcdefghijklmnopqrstuvwxyz"
	var suffix strings.Builder

	for i := 0; i < length; i++ {
		suffix.WriteByte(chars[g.rand.Intn(len(chars))])
	}

	return "-" + suffix.String()
}

func (g *Generator) generateNumericSuffix() string {
	length := g.config.SuffixLength
	if length < 1 {
		length = 1
	} else if length > 6 {
		length = 6
	}

	max := intPow(10, length) - 1
	min := intPow(10, length-1)

	if length == 1 {
		min = 0
	}

	num := g.rand.Intn(max-min+1) + min
	return fmt.Sprintf("-%d", num)
}

func (g *Generator) generateMixedSuffix() string {
	length := g.config.SuffixLength
	if length < 3 {
		length = 3
	} else if length > 8 {
		length = 8
	}

	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	var suffix strings.Builder

	for i := 0; i < length; i++ {
		suffix.WriteByte(chars[g.rand.Intn(len(chars))])
	}

	return "-" + suffix.String()
}

func (g *Generator) generateTimestampSuffix() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("-%d", timestamp%100000000)
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// DefaultConfig returns a default configuration for the generator
func DefaultConfig() Config {
	return Config{
		Category:       "all",
		SuffixType:     SuffixMixed,
		SuffixLength:   4,
		IncludeNumbers: true,
		UseEmojis:      false,
		Count:          1,
		Seed:           0,
	}
}