package main

import (
	
	"fmt"
)

func main(){
	fmt.Println("#####################To-Do List App#######################")
	var task1="Go out for a walk"
	var taskList = [] string {task1}
	lister(taskList)

}

func lister(taskItems []string){
	fmt.Println("~ List of to-do's:")
	for index, task := range taskItems{
		fmt.Printf("%d. %s \n", index+1,task)
	}
}
func addTask(taskItems []string, newtask string){
	var updatedTaskItems = append(taskItems,newtask)
	lister(updatedTaskItems)
}