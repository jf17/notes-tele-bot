package notestelebot

	

import (
	//"strconv"
	"time"
)

type Notes struct {
	Create,LastEdit time.Time // время создания заметки и время последнего редактирования
	Note string // заметки
	Tag []string // тэги
}

func GetListNotes() {
     // TODO code 
}


func AddNewNote() {
   // TODO code 
}

func DeleteNote() {
	// TODO code 
 }
 
 
func parseFile() {
	// TODO code 
 }