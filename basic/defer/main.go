package main

//
// defer работает от последнего к первому тоесть это LIFO
// прервать можно с помощью os.Exit()
// все defer привязаны к своей функции. Свое выполнениея они начинают после окончания выполнения
//      функции в которой они обявлены


// startEx3() - простой пример того как работает очередь деферов
// startEx2() - пример с замыканием переменной в дефере
// startEx1() - просто набор деферов которые вызываются в разных функциях
func main() {
	startEx1()
}