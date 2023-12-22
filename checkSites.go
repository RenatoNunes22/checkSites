package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	apiURL       = "https://seu-servico-whatsapp-api.com/send-message"
	seuNumero    = "seu_numero"
	numeroDestino = "numero_destino"
)

type Mensagem struct {
	NumeroDestino string `json:"numero_destino"`
	Texto         string `json:"texto"`
}

const delay = 1

func main() {
	
	intro()
	displayMenu()

	command := commadRead()
	
	switch command {
		case 1: 
		for {
			startMonitoring()
			time.Sleep(delay * time.Minute)
			}
			case 2:
				fmt.Println("Exibindo Logs...")
				readLog()
			case 0:
				fmt.Println("Saindo do Programa...")
				os.Exit(0)
			default:
				fmt.Println("Comando inválido!")
				os.Exit(-1)
		}
	
	
}

func intro() {
	version := 1.1
	
	fmt.Println("Este programa está na versão:", version)
}

func commadRead() int {
	var command int
	fmt.Scan(&command)

	return command
}

func displayMenu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func startMonitoring() {
	fmt.Println("Monitorando...") 
	sites := readFileSites()
	
	for _, site := range sites {
		checkSite(site)
	}
}

func checkSite(site string) {

	client := &http.Client{
		Timeout: 5 * time.Second, 
	}

	resp, err := client.Get(site)


	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Timeout ao acessar o site:", site)
			registerLog(site, false)
		} else {
			fmt.Println("Erro ao acessar o site:", err)
		}
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readFileSites() []string {

	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	read := bufio.NewReader(file)

	for {
		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Ocorreu um erro:", err)
		} 
	}
	file.Close()
	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + fmt.Sprint(status) + "\n")

	file.Close()
}

func readLog() {

file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file)) 

}
