package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/aravindcm49/dir-init/internal/categories"
	"github.com/aravindcm49/dir-init/internal/utils"
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

// GenerateSingleName generates a single directory name with the specified category and suffix type
func (g *Generator) GenerateSingleName(category string, suffixType SuffixType, length int) (string, error) {
	// Create a temporary config for this single generation
	tempConfig := Config{
		Category:      category,
		SuffixType:    suffixType,
		SuffixLength:  length,
		IncludeNumbers: true,
		UseEmojis:     false,
		Count:         1,
		Seed:          g.rand.Int63(),
	}

	// Create a new generator with this config
	tempGenerator := &Generator{
		config: tempConfig,
		rand:   g.rand,
	}

	// Generate the name
	category = tempGenerator.selectCategory()
	prefix := tempGenerator.selectWordFromCategory(category)
	suffix := tempGenerator.generateSuffix()
	name := fmt.Sprintf("%s%s", prefix, suffix)

	// Validate the name
	if !utils.IsValidDirectoryName(name) {
		return "", fmt.Errorf("generated name '%s' is not valid for filesystem", name)
	}

	return name, nil
}

// CreateDirectory generates and creates a single directory
func (g *Generator) CreateDirectory(category string, suffixType SuffixType, length int) error {
	name, err := g.GenerateSingleName(category, suffixType, length)
	if err != nil {
		return err
	}

	err = utils.CreateDirectory(name)
	if err != nil {
		return fmt.Errorf("failed to create directory '%s': %v", name, err)
	}

	return nil
}

// CreateDirectories generates and creates multiple directories
func (g *Generator) CreateDirectories(category string, suffixType SuffixType, length, count int) ([]string, error) {
	if count <= 0 {
		return nil, fmt.Errorf("count must be positive")
	}

	if count > 20 {
		return nil, fmt.Errorf("cannot create more than 20 directories at once")
	}

	names := make([]string, 0, count)
	createdNames := make([]string, 0, count)

	for i := 0; i < count; i++ {
		name, err := g.GenerateSingleName(category, suffixType, length)
		if err != nil {
			continue
		}

		err = utils.CreateDirectory(name)
		if err != nil {
			continue
		}

		names = append(names, name)
		createdNames = append(createdNames, name)
	}

	if len(createdNames) == 0 {
		return nil, fmt.Errorf("failed to create any directories")
	}

	return createdNames, nil
}

// GenerateEnhancedName generates names with the new enhanced format: {techstack}-{framework}-{category}-{suffix}
func (g *Generator) GenerateEnhancedName(techStack, framework, category string, suffixType SuffixType, length int) (string, error) {
	// Generate prefix part: {techstack}-{framework}
	prefix := fmt.Sprintf("%s-%s", techStack, framework)

	// Generate category word
	categoryWord := g.selectWordFromCategory(category)

	// Generate suffix
	suffix := g.generateSuffixWithConfig(suffixType, length)

	// Combine all parts: {techstack}-{framework}-{categoryword}-{suffix}
	name := fmt.Sprintf("%s-%s%s", prefix, categoryWord, suffix)

	// Validate the name
	if !utils.IsValidDirectoryName(name) {
		return "", fmt.Errorf("generated name '%s' is not valid for filesystem", name)
	}

	return name, nil
}

// generateSuffixWithConfig generates suffix with specific configuration
func (g *Generator) generateSuffixWithConfig(suffixType SuffixType, length int) string {
	switch suffixType {
	case SuffixAlpha:
		return g.generateAlphaSuffixWithLength(length)
	case SuffixNumeric:
		return g.generateNumericSuffixWithLength(length)
	case SuffixMixed:
		return g.generateMixedSuffixWithLength(length)
	case SuffixTimestamp:
		return g.generateTimestampSuffix()
	default:
		return g.generateMixedSuffixWithLength(length)
	}
}

// generateAlphaSuffixWithLength generates alphabetic suffix with specific length
func (g *Generator) generateAlphaSuffixWithLength(length int) string {
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

// generateNumericSuffixWithLength generates numeric suffix with specific length
func (g *Generator) generateNumericSuffixWithLength(length int) string {
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

// generateMixedSuffixWithLength generates mixed suffix with specific length
func (g *Generator) generateMixedSuffixWithLength(length int) string {
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