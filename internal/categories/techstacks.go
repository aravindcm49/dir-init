package categories

// TechStacks combines tech stack with language in minified format
var TechStackWords = []string{
	// Frontend JavaScript
	"fejs",
	// Frontend TypeScript
	"fets",
	// Backend Python
	"bepy",
	// Backend Node.js
	"bens",
	// Backend Java
	"bejv",
	// Full Stack JavaScript
	"fsjs",
	// DevOps Go
	"dogo",
	// Mobile Kotlin
	"mokt",
	// Mobile Swift
	"mosw",
	// Data Science Python
	"dspyt",
	// Machine Learning Python
	"mlpy",
	// Backend Ruby
	"berb",
	// Backend C++
	"becp",
	// Full Stack TypeScript
	"fsts",
	// DevOps Python
	"dopy",
	// Data Science R
	"dsr",
	// Machine Learning R
	"mlr",
	// Backend Rust
	"bers",
	// Frontend Dart
	"fedt",
	// Mobile Flutter
	"mofl",
	// DevOps Rust
	"dors",
	// Backend PHP
	"beph",
	// Backend C#
	"becs",
	// Full Stack Python
	"fspy",
	// DevOps Node.js
	"donj",
	// DevOps Java
	"dojv",
	// Data Science Scala
	"dssc",
	// Machine Learning Julia
	"mljl",
}

// TechStackDescriptions maps tech stack codes to human-readable descriptions
func TechStackDescriptions() map[string]string {
	return map[string]string{
		"fejs": "Frontend JavaScript",
		"fets": "Frontend TypeScript",
		"bepy": "Backend Python",
		"bens": "Backend Node.js",
		"bejv": "Backend Java",
		"fsjs": "Full Stack JavaScript",
		"dogo": "DevOps Go",
		"mokt": "Mobile Kotlin",
		"mosw": "Mobile Swift",
		"dspyt": "Data Science Python",
		"mlpy": "Machine Learning Python",
		"berb": "Backend Ruby",
		"becp": "Backend C++",
		"fsts": "Full Stack TypeScript",
		"dopy": "DevOps Python",
		"dsr": "Data Science R",
		"mlr": "Machine Learning R",
		"bers": "Backend Rust",
		"fedt": "Frontend Dart",
		"mofl": "Mobile Flutter",
		"dors": "DevOps Rust",
		"beph": "Backend PHP",
		"becs": "Backend C#",
		"fspy": "Full Stack Python",
		"donj": "DevOps Node.js",
		"dojv": "DevOps Java",
		"dssc": "Data Science Scala",
		"mljl": "Machine Learning Julia",
	}
}