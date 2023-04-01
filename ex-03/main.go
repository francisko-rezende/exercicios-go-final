package main

import "fmt"

type Veiculo struct {
	Marca  string
	Modelo string
}

type Carro struct {
	Veiculo
	NumeroDePortas int
}

type Moto struct {
	Veiculo
	Cilindradas int
}

type Revisao interface {
	fazerRevisao() string
}

func (c Carro) fazerRevisao() string {
	return fmt.Sprint("Revisão do carro agendada. Veículo: ", c.Marca, ", ", c.Modelo, "\n")
}

func (m Moto) fazerRevisao() string {
	return fmt.Sprint("Revisão da moto agendada. Veículo: ", m.Marca, ", ", m.Modelo, "\n")
}

func agendarRevisao(r Revisao) {
	fmt.Printf(r.fazerRevisao())
}

func main() {

	v1 := Veiculo{Marca: "Fiat", Modelo: "Uno"}
	c := Carro{v1, 2}
	v2 := Veiculo{Marca: "Honda", Modelo: "CG"}
	m := Moto{Veiculo: v2, Cilindradas: 125}

	agendarRevisao(c)
	agendarRevisao(m)
}
