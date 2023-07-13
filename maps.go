package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Lourenço",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"nome":      "Lucas",
			"sobrenome": "Lourenço",
		},
		"curso": {
			"curso_nome": "engenharia de Software",
			"duração":    "3 anos",
		},
	}

	// fmt.Println(usuario2["curso"]["curso_nome"])
	delete(usuario2, "curso")
	usuario2["signo"] = map[string]string{
		"signo": "Sagitario",
	}

	fmt.Println(usuario2)
}
