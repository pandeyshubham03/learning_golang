package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (eb englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi There!"
}

func (sp spanishBot) getGreeting() string {
	// very custom logic for generating an Spanish greeting
	return "Hola!"
}
