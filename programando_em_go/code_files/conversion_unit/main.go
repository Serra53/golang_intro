
package main


import (
    "fmt"
    "os"
    "strconv"
)

func main(){
    //os.Args is of type slice
    //containing all the arguments
    //passed by the program

    if len(os.Args) < 3 {
        fmt.Println("Uso: conversor <valores> <unidade>")
        os.Exit(1)
    }

    unidadeOrigem := os.Args[len(os.Args)-1]
    valoresOrigem := os.Args[1: len(os.Args)-1]

    var unidadeDestino string

    if unidadeOrigem == "celsius" {
        unidadeDestino = "fahrenheit"
    } else if unidadeOrigem == "quilometros" {
        unidadeDestino = "milhas"
    } else {
        fmt.Printf("A unidade %v não é reconhecida. Tente celsius ou quilometros.\n", unidadeOrigem)
        os.Exit(1)
    }

    for i, v := range valoresOrigem{
        valorOrigem, err := strconv.ParseFloat(v, 64)
        if err != nil {
            fmt.Printf("O valor %s da posição %d não é um número válido!\n", v,i)
            os.Exit(1)
        }
        var valorDestino float64
        if unidadeOrigem == "celsius"{
            //temp
            valorDestino = valorOrigem * 1.8 + 32
        } else {
            //distancia
            valorDestino = valorOrigem / 1.60934
        }
        //%.2f informa a função que o
        //valor do tipo float deve ser
        //arrendodado para as duas casas
        //decimais
        fmt.Printf("%.2f %s = %.2f %s \n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
    }
    
}
