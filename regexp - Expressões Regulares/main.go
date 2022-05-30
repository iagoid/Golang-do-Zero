package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Para todos os metodos abaixo
	// caso seja retornado um valor nil NÃO há correspondência.

	// Match informa se a slice de bytes contém qualquer correspondência
	// *Consultas mais complicadas precisam usar o Compile e a interface Regexp completa.
	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`a(b`, []byte(`seafood`))
	fmt.Println(matched, err)

	// MatchString igual a de bytes porém com string
	matched, err = regexp.MatchString(`foo.*`, "seafood")
	fmt.Println(matched, err)

	// Match informa se a slice de byte contém qualquer correspondência
	// MatchString informa se a string contém qualquer correspondência
	re := regexp.MustCompile(`foo.?`)
	fmt.Println(re.Match([]byte(`seafood fool`)))
	fmt.Println(re.Match([]byte(`something else`)))
	re = regexp.MustCompile(`(gopher){2}`)
	fmt.Println(re.MatchString("gopher"))
	fmt.Println(re.MatchString("gophergopher"))
	fmt.Println(re.MatchString("gophergophergopher"))

	// Find retorna um slice contendo o texto da correspondência mais à esquerda
	// FindAll retorna um slice de todas as correspondências
	re = regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.Find([]byte(`seafood fool`)))
	fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), -1))

	// FindAllIndex retorna um slice de indices, do inicio ao fim da busca
	// O segundo argumento define quantos dos items encontrados eu quero que devolva
	// numero -1 define o máximo como o tamanho da palavra
	word := "Promoção"
	content := []byte(word)
	re = regexp.MustCompile(`o.`)
	fmt.Println(re.FindAllIndex(content, 1))

	allIndexes := re.FindAllIndex(content, -1)
	fmt.Println(allIndexes)
	for intGroup := range allIndexes[0] {
		posInitial := allIndexes[intGroup][0]
		posFinal := allIndexes[intGroup][1]
		substring := word[posInitial:posFinal]
		fmt.Println(substring)
	}

	// FindAllString retorna um slice de strings com todas as correspondência
	// FindAllStringIndex retorna um slice de indexes com todas as correspondência
	// O segundo argumento define quantos dos items encontrados eu quero que devolva
	// numero -1 define o máximo como o tamanho da palavra
	re = regexp.MustCompile(`a.`)
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
	fmt.Println(re.FindAllStringIndex("paranormal", -1))
	fmt.Println(re.FindAllStringIndex("paranormal", 2))
	fmt.Println(re.FindAllStringIndex("graal", -1))
	fmt.Println(re.FindAllStringIndex("none", -1))

	// FindAllStringSubmatch retorna um slice da substring encontrada em todas as correspondência
	// FindAllStringSubmatchIndex retorna um slice com indices da substring encontrada em todas as correspondência
	// O segundo argumento define quantos dos items encontrados eu quero que devolva
	// numero -1 define o máximo como o tamanho da palavra
	fmt.Println("\nAllSubmatch")
	re = regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))

	// FindAllSubmatchIndex retorna um slice de todos os indices as correspondências sucessivas
	// FindIndex retorna um slice de dois elementos de inteiros definindo a localização da correspondência mais à esquerda
	re = regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`), -1))
	fmt.Printf("%q\n", re.FindAllSubmatchIndex([]byte(`seafood fool`), -1))
	content = []byte(`
		# comment line
		option1: value1
		option2: value2
	`)
	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	allIndexes = pattern.FindAllSubmatchIndex(content, -1)
	for _, loc := range allIndexes {
		fmt.Println(loc)
		fmt.Println(string(content[loc[0]:loc[1]]))
		fmt.Println(string(content[loc[2]:loc[3]]))
		fmt.Println(string(content[loc[4]:loc[5]]))
	}

	loc := pattern.FindIndex(content)
	fmt.Println(loc)
	fmt.Println(string(content[loc[0]:loc[1]]))

	// FindString retorna uma string contendo o texto da correspondência mais à esquerda
	// Se não houver correspondência, o valor de retorno será uma string vazia,
	// mas também estará vazia se a expressão regular corresponder com sucesso a uma string vazia.
	// *Use FindStringIndex ou FindStringSubmatch se for necessário distinguir esses casos.
	re = regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindString("seafood fool"))
	fmt.Printf("%q\n", re.FindString("meat"))

	// FindStringIndex retorna um slice de dois elementos as posições das stringsa encontradas
	// FindStringSubmatch retorna um slice de strings contendo o texto da correspondência
	// mais à esquerda da expressão regular, se houver, de suas subexpressões
	// FindStringSubmatchIndex retorna um slice de indices contendo o texto da correspondência
	// mais à esquerda da expressão regular, se houver, de suas subexpressões
	// correspondência mais à esquerda da expressão regular.
	re = regexp.MustCompile(`ab?`)
	fmt.Println(re.FindStringIndex("tablett"))
	fmt.Println(re.FindStringIndex("foo") == nil)

	re = regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
	fmt.Printf("%q\n", re.FindStringSubmatchIndex("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatchIndex("-abzc-"))

	// FindSubmatch retorna o texto da correspondência mais à esquerda
	// e as correspondência de sua subexpressõe
	fmt.Println("\n\n============================")
	re = regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindSubmatch([]byte(`seafood fool`)))

	// FindSubmatchIndex retorna os pares de índices que identificam a correspondência mais à esquerda
	re = regexp.MustCompile(`a(x*)b`)
	fmt.Println(re.FindSubmatchIndex([]byte("-ab-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-ab-axb-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-ab-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-foo-")))

	// NumSubexp retorna o número de subexpressões entre parênteses neste Regexp.
	re0 := regexp.MustCompile(`a.`)
	fmt.Printf("%d\n", re0.NumSubexp())
	re = regexp.MustCompile(`(.*)((a)b)(.*)a`)
	fmt.Println(re.NumSubexp())

	// ReplaceAll substitui as correspondências do Regexp pelo texto de substituição
	// Dentro de repl, os sinais $ são interpretados como em Expandir, então, por exemplo,
	// $ 1 representa o texto da primeira subcorrespondência.
	fmt.Println("\n\nReplaceAll")
	re = regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%s\n", re.ReplaceAll([]byte("axxxxxxxxxxb"), []byte("substituido")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("ab-axxb"), []byte("$1")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("ab-axxb"), []byte("$1substituido")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("ab-axxb"), []byte("${1}substituido")))

	// ReplaceAllString as correspondências do Regexp pelo texto de substituição
	// Dentro de repl, os sinais $ são interpretados como em Expandir, então, por exemplo,
	// $ 1 representa o texto da primeira subcorrespondência.
	fmt.Println()
	fmt.Printf("%s\n", re.ReplaceAllString("ab-axxb", "substituido"))
	fmt.Printf("%s\n", re.ReplaceAllString("ab-axxb", "$1"))
	fmt.Printf("%s\n", re.ReplaceAllString("ab-axxb", "$1substituido"))
	fmt.Printf("%s\n", re.ReplaceAllString("ab-axxb", "${1}substituido"))

	// Split os slices em substrings separadas pela expressão
	// e retorna um slice das substrings entre essas correspondências de expressão.
	// O slice retornada por esse método consiste em todas as substrings de não contidas no slice retornado por FindAllString.
	fmt.Println("\nSplit")
	fmt.Printf("%s\n", regexp.MustCompile("a*").Split("aBaaBaCCaDaaaE", 5))

	// SubexpNames retorna os nomes das subexpressões entre parênteses
	// neste Regexp. O nome da primeira subexpressão é names[1], de modo que,
	// o nome de m[i] será SubexpNames()[i].
	// Como o Regexp como um todo não pode ser nomeado, names[0] é sempre a string vazia.
	// A fatia não deve ser modificada.
	fmt.Println("\nNomes no regex")
	re = regexp.MustCompile(`(?P<nome>[a-zA-Z]+) (?P<sobrenome>[a-zA-Z]+)`)

	nome := "Alan Turing"
	if re.MatchString(nome) {
		fmt.Printf("%q\n", re.SubexpNames()) //todos os nomes
		reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
		fmt.Println(reversed)
		fmt.Println(re.ReplaceAllString(nome, reversed))
	}

}
