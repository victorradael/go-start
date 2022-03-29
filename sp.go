package main

import (
	"fmt"
)

func main() {
	showNames()

	// intro()
	// showMenu()
	// command := readTerminal()

	// switch command {
	// case 1:
	// 	startMonitoring()
	// case 2:
	// 	fmt.Print("Show Logs ... \n")
	// case 0:
	// 	fmt.Print("Exit \n")
	// 	os.Exit(0)
	// default:
	// 	fmt.Print("Command Not Found")
	// 	os.Exit(-1)
	// }

}

func intro() {
	var name string = "Victor"
	version := 1.1

	fmt.Println("Olá, Mr.", name)
	fmt.Println("Program Version", version)
}

func showMenu() {
	fmt.Println("1 -  Iniciar o Monitoramento")
	fmt.Println("2 -  Exibir Logs")
	fmt.Println("0 -  Sair")
}

func readTerminal() int {
	var command int

	// fmt.Scanf("%d", &command)
	fmt.Scan(&command)

	return command
}

// func startMonitoring() {
// 	fmt.Print("Start Monitoring! \n")
// 	var sites [4]string
// 	sites[0] = "https://radaeldeveloper.vercel.app/"
// 	sites[1] = "https://www.nasa.gov/"
// 	sites[2] = "https://www.alura.com.br/"

// 	response, error := http.Get(siteTwo)

// 	fmt.Println(error)

// 	if response.StatusCode == 200 {
// 		fmt.Println(siteTwo, "Online")
// 	} else {
// 		fmt.Println(siteTwo, "Down", response.StatusCode)
// 	}
// }

func showNames() {
	names := []string{}
	names = append(names, "Victor")
	fmt.Println(names)
	fmt.Println("Tamanho", len(names))
	fmt.Println("Capacidade", cap(names))
	names = append(names, "Gabriel")
	fmt.Println(names)
	fmt.Println("Tamanho", len(names))
	fmt.Println("Capacidade", cap(names))
	names = append(names, "Radael")
	fmt.Println(names)
	fmt.Println("Tamanho", len(names))
	fmt.Println("Capacidade", cap(names))
	names = append(names, "Alves")
	fmt.Println(names)
	fmt.Println("Tamanho", len(names))
	fmt.Println("Capacidade", cap(names))
	names = append(names, "Kdt")
	fmt.Println(names)
	fmt.Println("Tamanho", len(names))
	fmt.Println("Capacidade", cap(names))
}

// A Capacidade do Slice dobra quando a capacidade é superada.
