package main

import "fmt"

func fibonacci (n int ) int { //fibonacci recursivo retorna un valor
		if(n == 0){
			return(0)
		}else if(n==1){
			return(1)
		}else{
			return(fibonacci(n-1) + fibonacci(n-2))
		}
}

func fibonacciSucesion(n int) { //para imprimir toda la sucesion
	for i:=0; i<n; i++{
		num := fibonacci(i)
		fmt.Print(num, ", ")
	}
}

 func variadic( args ...int) int { //retorna el numero mas grande
	var mayor int
	for i:=0; i<len(args); i++ { 
		if(i==0){
			mayor= args[i]
		}else if (args[i]>mayor){
			mayor = args[i]
		}
	}
	return mayor
 }



func main () {

	generadorImpares := func() func() uint { //Genera impares
		i := uint(1)  
		return func() uint {
			var impar = i
			i += 2
			return impar
		}
	}
	
	
	intercambiar := func (a, b *int) { //intercambia valores con punteros
		aux := *a
		*a = *b
		*b = aux
	}

	var op uint = 1
	for op != 0{
		fmt.Println("")
		fmt.Println("1.- Fibonacci")
		fmt.Println("2.- Número más grande")
		fmt.Println("3.- Generador de impares")
		fmt.Println("4.- Intercambiar enteros (punteros)")
		fmt.Println("0.- Salir")

		fmt.Scan(&op)

		switch {
		case op == 1:
			var n int
			fmt.Println("Hasta que posición quiere la sucesión de Fibonacci?")
			fmt.Scan(&n)
			fibonacciSucesion(n)

		case op == 2:
			var n uint64;
			fmt.Println("Cuántos números?")
			fmt.Scan(&n);
			s:= make([]int, 0, n);
			var e int;
			for i:=0; i<int(n); i++ {	
				fmt.Scan(&e)
				s = append(s, e)
			}
			mayor := variadic(s...)
			fmt.Println("Mas grande:", mayor)

		case op == 3:
			fmt.Println("Primeros 20 impares:")
			sigImpar := generadorImpares()
			for i:=1; i<20; i++{
				fmt.Println(sigImpar())
			}

		case op == 4:
			var a int
			var b int
			fmt.Println("Escriba el número a:")
			fmt.Scan(&a)
			fmt.Println("Escriba el número b:")
			fmt.Scan(&b)
			intercambiar(&a, &b)
			fmt.Println("Intercambiados:")
			fmt.Println("Número a:", a)
			fmt.Println("Número b:", b)
			
		case op == 0:
			fmt.Println("Saliendo...")

		default:
			fmt.Println("Opción inválida")
		
		}
	}  
	
}