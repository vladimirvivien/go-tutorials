package hellolib

var greetings = map[string]string{
	"English": "Hello, World!",
	"French":  "Salut Monde",
	"Chinese": "世界您好",
	"Klingon": "qo' vIvan",
	"Hindi":   "हैलो वर्ल्ड",
	"Korean":  "안녕하세요",
	"Russian": "привет мир",
	"Swahili": "Wapendwa Dunia",
	"Spanish": "Hola Mundo",
	"Turkish": "Merhaba Dünya",
	"Kreyol": "Alo tout moun!"
}

// Hello returns "Hello World" greeting in one or more in specified langs
func Hello(langs ...string) []string {
	if len(langs) == 0 {
		return []string{greetings["English"]}
	}

	var results []string
	for _, lang := range langs {
		if greeting, ok := greetings[lang]; ok {
			results = append(results, greeting)
		}
	}

	return results
}
