package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	configMutex sync.RWMutex
	configPath  string
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		configPath = ".dir-init/config-v2.yaml"
	} else {
		configPath = filepath.Join(home, ".dir-init", "config-v2.yaml")
	}
}

// GetConfigPath returns the path to the config file
func GetConfigPath() string {
	return configPath
}

// LoadConfig loads the user configuration from ~/.dir-init/config.yaml
func LoadConfig() (*Config, error) {
	configMutex.RLock()
	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configMutex.RUnlock()
		// Initialize with defaults if it doesn't exist
		if err := InitConfig(); err != nil {
			return nil, fmt.Errorf("failed to initialize config: %w", err)
		}
		configMutex.RLock()
	}
	defer configMutex.RUnlock()

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Initialize maps if nil
	if config.Frameworks == nil {
		config.Frameworks = make(map[string][]Framework)
	}
	if config.Categories == nil {
		config.Categories = make(map[string][]string)
	}

	return &config, nil
}

// SaveConfig saves the configuration to ~/.dir-init/config.yaml
func SaveConfig(config *Config) error {
	configMutex.Lock()
	defer configMutex.Unlock()

	// Create directory if it doesn't exist
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Marshal to YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Add header comment
	header := "# dir-init Custom Collections\n# Auto-generated and manually editable\n\n"
	finalData := append([]byte(header), data...)

	// Write to temp file first (atomic write)
	tempPath := configPath + ".tmp"
	if err := os.WriteFile(tempPath, finalData, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	// Rename temp file to actual config file
	if err := os.Rename(tempPath, configPath); err != nil {
		os.Remove(tempPath) // Clean up temp file
		return fmt.Errorf("failed to save config file: %w", err)
	}

	return nil
}

// SaveTechStack adds a tech stack to the config
func SaveTechStack(code, description string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Check if already exists
	for _, ts := range config.TechStacks {
		if ts.Code == code {
			return fmt.Errorf("tech stack '%s' already exists", code)
		}
	}

	// Add new tech stack
	config.TechStacks = append(config.TechStacks, TechStack{
		Code:        code,
		Description: description,
	})

	return SaveConfig(config)
}

// SaveFramework adds a framework to the config
func SaveFramework(techStack, code, description string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Initialize frameworks map for tech stack if needed
	if config.Frameworks[techStack] == nil {
		config.Frameworks[techStack] = []Framework{}
	}

	// Check if already exists
	for _, fw := range config.Frameworks[techStack] {
		if fw.Code == code {
			return fmt.Errorf("framework '%s' already exists for tech stack '%s'", code, techStack)
		}
	}

	// Add new framework
	config.Frameworks[techStack] = append(config.Frameworks[techStack], Framework{
		Code:        code,
		Description: description,
	})

	return SaveConfig(config)
}

// SaveCategoryWord adds a word to a category in the config
func SaveCategoryWord(category, word string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Initialize category if needed
	if config.Categories[category] == nil {
		config.Categories[category] = []string{}
	}

	// Check if already exists
	for _, w := range config.Categories[category] {
		if w == word {
			return fmt.Errorf("word '%s' already exists in category '%s'", word, category)
		}
	}

	// Add new word
	config.Categories[category] = append(config.Categories[category], word)

	return SaveConfig(config)
}

// InitConfig creates an example config file
func InitConfig() error {
	// Check if config already exists
	if _, err := os.Stat(configPath); err == nil {
		return nil // Already exists, nothing to do
	}

	// Create default config with all values
	defaultConfig := &Config{
		TechStacks: []TechStack{
			{Code: "rct", Description: "React"},
			{Code: "vue", Description: "Vue.js"},
			{Code: "ng", Description: "Angular"},
			{Code: "svelte", Description: "Svelte"},
			{Code: "nxt", Description: "Next.js"},
			{Code: "nuxt", Description: "Nuxt.js"},
			{Code: "sol", Description: "Solid"},
			{Code: "qwik", Description: "Qwik"},
			{Code: "pre", Description: "Preact"},
			{Code: "grt", Description: "Gatsby"},
			{Code: "astro", Description: "Astro"},
			{Code: "remix", Description: "Remix"},
			{Code: "html", Description: "HTML/CSS/JS"},
			{Code: "node", Description: "Node.js"},
			{Code: "py", Description: "Python"},
			{Code: "go", Description: "Go"},
			{Code: "java", Description: "Java"},
			{Code: "ruby", Description: "Ruby"},
			{Code: "php", Description: "PHP"},
			{Code: "rust", Description: "Rust"},
			{Code: "csharp", Description: "C#"},
			{Code: "deno", Description: "Deno"},
			{Code: "bun", Description: "Bun"},
			{Code: "spring", Description: "Spring Boot"},
			{Code: "django", Description: "Django"},
			{Code: "flask", Description: "Flask"},
			{Code: "fastapi", Description: "FastAPI"},
			{Code: "express", Description: "Express"},
			{Code: "nest", Description: "NestJS"},
			{Code: "rails", Description: "Rails"},
			{Code: "laravel", Description: "Laravel"},
			{Code: "none", Description: "None"},
		},
		Categories: map[string][]string{
			"food": {
				"pizza", "burger", "taco", "pasta", "sushi", "donut", "sandwich", "salad",
				"soup", "steak", "chicken", "fish", "rice", "noodles", "curry", "stew",
				"bbq", "kebab", "wrap", "panini", "quesadilla", "burrito", "nachos",
				"lasagna", "risotto", "paella", "couscous", "tabbouleh", "hummus",
				"cake", "cupcake", "muffin", "cookie", "brownie", "pie", "icecream",
				"gelato", "sorbet", "pudding", "flan", "tiramisu", "cheesecake",
				"croissant", "danish", "eclair", "profiterole", "macaron", "meringue",
				"lollipop", "candy", "chocolate", "fudge", "toffee", "brittle",
				"coffee", "tea", "espresso", "latte", "cappuccino", "americano", "mocha",
				"juice", "smoothie", "milkshake", "soda", "water", "lemonade", "icedtea",
				"hotchocolate", "chai", "matcha", "boba", "cocktail", "mocktail",
				"chips", "popcorn", "pretzels", "nuts", "seeds", "crackers", "bread",
				"cheese", "olives", "pickles", "dips", "salsa", "guacamole", "bruschetta",
				"canapes", "springrolls", "wings", "onionrings", "fries", "mozzarella",
				"calamari", "samosas", "pakoras",
			},
			"animals": {
				"penguin", "koala", "dolphin", "eagle", "tiger", "panda", "turtle", "rabbit",
				"fox", "wolf", "bear", "lion", "otter", "meerkat", "sloth", "hippo",
				"giraffe", "zebra", "elephant", "rhino", "monkey", "gorilla", "orangutan",
				"chimpanzee", "lemur", "kangaroo", "wallaby", "wombat", "platypus",
				"armadillo", "hedgehog", "porcupine", "ferret", "mongoose", "badger",
				"raccoon", "skunk", "opossum", "coyote", "lynx", "bobcat", "cougar",
				"parrot", "cockatoo", "macaw", "toucan", "cockatiel", "budgie",
				"canary", "finch", "sparrow", "robin", "bluejay", "cardinal",
				"falcon", "hawk", "owl", "vulture", "condor", "flamingo",
				"pelican", "seagull", "pigeon", "dove", "swan", "goose", "duck",
				"peacock", "turkey", "quail", "pheasant", "partridge",
				"shark", "whale", "porpoise", "orca", "beluga", "narwhal",
				"octopus", "squid", "cuttlefish", "jellyfish", "starfish", "seahorse",
				"clownfish", "angelfish", "betta", "goldfish", "koi", "tuna", "salmon",
				"trout", "bass", "catfish", "swordfish", "marlin", "halibut",
				"butterfly", "moth", "dragonfly", "damselfly", "beetle", "ladybug",
				"ant", "bee", "wasp", "hornet", "grasshopper", "cricket", "prayingmantis",
				"spider", "scorpion", "centipede", "millipede", "earthworm", "leech",
				"slug", "snail", "crayfish", "lobster", "crab", "shrimp", "prawn",
			},
			"pop": {
				"ninja", "samurai", "wizard", "knight", "viking", "pirate", "astronaut",
				"robot", "superhero", "detective", "warrior", "mage", "sorcerer", "paladin",
				"ranger", "cleric", "druid", "assassin", "barbarian", "monk", "bard",
				"healer", "summoner", "elementalist", "chronomancer", "pyromancer",
				"cryomancer", "geomancer", "aeromancer", "necromancer",
				"musician", "artist", "painter", "sculptor", "writer", "author", "poet",
				"dancer", "actor", "director", "producer", "composer", "conductor", "singer",
				"guitarist", "pianist", "drummer", "violinist", "cellist", "flutist",
				"photographer", "filmmaker", "videographer", "editor", "designer",
				"architect", "chef", "baker", "mixologist", "bartender",
				"pharaoh", "emperor", "king", "queen", "prince", "princess", "duke",
				"duchess", "baron", "baroness", "squire", "serf", "peasant",
				"gladiator", "centurion", "legion", "spartan", "athenian", "roman",
				"norse", "celtic", "greek", "egyptian", "maya", "aztec",
				"influencer", "blogger", "vlogger", "streamer", "gamer", "esports",
				"contentcreator", "socialmedia", "tiktok", "instagram", "youtube",
				"brandambassador", "spokesperson", "representative",
				"ambassador", "diplomat", "negotiator", "mediator", "arbitrator",
			},
			"silly": {
				"potato", "banana", "unicorn", "noodle", "pickle", "muffin", "cupcake",
				"cookie", "marshmallow", "popcorn", "cucumber", "broccoli", "carrot",
				"tomato", "pepper", "onion", "garlic", "ginger", "lettuce", "spinach",
				"mushroom", "avocado", "papaya", "mango", "kiwi", "pomegranate",
				"pineapple", "coconut", "watermelon", "honeydew", "cantaloupe", "fig",
				"date", "prune", "apricot", "peach", "plum", "cherry", "berry",
				"rubberduck", "sockpuppet", "paperclip", "stapler", "highlighter",
				"gluestick", "scissors", "ruler", "protractor", "compass", "eraser",
				"calculator", "abacus", "typewriter", "telegraph", "telephone", "radio",
				"television", "computer", "keyboard", "mouse", "monitor", "printer",
				"scanner", "camera", "microphone", "speaker", "headphones", "earbuds",
				"happy", "sad", "angry", "excited", "nervous", "confused", "surprised",
				"shocked", "amazed", "bored", "tired", "sleepy", "hungry", "thirsty",
				"curious", "playful", "goofy", "weird", "strange", "bizarre", "odd",
				"peculiar", "quirky", "crazy", "wild", "silly", "funny",
				"hilarious", "comical", "absurd", "ridiculous", "ludicrous",
				"butterfly", "dragonfly", "firefly", "lightningbug", "ladybug",
				"jellybean", "poprocks", "cottoncandy", "licorice", "taffy",
				"gummybear", "chocolatechip", "peanutbutter", "strawberry", "blueberry",
				"raspberry", "blackberry", "cranberry", "gooseberry", "elderberry",
			},
			"dev": {
				"github", "gitlab", "bitbucket", "mercurial", "svn", "cvs", "perforce",
				"stash", "source", "repository", "repo", "branch", "trunk", "tag",
				"commit", "push", "pull", "merge", "rebase", "cherry-pick", "fork",
				"clone", "remote", "origin", "upstream", "downstream", "head",
				"master", "main", "develop", "feature", "release", "hotfix",
				"aws", "gcp", "azure", "heroku", "digitalocean", "linode", "vultr",
				"cloudflare", "vercel", "netlify", "railway", "flyio", "render",
				"lambda", "ec2", "gke", "aks", "eks", "compute", "serverless",
				"container", "vm", "instance", "machine", "node", "pod", "cluster",
				"react", "vue", "angular", "svelte", "nextjs", "nuxt", "gatsby",
				"express", "koa", "fastify", "flask", "django", "rails", "laravel",
				"spring", "symfony", "aspnet", "bun", "bunchee",
				"webpack", "vite", "rollup", "parcel", "esbuild", "snowpack",
				"babel", "typescript", "coffeescript", "jsx", "tsx", "sass", "less",
				"docker", "kubernetes", "helm", "istio", "linkerd", "consul", "etcd",
				"terraform", "ansible", "puppet", "chef", "saltstack", "fabric",
				"jenkins", "travis", "circleci", "githubactions", "gitlabci", "bamboo",
				"gradle", "maven", "npm", "yarn", "pnpm", "pip", "composer", "gem",
				"brew", "apt", "yum", "dnf", "pacman", "emerge", "pkg",
				"jest", "cypress", "selenium", "puppeteer", "playwright", "cucumber",
				"mocha", "jasmine", "qunit", "vitest", "ava", "tape", "chai", "sinon",
				"eslint", "prettier", "stylelint", "sonarqube", "codeclimate", "coveralls",
				"codecov", "dependabot", "renovate", "snyk", "githubsecurity",
			},
		},
	}

	return SaveConfig(defaultConfig)
}
