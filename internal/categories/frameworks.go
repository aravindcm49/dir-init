package categories

// FrameworkWords contains minified framework names
var FrameworkWords = []string{
	// Frontend/Full Stack Frameworks
	"rct",    // React
	"vue",    // Vue
	"ng",     // Angular
	"svelte", // Svelte
	"nxt",    // Next.js
	"nuxt",   // Nuxt.js
	"grt",    // Gatsby
	"pre",    // Preact
	"sol",    // Solid
	"qwik",   // Qwik

	// Backend Frameworks
	"exp",    // Express
	"fast",   // FastAPI
	"koa",    // Koa
	"hel",    // Hapi
	"nest",   // NestJS
	"type",   // TypeORM
	"seq",    // Sequelize
	"pr",     // Prisma
	"dr",     // Drizzle
	"knx",    // Knex

	// Python Frameworks
	"dja",    // Django
	"flk",    // Flask
	"fas",    // FastAPI (already listed but also here for Python focus)
	"py",     // Pyramid
	"tur",    // Tornado
	"chalc",  // Chalice
	"zap",    // Zappa

	// Java Frameworks
	"spr",    // Spring
	"jvax",   // Jakarta EE
	"mic",    // Micronaut
	"quark",  // Quarkus
	"vertx",  // Vert.x
	"drop",   // Dropwizard
	"weld",   // Weld
	"arq",    // Arquillian

	// Go Frameworks
	"gor",    // Gin
	"htt",    // HTTPRouter
	"chir",   // Chi React
	"eve",    // Echo
	"fib",    // Fiber
	"gof",    // Go-Framework
	"rel",    // Revel
	"beego",  // Beego

	// Ruby Frameworks
	"rls",    // Rails
	"sin",    // Sinatra
	"grap",   // Grape
	"pad",    // Padrino
	"han",    // Hanami

	// PHP Frameworks
	"lav",    // Laravel
	"sym",    // Symfony
	"slim",   // Slim
	"fat",    // Fat-Free
	"cake",   // CakePHP
	"code",   // CodeIgniter
	"phal",   // Phalcon
	"zend",   // Zend

	// Mobile Frameworks
	"flt",    // Flutter
	"rnm",    // React Native
	"nativ",  // NativeScript
	"ionic",  // Ionic
	"cap",    // Capacitor
	"expo",   // Expo
	"mon",    // Monaca
	"ons",    // Onsen UI

	// DevOps/Infra Frameworks
	"kt",     // Ktor
	"finc",   // Finatra
	"twil",   // Twirl
	"play",   // Play Framework
	"akka",   // Akka
	"lag",    // Lagom
	"glue",   // Gluecodium

	// Data Science/Machine Learning
	"tf",     // TensorFlow
	"pt",     // PyTorch
	"ker",    // Keras
	"jup",    // Jupyter
	"dg",     // Dask
	"spk",    // Spark
	"h2o",    // H2O
	"xgb",    // XGBoost

	// Generic/Web Frameworks
	"web",    // Generic web
	"http",   // HTTP server
	"api",    // API framework
	"serv",   // Server framework
	"arch",   // Architecture framework
	"bld",    // Build framework
	"tool",   // Tool framework
}

// FrameworkDescriptions maps framework codes to human-readable descriptions
func FrameworkDescriptions() map[string]string {
	return map[string]string{
		"rct":    "React",
		"vue":    "Vue.js",
		"ng":     "Angular",
		"svelte": "Svelte",
		"nxt":    "Next.js",
		"nuxt":   "Nuxt.js",
		"grt":    "Gatsby",
		"pre":    "Preact",
		"sol":    "Solid",
		"qwik":   "Qwik",
		"exp":    "Express",
		"fast":   "FastAPI",
		"koa":    "Koa",
		"hel":    "Hapi",
		"nest":   "NestJS",
		"type":   "TypeORM",
		"seq":    "Sequelize",
		"pr":     "Prisma",
		"dr":     "Drizzle",
		"knx":    "Knex",
		"dja":    "Django",
		"flk":    "Flask",
		"fas":    "FastAPI",
		"py":     "Pyramid",
		"tur":    "Tornado",
		"chalc":  "Chalice",
		"zap":    "Zappa",
		"spr":    "Spring",
		"jvax":   "Jakarta EE",
		"mic":    "Micronaut",
		"quark":  "Quarkus",
		"vertx":  "Vert.x",
		"drop":   "Dropwizard",
		"weld":   "Weld",
		"arq":    "Arquillian",
		"gor":    "Gin",
		"htt":    "HTTPRouter",
		"chir":   "Chi",
		"eve":    "Echo",
		"fib":    "Fiber",
		"gof":    "Go-Framework",
		"rel":    "Revel",
		"beego":  "Beego",
		"rls":    "Rails",
		"sin":    "Sinatra",
		"grap":   "Grape",
		"pad":    "Padrino",
		"han":    "Hanami",
		"lav":    "Laravel",
		"sym":    "Symfony",
		"slim":   "Slim",
		"fat":    "Fat-Free",
		"cake":   "CakePHP",
		"code":   "CodeIgniter",
		"phal":   "Phalcon",
		"zend":   "Zend",
		"flt":    "Flutter",
		"rnm":    "React Native",
		"nativ":  "NativeScript",
		"ionic":  "Ionic",
		"cap":    "Capacitor",
		"expo":   "Expo",
		"mon":    "Monaca",
		"ons":    "Onsen UI",
		"kt":     "Ktor",
		"finc":   "Finatra",
		"twil":   "Twirl",
		"play":   "Play Framework",
		"akka":   "Akka",
		"lag":    "Lagom",
		"glue":   "Gluecodium",
		"tf":     "TensorFlow",
		"pt":     "PyTorch",
		"ker":    "Keras",
		"jup":    "Jupyter",
		"dg":     "Dask",
		"spk":    "Spark",
		"h2o":    "H2O",
		"xgb":    "XGBoost",
		"web":    "Web",
		"http":   "HTTP Server",
		"api":    "API",
		"serv":   "Server",
		"arch":   "Architecture",
		"bld":    "Build",
		"tool":   "Tool",
	}
}