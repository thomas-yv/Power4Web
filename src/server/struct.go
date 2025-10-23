package server

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	Title string
	Todos []Todo
}
