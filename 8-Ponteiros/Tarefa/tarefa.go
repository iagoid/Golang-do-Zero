package main

type Tarefa struct {
	titulo  string
	proxima *Tarefa
}

func main() {
	t1 := Tarefa{titulo: "Comprar pão"}
	t2 := Tarefa{titulo: "Comprar leite"}
	t3 := Tarefa{titulo: "Comprar ovos"}
	t4 := Tarefa{titulo: "Comprar Carne"}
	t5 := Tarefa{titulo: "Dormir cedo"}

	// Cada tarefa aponta para outra tarefa
	// Assim quando quero remove-la(ext3) é so fazer o valor anterior(t2)
	// apontar para t4 e retirar o valor apontado por t3
	t1.proxima = &t2
	t2.proxima = &t3
	t3.proxima = &t4
	t4.proxima = &t5

}
