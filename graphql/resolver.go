package graphql

import (
	"context"
	"github.com/NPC-GO/MaJaJalist-backend/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type resolver struct {
}

func (r *resolver) MyMutation() MyMutationResolver {
	return &myMutation{r}
}
func (r *resolver) MyQuery() MyQueryResolver {
	return &myQuery{r}
}
func (r *resolver) Todo() TodoResolver {
	return &todo{r}
}
func (r *resolver) TodoStatus() TodoStatusResolver {
	return &todoStatus{r}
}
func (r *resolver) User() UserResolver {
	return &user{r}
}

type myMutation struct{ *resolver }

func (r *myMutation) Login(ctx context.Context, input *model.LoginInput) (bool, error) {
	panic("not implemented")
}
func (r *myMutation) DeleteTodo(ctx context.Context, ids []string) ([]string, error) {
	panic("not implemented")
}
func (r *myMutation) ChangeTextInTodo(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}
func (r *myMutation) CompleteTodo(ctx context.Context, ids []string) ([]string, error) {
	panic("not implemented")
}
func (r *myMutation) CreateTodo(ctx context.Context, input *model.CreateTodoInput) (*model.Todo, error) {
	panic("not implemented")
}
func (r *myMutation) ChangeTodoConfigValue(ctx context.Context, todoConfig model.ChangeTodoConfigValueInput) ([]*model.TodoConfig, error) {
	panic("not implemented")
}
func (r *myMutation) AddAvailableTodoConfig(ctx context.Context, key []string) ([]string, error) {
	panic("not implemented")
}
func (r *myMutation) DeleteAvailableTodoConfig(ctx context.Context, key []string) ([]string, error) {
	panic("not implemented")
}
func (r *myMutation) CreateUser(ctx context.Context, input *model.CreateUserInput) (*model.User, error) {
	panic("not implemented")
}
func (r *myMutation) CheckIDAvailability(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *myMutation) ChangeMyUserConfigValue(ctx context.Context, input []*model.ConfigInput) ([]*model.Config, error) {
	panic("not implemented")
}
func (r *myMutation) AddAvailableUserConfig(ctx context.Context, key []string) ([]string, error) {
	panic("not implemented")
}
func (r *myMutation) DeleteAvailableUserConfig(ctx context.Context, key []string) ([]string, error) {
	panic("not implemented")
}

type myQuery struct{ *resolver }

func (r *myQuery) TodobyUserID(ctx context.Context, id string) ([]*model.Todo, error) {
	panic("not implemented")
}
func (r *myQuery) AvailableTodoConfig(ctx context.Context) ([]string, error) {
	panic("not implemented")
}
func (r *myQuery) MyUserData(ctx context.Context) (*model.User, error) {
	panic("not implemented")
}
func (r *myQuery) MyUserConfig(ctx context.Context) ([]*model.Config, error) {
	panic("not implemented")
}
func (r *myQuery) AvailableUserConfig(ctx context.Context) ([]string, error) {
	panic("not implemented")
}

type todo struct{ *resolver }

func (r *todo) Status(ctx context.Context, obj *model.Todo) (*model.TodoStatus, error) {
	panic("not implemented")
}
func (r *todo) Author(ctx context.Context, obj *model.Todo) (*model.TodoAuthorLayout, error) {
	panic("not implemented")
}
func (r *todo) TodoConfig(ctx context.Context, obj *model.Todo) ([]*model.TodoConfig, error) {
	panic("not implemented")
}

type todoStatus struct{ *resolver }

func (r *todoStatus) Readonly(ctx context.Context, obj *model.TodoStatus) (bool, error) {
	panic("not implemented")
}

type user struct{ *resolver }

func (r *user) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	panic("not implemented")
}
func (r *user) UserConfig(ctx context.Context, obj *model.User) ([]*model.Config, error) {
	panic("not implemented")
}
