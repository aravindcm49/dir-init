package categories

// TechToFrameworkMapping maps tech stack codes to available frameworks
func TechToFrameworkMapping() map[string][]string {
	return map[string][]string{
		// Frontend JavaScript
		"fejs": {
			"rct", "vue", "ng", "svelte", "nxt", "nuxt", "grt", "pre", "sol", "qwik",
			"flt", "rnm", "nativ", "ionic", "cap", "expo", "mon", "ons",
		},
		// Frontend TypeScript
		"fets": {
			"rct", "vue", "ng", "svelte", "nxt", "nuxt", "grt", "pre", "sol", "qwik",
			"flt", "rnm", "nativ", "ionic", "cap", "expo", "mon", "ons",
		},
		// Backend Python
		"bepy": {
			"dja", "flk", "fas", "py", "tur", "chalc", "zap",
			"spr", "mic", "quark", "drop",
		},
		// Backend Node.js
		"bens": {
			"exp", "fast", "koa", "hel", "nest",
			"type", "seq", "pr", "dr", "knx",
		},
		// Backend Java
		"bejv": {
			"spr", "jvax", "mic", "quark", "vertx", "drop", "weld", "arq",
		},
		// Full Stack JavaScript
		"fsjs": {
			"rct", "vue", "ng", "svelte", "nxt", "nuxt", "grt", "pre", "sol", "qwik",
			"exp", "fast", "koa", "hel", "nest",
			"web", "http", "api", "serv", "arch", "bld", "tool",
		},
		// DevOps Go
		"dogo": {
			"gor", "htt", "eve", "fib", "gof", "rel", "beego",
			"kt", "fin", "twil", "play", "akka", "lag", "glue",
		},
		// Mobile Kotlin
		"mokt": {
			"flt", "cap", "exp", "mon", "ons",
		},
		// Mobile Swift
		"mosw": {
			"flt", "cap", "exp", "mon", "ons",
		},
		// Data Science Python
		"dspyt": {
			"tf", "pt", "ker", "jup", "dg", "spk", "h2o", "xgb",
		},
		// Machine Learning Python
		"mlpy": {
			"tf", "pt", "ker", "jup", "dg", "spk", "h2o", "xgb",
		},
		// Backend Ruby
		"berb": {
			"rls", "sin", "grap", "pad", "han",
		},
		// Backend C++
		"becp": {
			"web", "http", "api", "serv", "arch", "bld", "tool",
		},
		// Full Stack TypeScript
		"fsts": {
			"rct", "vue", "ng", "svelte", "nxt", "nuxt", "grt", "pre", "sol", "qwik",
			"exp", "fast", "koa", "hel", "nest",
			"type", "seq", "pr", "dr", "knx",
			"web", "http", "api", "serv", "arch", "bld", "tool",
		},
		// DevOps Python
		"dopy": {
			"fas", "py", "tur", "chalc", "zap",
			"web", "http", "api", "serv", "arch", "bld", "tool",
		},
		// Data Science R
		"dsr": {
			"jup", "dg", "spk", "h2o", "xgb",
		},
		// Machine Learning R
		"mlr": {
			"jup", "dg", "spk", "h2o", "xgb",
		},
		// Backend Rust
		"bers": {
			"web", "http", "api", "serv", "arch", "bld", "tool",
		},
		// Frontend Dart
		"fedt": {
			"flt", "rnm", "nativ", "ionic", "cap", "expo", "mon", "ons",
		},
		// Mobile Flutter
		"mofl": {
			"flt", "cap", "exp", "mon", "ons",
		},
		// DevOps Rust
		"dors": {
			"web", "http", "api", "serv", "arch", "bld", "tool",
		},
		// Backend PHP
		"beph": {
			"lav", "sym", "slim", "fat", "cake", "code", "phal", "zend",
		},
		// Backend C#
		"becs": {
			"asp", "blaz", "nfx", "serv", "http", "web",
		},
		// Full Stack Python
		"fspy": {
			"dja", "flk", "fas", "py", "tur", "chalc", "zap",
			"web", "http", "api", "serv", "arch", "bld", "tool",
		},
		// DevOps Node.js
		"donj": {
			"exp", "fast", "koa", "hel", "nest",
			"type", "seq", "pr", "dr", "knx",
			"web", "http", "api", "serv", "arch", "bld", "tool",
		},
		// DevOps Java
		"dojv": {
			"spr", "jvax", "mic", "quark", "vertx", "drop", "weld", "arq",
			"web", "http", "api", "serv", "arch", "bld", "tool",
		},
		// Data Science Scala
		"dssc": {
			"spk", "akka", "lag", "jup", "dg",
		},
		// Machine Learning Julia
		"mljl": {
			"jup", "dg", "spk", "h2o", "xgb",
		},
	}
}

// GetFrameworksForTech returns available frameworks for a given tech stack code
func GetFrameworksForTech(techStack string) []string {
	frameworks, exists := TechToFrameworkMapping()[techStack]
	if !exists {
		return []string{} // Return empty array for unknown tech stacks
	}
	return frameworks
}

// IsCompatible checks if a framework is compatible with a given tech stack
func IsCompatible(techStack, framework string) bool {
	frameworks := GetFrameworksForTech(techStack)
	for f := range frameworks {
		if frameworks[f] == framework {
			return true
		}
	}
	return false
}