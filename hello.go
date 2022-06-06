package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringTimes = 3
const delay = 5

func main() {

	showIntroduction()

	for {
		showMenu()
		command := getCommand()

		switch command {
		case 1:
			startMonitor()
		case 2:
			showLogs()
		case 0:
			exitProgram(0)
		default:
			exitProgram(-1)
		}
	}
}

func showIntroduction() {
	name := "Ronaldo"
	version := 1.1
	fmt.Println("Olá, sr(a).", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)
	fmt.Println()
	return command
}

func exitProgram(code int) {
	if code == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço este comando")
	}
	os.Exit(code)
}

func startMonitor() {
	fmt.Println("Monitorando...")
	sites := readSitesFromFile()

	for i := 0; i < monitoringTimes; i++ {
		for j, site := range sites {
			fmt.Println("Testando site", j, "-", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}

	fmt.Println()
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		saveLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		saveLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	sites = strings.Split(string(file), "\n")
	return sites
}

func saveLog(site string, isOnline bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("[02/01/2006 15:04:05]") + " - " + site + " - online: " + strconv.FormatBool(isOnline) + "\n")
	file.Close()
}

func showLogs() {
	fmt.Println("Exibindo Logs...")
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
