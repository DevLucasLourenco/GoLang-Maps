package main

import "fmt"

type informacao struct {
	nome        string
	idade       uint8
	localizacao endereco
}

type endereco struct {
	rua    string
	numero uint16
	cidade string
	estado string
}

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
			"curso_nome": "Engenharia de Software",
			"duração":    "5 anos",
		},
	}

	// fmt.Println(usuario2["curso"]["curso_nome"])
	delete(usuario2, "curso")
	usuario2["signo"] = map[string]string{
		"signo": "Sagitario",
	}

	fmt.Println(usuario2)

	infos()
}

func infos() {

	enderecamento := endereco{rua: "graccho rangel", numero: 553, cidade: "Niterói", estado: "RJ"}
	dado_pessoal := informacao{nome: "lucas", idade: 20, localizacao: enderecamento}
	fmt.Println(dado_pessoal)
}
