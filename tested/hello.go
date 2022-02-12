package tested

const hello = "Hello, "
const kannadaHello = " Avare,"
const kannada = "Kannada"
const world = "World"
const tulu = "Tulu"
const tuluHello = " Mere,"

func Hello(name, lang string) string {

	if name == "" {
		name = world
	}

	return addGreetings(name, lang)
}

func addGreetings(name, lang string) (result string) {

	switch lang {
	case kannada:
		result = name + kannadaHello
	case tulu:
		result = name + tuluHello
	default:
		result = hello + name
	}

	return
}
