package controllers

import (
	"FirstProyectWebEngineering/models"
	"FirstProyectWebEngineering/services"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TaskResponse struct {
	ID           primitive.ObjectID `json:"ID"`
	Name         string              `json:"Name"`
	Description  string              `json:"Description"`
	EmployeeName string              `json:"EmployeeName"`
	ProjectName  string              `json:"ProjectName"`
	StartDate    time.Time           `json:"StartDate"`
	EndDate      time.Time           `json:"EndDate"`
	IsCompleted   bool               `json:"IsCompleted"` 
}

type TaskDuration struct {
	ID             primitive.ObjectID `json:"ID"`
	Name           string              `json:"Name"`
	Description    string              `json:"Description"`
	EmployeeName   string              `json:"EmployeeName"`
	ProjectName    string              `json:"ProjectName"`
	StartDate      time.Time           `json:"StartDate"`
	EndDate        time.Time           `json:"EndDate"`
	IsCompleted     bool               `json:"IsCompleted"`
	DurationDays   int                 `json:"DurationDays"` 
	OverdueDays    int                 `json:"OverdueDays"`   
}


type TaskSummaryResponse struct {
	ID           primitive.ObjectID `json:"id"`
	Name         string              `json:"name"`
	IsCompleted   bool               `json:"is_completed"`
}

type TaskController struct {
	Service          services.TaskService
	EmployeeService  services.EmployeeService
	ProyectService   services.ProyectService
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := c.EmployeeService.GetEmployeeByID(task.ID_Employee)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	_, err = c.ProyectService.GetProyectByID(task.ID_Project)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	result, err := c.Service.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var taskResponses []TaskResponse
	for _, task := range tasks {
		employee, _ := c.EmployeeService.GetEmployeeByID(task.ID_Employee)
		project, _ := c.ProyectService.GetProyectByID(task.ID_Project)

		taskResponse := TaskResponse{
			ID:           task.ID,
			Name:         task.Name,
			Description:  task.Description,
			EmployeeName: employee.Name,
			ProjectName:  project.Name,
			StartDate:    task.StartDate, 
			EndDate:      task.EndDate,   
			IsCompleted:  task.IsCompleted, 
		}
		taskResponses = append(taskResponses, taskResponse)
	}

	json.NewEncoder(w).Encode(taskResponses)
}

// New method for getting a summary of tasks
func (c *TaskController) GetTaskSummary(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var taskSummaries []TaskSummaryResponse
	for _, task := range tasks {
		taskSummary := TaskSummaryResponse{
			ID:           task.ID,
			Name:         task.Name,
			IsCompleted:  task.IsCompleted,
		}
		taskSummaries = append(taskSummaries, taskSummary)
	}

	json.NewEncoder(w).Encode(taskSummaries)
}

func (c *TaskController) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := c.Service.GetTaskByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_,  err = c.EmployeeService.GetEmployeeByID(task.ID_Employee)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	_, err = c.ProyectService.GetProyectByID(task.ID_Project)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}


	json.NewEncoder(w).Encode(task)
}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = c.EmployeeService.GetEmployeeByID(task.ID_Employee)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	_, err = c.ProyectService.GetProyectByID(task.ID_Project)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	result, err := c.Service.UpdateTask(id, &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	result, err := c.Service.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

type DateRangeRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (c *TaskController) GetTasksBetweenDates(w http.ResponseWriter, r *http.Request) {
	var dateRange DateRangeRequest

	if err := json.NewDecoder(r.Body).Decode(&dateRange); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse("2006-01-02", dateRange.StartDate)
	if err != nil {
		http.Error(w, "Invalid start date", http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse("2006-01-02", dateRange.EndDate)

	if err != nil {
		http.Error(w, "Invalid end date", http.StatusBadRequest)
		return
	}

	fmt.Printf("Starting tasks between %s and %s", startDate, endDate)

	tasks, err := c.Service.GetTasksBetweenDates(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("tasks %s", tasks)

	var filteredTasks []TaskResponse
	for _, task := range tasks {
		if task.IsCompleted{ 
			employee, _ := c.EmployeeService.GetEmployeeByID(task.ID_Employee)
			project, _ := c.ProyectService.GetProyectByID(task.ID_Project)

			taskResponse := TaskResponse{
				ID:           task.ID,
				Name:         task.Name,
				Description:  task.Description,
				EmployeeName: employee.Name,
				ProjectName:  project.Name,
				StartDate:    task.StartDate,
				EndDate:      task.EndDate,
				IsCompleted:  task.IsCompleted,
			}
			filteredTasks = append(filteredTasks, taskResponse)
		}
	}

	json.NewEncoder(w).Encode(filteredTasks)
}







