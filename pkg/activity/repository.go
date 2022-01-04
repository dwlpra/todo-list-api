package activity

import (
	"context"
	"database/sql"
	"errors"
	"time"

	resp "github.com/dwlpra/todo-list-api/libs/response"
)

type Repository interface {
	GetAll() chan resp.Result
	GetOne(id string) chan resp.Result
	Create(req RequestBody) chan resp.Result
	Delete(id string) chan resp.Result
	Update(req RequestBody, id string) chan resp.Result
}

type repository struct {
	connDB *sql.DB
}

func NewRepository(connDB *sql.DB) Repository {
	return &repository{
		connDB: connDB,
	}
}

func (repo *repository) GetAll() chan resp.Result {
	result := make(chan resp.Result)
	go func() {
		activities := []Activities{}
		ctx := context.Background()
		rows, err := repo.connDB.QueryContext(ctx, "SELECT * FROM activities WHERE deleted_at IS NULL")
		if err != nil {
			result <- resp.Result{Err: err}
			return
		}
		for rows.Next() {
			activity := Activities{}
			if err := rows.Scan(
				&activity.ID,
				&activity.Email,
				&activity.Title,
				&activity.CreatedAt,
				&activity.UpdatedAt,
				&activity.DeletedAt,
			); err != nil {
				result <- resp.Result{Err: err}
				return
			}
			activities = append(activities, activity)
		}

		result <- resp.Result{Data: activities}
		defer rows.Close()

	}()
	return result
}

func (repo *repository) GetOne(id string) chan resp.Result {
	result := make(chan resp.Result)
	go func() {

		row := repo.connDB.QueryRow("SELECT * FROM activities WHERE id=? and deleted_at IS NULL", id)
		activity := Activities{}
		switch err := row.Scan(
			&activity.ID,
			&activity.Email,
			&activity.Title,
			&activity.CreatedAt,
			&activity.UpdatedAt,
			&activity.DeletedAt,
		); err {
		case sql.ErrNoRows:
			result <- resp.Result{Err: err}
			return
		case nil:
			result <- resp.Result{Data: activity}
			return
		}

	}()
	return result
}

func (repo *repository) Create(req RequestBody) chan resp.Result {
	result := make(chan resp.Result)
	go func() {
		ctx := context.Background()
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		res, err := repo.connDB.ExecContext(ctx, "INSERT INTO activities (email, title, created_at, updated_at) VALUES(?,?,?,?)", req.Email, req.Title, timeNow, timeNow)
		if err != nil {
			result <- resp.Result{Err: err}
			return
		}
		id, _ := res.LastInsertId()
		activity := InsertActivity{
			ID:        id,
			Email:     req.Email,
			Title:     req.Title,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		}

		result <- resp.Result{Data: activity}
	}()
	return result
}

func (repo *repository) Delete(id string) chan resp.Result {
	result := make(chan resp.Result)
	go func() {
		ctx := context.Background()
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		res, err := repo.connDB.ExecContext(ctx, "UPDATE activities SET deleted_at=? WHERE id=? and deleted_at IS NULL", timeNow, id)
		if err != nil {
			result <- resp.Result{Err: err}
			return
		}
		status, _ := res.RowsAffected()
		if status == 0 {
			err := errors.New("no row deleted")
			result <- resp.Result{Err: err}
			return
		}
		result <- resp.Result{Data: resp.EmptyResp{}}

	}()
	return result
}

func (repo *repository) Update(req RequestBody, id string) chan resp.Result {
	result := make(chan resp.Result)
	go func() {
		ctx := context.Background()
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		_, err := repo.connDB.ExecContext(ctx, "UPDATE activities SET title=?, updated_at=? WHERE id=? and deleted_at IS NULL", req.Title, timeNow, id)
		if err != nil {
			result <- resp.Result{Err: err}
			return
		}
		getData := <-repo.GetOne(id)
		result <- resp.Result{Data: getData.Data}

	}()
	return result
}
