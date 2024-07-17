package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Шалом!")
	var Args = os.Args[1:]
	var Config config
	var Handlers = make(map[string]handler)
	switch Args[0] {
	case "configfile":
		if len(Args) == 2 {
			fmt.Println("Вказано шлях до конфіг файлу")
			Config = configfile(Args[1])
		} else {
			fmt.Println("Не вказано шлях до конфіг файлу")
		}
	}
	for _, Handler := range Config.Handlers {
		Handlers[Handler.Description] = Handler
	}
	for _, Binding := range Config.Bindings {
		// fmt.Println("!!!!!!")
		if len(Binding.Address) == 0 {
			panic("Не вказано адресу біндінгу")
		}
		fmt.Println(Binding.Address)
		BindingHandler := Handlers[Binding.Handler]
		if len(BindingHandler.Type) > 0 {
			go bind(Binding, Handlers[Binding.Handler], Binding.Buffer)
		} else {
			panic(fmt.Sprintf("Хендлер %s вказаний для біндінгу %s якісь не дуже. Його не існує.", Binding.Handler, Binding.Description))
		}
	}
	for {
		time.Sleep(1 * time.Hour)
	}
}
