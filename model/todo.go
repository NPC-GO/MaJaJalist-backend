package model

type Todo struct {
	ID          string `json:"id"`
	TextContent string `json:"textContent" pg:"textcontent"`
	Completed   bool
	Deleted     bool
	AuthorID    string `pg:"authorid"`
}

type TodoAuthorLayout struct {
	tableName struct{} `pg:"users"`
	ID        string   `json:"id"`
	NickName  string   `json:"nickName" pg:"nickname"`
	Avatar    string   `json:"avatar"`
	Special   bool     `json:"special"`
}
type TodoStatus struct {
	Completed bool `json:"completed"`
	Deleted   bool `json:"deleted"`
}
type CreateTodoInput struct {
	TextContent string                 `json:"textContent"`
	Status      *CreateTodoInputStatus `json:"status"`
}

type CreateTodoInputStatus struct {
	Complete *bool `json:"complete"`
	Deleted  *bool `json:"deleted"`
}

type TodoConfig struct {
	tableName struct{}  `pg:"todoconfig"`
	TodoID    string    `json:"todoID" pg:"todoid"`
	Config    []*Config `json:"config" pg:"-"`
	Key       string    `pg:"key"`
	Value     string    `pg:"value"`
}

type ChangeTodoConfigValueInput struct {
	TodoID string         `json:"todoID"`
	Config []*ConfigInput `json:"config"`
}
