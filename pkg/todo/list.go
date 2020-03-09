package todo

import (
	"fmt"
	"todo/pkg/infra"
	"todo/pkg/utils"
)

func List() {
	// todoText := utils.ApplyStyle("bold", "yellow", "todo")
	// doingText := utils.ApplyStyle("bold", "red", "doing")
	// doneText := utils.ApplyStyle("bold", "blue", "done")
	utils.ClearScreen()
	listTodo()
	listDoing()
	listDone()

}

func listTodo() {
	todoText := utils.ApplyStyle("bold", "yellow", "todo")

	listTodo, err := infra.GetRedis().LRange("todo", 0, 1000000).Result()
	if err != nil {
		panic(err)
	}
	if len(listTodo) > 0 {
		fmt.Println(todoText)
		for _, v := range listTodo {
			fmt.Println("🧩 ", v)
		}
	}
}
func listDoing() {
	doingText := utils.ApplyStyle("bold", "red", "doing")

	listDoing, err := infra.GetRedis().LRange("doing", 0, 1000000).Result()
	if err != nil {
		panic(err)
	}
	if len(listDoing) > 0 {
		fmt.Println("\n", doingText)
		for _, v := range listDoing {
			fmt.Println("♨️  ", v)
		}
	}
}
func listDone() {
	doneText := utils.ApplyStyle("bold", "blue", "done")

	listDone, err := infra.GetRedis().LRange("done", 0, 1000000).Result()
	if err != nil {
		panic(err)
	}
	if len(listDone) > 0 {
		fmt.Println("\n", doneText)
		for _, v := range listDone {
			fmt.Println("✅ ", v)
		}
	}
}
