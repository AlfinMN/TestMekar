package domain

import (
	"BackendTestMekar/domain/user"
	"database/sql"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {

	userRepo := user.InitUserRepoImpl(db)
	userUsecase := user.InitUsecaseImpl(userRepo)
	user.UserController(r, userUsecase)
}
