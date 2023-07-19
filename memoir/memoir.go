package memoir

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	defaultProjectName = "default"
	dailyDirectory     = "daily"
	dateFormatFileName = "20060102"
	dateFormat         = "2006/01/02"
)

type Memoir struct {
	Path           string
	CurrentProject string
	TasksFile      string
	Today          string
	DailyTasks     []Task
	// Projects       []string
}

type Task struct {
	Title string
	Done  bool
}

// Initialises a new memoir strucutre in the user's home folder
func Init() (*Memoir, error) {
	// Set dir as .memoir in user's  home directory
	dir, err := getMemoirPath()
	if err != nil {
		return nil, err
	}

	// Create directory
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0775)
		if err != nil {
			return nil, errors.New("failed to initilise memoir")
		}

		// Create tasks directory
		err = os.MkdirAll(filepath.Join(dir, dailyDirectory), 0775)
		if err != nil {
			return nil, errors.New("failed to initilise memoir")
		}

		// Create projects directory
		// err = os.MkdirAll(fmt.Sprintf("%s/projects", dir), 0775)
		// if err != nil {
		// 	return nil, errors.New("failed to initilise memoir")
		// }
	}

	// Create .current file
	file := filepath.Join(dir, ".current")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		err = os.WriteFile(file, []byte("default"), 0644)
		if err != nil {
			return nil, errors.New("failed to create .current file")
		}
	}

	// Create default
	m := &Memoir{
		Path: dir,
	}

	// Create default notebook
	// m.CreateProject(defaultProjectName)

	return m, nil

}

// Load memoir from the user's home directory
func Load() (*Memoir, error) {
	today := time.Now()
	return LoadFromDate(today)
}

// load memoir with specific date
func LoadFromDate(date time.Time) (*Memoir, error) {
	// Get memoir path
	dir, err := getMemoirPath()
	if err != nil {
		return nil, err
	}

	// Create memoir structure
	m := &Memoir{
		Path:      dir,
		TasksFile: filepath.Join(dir, dailyDirectory, fmt.Sprintf("%s.daily.md", date.Format(dateFormatFileName))),
		Today:     date.Format(dateFormat),
	}
	m.GetDailyTasks()

	return m, nil
}

// getMemoirPath returns the path to the memoir directory
func getMemoirPath() (string, error) {
	// Get user's home directory
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("failed to get user's home directory")
	}

	// Set memoir path
	dir = filepath.Join(dir, ".memoir")

	// Return memoir path
	return dir, nil
}

// GetDailyTasks gets the daily tasks for the current day
func (m *Memoir) GetDailyTasks() error {
	// Initialise tasks slice
	tasks := []Task{}

	// Open file
	file, err := os.Open(m.TasksFile)
	if err != nil {
		return errors.New("failed to open file")
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, parseTask(scanner.Text()))
	}

	// Set tasks
	m.DailyTasks = tasks

	return nil
}

func parseTask(strTask string) Task {
	done := true
	parts := strings.Split(strTask, "[x]")

	if len(parts) == 1 {
		parts = strings.Split(strTask, "[ ]")
		done = false
	}

	task := Task{
		Title: strings.Trim(parts[1], " "),
		Done:  done,
	}
	return task
}

// AddTask adds a new task to the current day
func (m *Memoir) AddTask(title string) error {
	task := Task{
		Title: title,
	}

	// Add task to the list
	m.DailyTasks = append(m.DailyTasks, task)

	m.SaveTasks()

	return nil
}

// DeleteTask deletes a task from the daily list
func (m *Memoir) DeleteTask(id int) error {
	if id < 0 || id > len(m.DailyTasks) {
		return errors.New("the task id doesn't exist")
	}
	// Rmove task from the list
	m.DailyTasks = append(m.DailyTasks[:id-1], m.DailyTasks[id:]...)

	m.SaveTasks()

	return nil
}

// SaveTasks saves the tasks to file
func (m *Memoir) SaveTasks() error {
	// Create file if it doesn't exist
	dailyFile := filepath.Join(m.TasksFile)
	file, err := os.OpenFile(dailyFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.New("error adding task")
	}

	for _, task := range m.DailyTasks {
		done := " "
		if task.Done {
			done = "x"
		}
		file.WriteString(fmt.Sprintf("- [%s] %s\n", done, task.Title))
	}

	defer file.Close()

	return nil
}

func (m *Memoir) GetAllDailies() []string {
	entries, _ := os.ReadDir(filepath.Join(m.Path, dailyDirectory))

	var days []string
	for _, e := range entries {
		parts := strings.Split(e.Name(), ".daily.md")
		days = append(days, parts[0])
	}

	// Reverse the order
	for i, j := 0, len(days)-1; i < j; i, j = i+1, j-1 {
		days[i], days[j] = days[j], days[i]
	}

	return days
}

// Get all projects
// func (m *Memoir) GetProjects() []string {
// 	// Get all directories within the memoir path
// 	dirs, _ := os.ReadDir(m.Path)

// 	// Initialise projects slice
// 	projects := []string{}

// 	// Add all directories to the projects slice
// 	for _, dir := range dirs {
// 		if dir.IsDir() {
// 			projects = append(projects, dir.Name())
// 		}
// 	}

// 	// Return projects
// 	return projects
// }

// Get notes for the current project
// func (m *Memoir) GetProjectNotes() []string {
// 	// Set path to the current project
// 	path := filepath.Join(m.Path, m.GetCurrentProject())

// 	// Get all directories within the project path
// 	dirs, _ := os.ReadDir(path)

// 	// Initialise notes slice
// 	notes := []string{}

// 	// Add diretory contents to the notes slice
// 	for _, dir := range dirs {
// 		notes = append(notes, dir.Name())
// 	}

// 	// Return notes
// 	return notes
// }

// Create a new project
// func (m *Memoir) CreateProject(name string) error {
// 	return os.MkdirAll(filepath.Join(m.Path, "projects", name), 0755)
// }

// Get the current project
// func (m *Memoir) GetCurrentProject() string {
// 	// Get current project from .current file
// 	file := filepath.Join(m.Path, ".current")

// 	// Read file contents
// 	proj, err := os.ReadFile(file)
// 	if err != nil {
// 		return ""
// 	}
// 	// Return project
// 	return string(proj)
// }

// Set the current project
// func (m *Memoir) SetCurrentProject(name string) error {
// 	// Set path to the project
// 	dir := filepath.Join(m.Path, name)

// 	// Check if project exists
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		return errors.New("project doesn't exist")
// 	}

// 	// Set the file to write to
// 	file := filepath.Join(m.Path, ".current")

// 	// Write to file
// 	err := os.WriteFile(file, []byte(name), 0644)
// 	if err != nil {
// 		return errors.New("error switching project")
// 	}

// 	// Set the current project
// 	m.CurrentProject = name

// 	// Return nil
// 	return nil
// }
