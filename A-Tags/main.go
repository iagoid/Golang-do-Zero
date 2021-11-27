package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	ID        string `Minha tag`
	AccountID string `chave1:"Valor1" chave2:"valor2"`
}

func main() {
	o := reflect.TypeOf(Order{})
	fmt.Println(o.Name())

	fmt.Println("Nome da segunda tag ", o.Field(1).Name)
	fmt.Println("Tipo da segunda tag ", o.Field(1).Type)
	fmt.Println("Numero de Campos ", o.NumField())

	id, _ := o.FieldByName("ID")
	fmt.Println(id.Tag)

	accountID, _ := o.FieldByName("AccountID")
	chave1, _ := accountID.Tag.Lookup("chave1")
	fmt.Println(chave1)

}
