package repository

import (
	"gin-micro-tm/config"
	"gin-micro-tm/model"
	"log"
)

// CreateItem ...
func CreateItem(task model.Task) (model.Task, error) {
	db := config.DB

	var err error

	crt, err := db.Prepare("insert into tasks (title, body) values (?, ?)")
	if err != nil {
		log.Panic(err)
		return task, err
	}

	res, err := crt.Exec(task.Title, task.Body)
	if err != nil {
		//log.Panic(err)
		return task, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		log.Panic(err)
		return task, err
	}

	task.ID = int64(rowID)

	// find item by id
	resval, err := GetItemByID(task.ID)
	if err != nil {
		log.Panic(err)
		return task, err
	}

	return resval, nil
}

// GetItemByID ...
func GetItemByID(id int64) (model.Task, error) {
	db := config.DB

	var item model.Task

	result, err := db.Query("select id, title, body from tasks where id = ?", id)
	if err != nil {
		// print stack trace
		log.Println("Error query item: " + err.Error())
		return item, err
	}

	for result.Next() {
		err := result.Scan(&item.ID, &item.Title, &item.Body)
		if err != nil {
			return item, err
		}
	}

	return item, nil
}

// GetItemAll ...
func GetItemAll() ([]model.Task, error) {
	db := config.DB
	var task model.Task
	var tasks []model.Task
	rows, err := db.Query("select id, title, body from tasks")
	if err != nil {
		log.Println("Error query get all items: " + err.Error())
		return tasks, err
	}
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.Title, &task.Body); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
