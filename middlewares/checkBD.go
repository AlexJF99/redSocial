package middlewares

import (
	"net/http"
	"redSocial/bd"
)

/*CheckBD allows knows the BD status*/
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
