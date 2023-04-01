package main

import "fmt"

// Crie uma aplicação em Go que simule uma carteira virtual de Bitcoins.
//     A aplicação deve possuir uma interface base chamada Carteira, que define métodos básicos para gerenciamento de Bitcoins, como enviar, receber e consultar saldo.
//     Crie uma struct chamada Endereco que implementa a interface Carteira e possui as propriedades de um endereço de carteira de Bitcoins, como chave pública e privada.
//     Crie uma função chamada enviarBitcoin que recebe como argumento um valor em Bitcoins e um endereço de destino, e utiliza a interface Carteira para realizar a transação.
// Por fim, crie uma instância de Endereco e utilize-a para realizar uma transação de Bitcoins.

func (c Coin) enviarBitcoin(v float64, e string) string {
	c.balance = c.balance - v
	return "Operação realizada com sucesso."
	//		for _, c := range coins {
	//		if c.Endereco.PublicKey == address{
	//			c.balance = c.balance + valor
	//		}
	//	}
}

func (c Coin) receberBitcoin(v float64, e string) string {
	fmt.Println(v)
	c.balance = c.balance + v
	fmt.Println(c.balance)
	return "Operação realizada com sucesso."
}

func (c Coin) consultarSaldo() string {
	return fmt.Sprint(c.balance)
}

type Carteira interface {
	enviarBitcoin() string
	receberBitcoin() string
	consultarSaldo() float64
}

type Endereco struct {
	PublicKey  string
	PrivateKey string
}

type Coin struct {
	name     string
	balance  float64
	Endereco Endereco
}

// var coins = []Coin{
// 	Coin{"Bitcoin", 0, Endereco{"11111", "9876543210"}},
// 	Coin{"LiteCoin", 0, Endereco{"22222", "9876543210"}},
// }

func main() {

	coin := Coin{"Bitcoin", 0, Endereco{"11111", "9876543210"}}
	// c2 := Coin{"LiteCoin", 0, Endereco{"22222", "9876543210"}}

	for {
		fmt.Println("O que você deseja fazer?")
		fmt.Println("1 - Enviar moedas.")
		fmt.Println("2 - Receber moedas.")
		fmt.Println("3 - Consultar saldo.")
		fmt.Println("4 - Sair.")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			var value float64
			fmt.Println("Qual valor você deseja enviar?")
			fmt.Scan(&value)
			var endereco string
			fmt.Println("Para qual endereço você deseja enviar?")
			fmt.Scan(&endereco)

			// var coin Coin

			// for

			coin.enviarBitcoin(value, endereco)

		case 2:
			var value float64
			fmt.Println("Qual valor você vai receber?")
			fmt.Scan(&value)
			var endereco string
			fmt.Println("Para qual endereço você deseja enviar?")
			fmt.Scan(&endereco)
			coin.receberBitcoin(value, endereco)

		case 3:
			fmt.Println("Saldo - Bitcoin :", coin.consultarSaldo())
		case 4:
			return
		default:
			return

		}
	}
}
