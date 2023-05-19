package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

type Target struct {
	InputFilePath  string // Путь до файла с паролями
	OutputFilePath string // Путь до файла, куда должны записываться результаты
	*Connection
}

func HackServer(ctx context.Context, target *Target) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go requester(ctx, target)
	go responser(ctx, target)
	for {
		select {
		case <-ctx.Done():
			time.Sleep(5 * time.Second)
		}
	}
}

func requester(ctx context.Context, target *Target) {
	passFile, err := os.Open(target.InputFilePath)
	if err != nil {
		fmt.Println("Не удалось открыть файл:", err)
		return
	}
	defer passFile.Close()
	scanner := bufio.NewScanner(passFile)
	for scanner.Scan() {
		go SendRequest(target.Connection, &Request{ctx, scanner.Text()})
	}

}

func responser(ctx context.Context, target *Target) {
	ctx, cancel := context.WithCancel(ctx)
	file, err := os.Create(target.OutputFilePath)
	if err != nil {
		fmt.Println("Не удалось создать файл:", err)
		return
	}
	defer file.Close()
	point := true
	for {
		select {
		case y := <-target.Connection.ResponseConn:
			if point {
				_, err := file.WriteString(fmt.Sprintf("pass: \"%s\"  %t\n", y.Password, y.Pass))
				if err != nil {
					fmt.Println("Ошибка при записи в файл:", err)
					return
				}
			}

			if y.Pass {
				point = false
				time.Sleep(3 * time.Second)
				cancel()
			}

		}
	}
}

func main() {
	requestChan := make(chan *Request)
	responseChan := make(chan *Response)
	defer close(requestChan)
	defer close(responseChan)

	connection := &Connection{
		RequestConn:  requestChan,
		ResponseConn: responseChan,
	}

	target := &Target{
		InputFilePath:  "darkweb2017-top10000.txt",
		OutputFilePath: "output.txt",
		Connection:     connection,
	}
	// Заменить "Password" на один из 10000 паролей
	server := NewVulnerableServer("123456", connection)

	go server.Run()

	// // Пробовать запускать с разными контекстами
	ctx := context.Background()
	HackServer(ctx, target)
}
