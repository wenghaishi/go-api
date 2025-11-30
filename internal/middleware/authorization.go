package middleware

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/wenghaishi/go-api/api"
	"github.com/wenghaishi/go-api/internal/tools"
	"net/http"
)

var UnAuthorizaedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizaedError)
			api.RequestErrorHandler(w, UnAuthorizaedError)
			return
		}

		var database *tools.DatabaseInerface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizaedError)
			api.RequestErrorHandler(w, UnAuthorizaedError)
			return
		}

		next.ServeHTTP(w, r)

	}
}
