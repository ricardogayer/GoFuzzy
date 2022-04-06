package main

import (
	"fmt"

	fuzzy "github.com/paul-mannino/go-fuzzywuzzy"
)

func main() {

	fmt.Println("Go Fuzzy Search - Lookup")
	fmt.Println("========================")
	fmt.Println("https://github.com/paul-mannino/go-fuzzywuzzy")

	departamentos := []string{"Marketing", "Financeiro", "Recursos Humanos", "Operações", "Tecnologia da Informação"}
	departamento := "Mkt"

	encontrado, err := fuzzy.ExtractOne(departamento, departamentos)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Departamento: ", encontrado.Match)
	fmt.Println("Score: ", encontrado.Score)

}
