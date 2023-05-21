package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Plane struct {
	title  int
	status string
}

// flying функция для полета самолета, в конце она отправляет самолет на посадку
func (p *Plane) flying(a *Airport) {
	p.status = "fly"

	r := rand.Intn(10)
	if r < 3 {
		r = 3
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r)*time.Second)
	defer cancel()

	if a.isClose() {
		fmt.Println("fly zap"+getCurrentTime(), p.title)
		cancel()
	}

	<-ctx.Done()
	fmt.Println("priletel"+getCurrentTime(), p.title)
	a.landingCh <- p
	//TODO:: логика полета.
	// Полет либо должен закончится по таймауту, либо если аэропорт скажет садить - мы закрываемся

	//TODO:: самолет нужно отправить на посадку
}

// servicing функция обслуживания самолета, в конце она отправляет самолет обратно на взлет
func (p *Plane) servicing(a *Airport) {

	if !a.isClose() {
		fmt.Println(getCurrentTime(), p.title)
		p.status = "on service"
		r := rand.Intn(3)
		if r < 1 {
			r = 1
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r)*time.Second)
		defer cancel()

		<-ctx.Done()
		fmt.Println("end service"+getCurrentTime(), p.title)

	}

	//TODO:: логика обслуживания самолета.
	// Обслуивание либо должено закончится по таймауту, либо если аэропорт скажет заканчивай - мы закрываемся

	p.status = "parking"

	//TODO:: самолет нужно отправить на попытку взлета
	a.takeoffCh <- p
}
