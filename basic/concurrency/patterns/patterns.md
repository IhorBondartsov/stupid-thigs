- fan-in

Описание: Групировка каналов для передачи всей информации через один канал. 


- fan-out
Описание: передача из одного канала в несколько
- Pipelines

Описание: конвеер из несольких рутин которые передают поочердно какую либо информацию и обрабатывают ее.

- pub-sub
```go
type Subscriber interface{
	Receive() chan<- Packet
}

type Publisher interface{
    Subscribe()chan<- Subscriber
    Unsubscriebe() chan<- Subscriber
}
```
##### Timing out [здесь](https://github.com/IhorBondartsov/stupid-things/tree/master/basic/concurrency/channels/timing%20out/main.go )
 
 moving on








