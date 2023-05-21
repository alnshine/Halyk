package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const PlaneValue = 10
const parallelAirTrafficController = 10

// TODO:: должна быть структура долдна быть потокобезопасная
type runway struct {
	mx     sync.Mutex
	isBusy bool
}

// useRunway испольщования взлетнопосадочной полосы
// в один моомент только один самолет может испольщовать поля данной структуры
func (r *runway) useRunway(plane *Plane, action string) {
	r.mx.Lock()
	r.isBusy = true

	fmt.Println(getCurrentTime(), plane.title, action)

	plane.status = action
	time.Sleep(time.Second * 1)

	r.isBusy = false
	r.mx.Unlock()
}

type Airport struct {
	runway

	// каналы для управления посадками и взлетами
	takeoffCh chan *Plane
	landingCh chan *Plane

	// поля которые помогут вам закрыть аэропорт
	stctx context.Context
	stop  context.CancelFunc

	// после которое говорит об завершении всех дел - программа может умирать
	done  chan struct{}
	close bool
}

// isClose функция для проверки - можно ли взлетать?
func (a *Airport) isClose() bool {
	err := a.stctx.Err()
	if err != nil {
		return true
	}
	return false
}

// NewAirport создание новой аэропорта и запуска его действия в отдельной горутине
func NewAirport() *Airport {
	stctx, stop := context.WithCancel(context.Background())

	a := &Airport{
		landingCh: make(chan *Plane),
		takeoffCh: make(chan *Plane),

		done: make(chan struct{}),

		stctx: stctx,
		stop:  stop,
	}

	// запускаем ассинхроную функцию для функционирования аэропорта
	//TODO:: как обработать ошибку от функции airportProcess?
	go func() {
		err := a.airportProcess()
		if err != nil {
			os.Exit(1)
		}
	}()

	return a
}

// airTrafficController создание воркера для уплавления самолетами
func (a *Airport) airTrafficController(wg *sync.WaitGroup, activePlanesCount *int64) {
	for {
		// TODO:: проверяем есть ли неприпаркованыне самолеты
		if atomic.LoadInt64(activePlanesCount) <= 0 {
			fmt.Println("activePlanesCount=0")
			return
		}
		select {
		case plane, _ := <-a.takeoffCh: // логика взлета самолета
			//TODO:: проверка - можно ли совершать вылеты,
			// если нет - самолет остаётся в статусе parking и более ничего не делает
			// запрет на взлет нужно залогировать

			if a.isClose() {
				fmt.Println("airport zap"+getCurrentTime(), plane.title)
				wg.Done()
				atomic.AddInt64(activePlanesCount, -1)
				continue
			}

			a.useRunway(plane, "takeoff")
			go plane.flying(a) // полетел

		case plane, _ := <-a.landingCh: // логика посадки самолета
			a.useRunway(plane, "landing")

			//TODO:: обслуживаться одновременно могут только 3 самодета.
			go plane.servicing(a) // на сервисе

		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

// airportProcess фнкция обслудивания самолетов - создание диспетчеров
func (a *Airport) airportProcess() error {
	wg := &sync.WaitGroup{}
	activePlanesCount := int64(PlaneValue)

	wg.Add(parallelAirTrafficController)
	for i := 0; i < parallelAirTrafficController; i++ {
		go a.airTrafficController(wg, &activePlanesCount)
	}

	wg.Wait()
	a.done <- struct{}{}

	//TODO:: нужно дождаться завершения всех самолетных дел,
	// после отправить сигнал в метод Close(), что можно закрываться

	return nil
}

// Start запуск работы аэропорта
func (a *Airport) Start() [PlaneValue]*Plane {
	planes := [PlaneValue]*Plane{}

	for i, _ := range planes {
		planes[i] = &Plane{title: i, status: "starting"}
		go func(p *Plane) {
			a.takeoffCh <- p
		}(planes[i])
	}

	return planes
}

// Close остановку работы аэропорта
func (a *Airport) Close(seconds time.Duration) {
	time.Sleep(time.Second * seconds)

	fmt.Printf("Airpost is closing...\n")
	//TODO:: новые самолеты не должны вылетать, остальные должны пойти на посадку в срочном порядке
	// также остановить обслуживание если оно проходит в данный момент
	// вы обязаны дождаться пока все самолеты не закончат все свои дела
	a.stop()
	<-a.done
}
