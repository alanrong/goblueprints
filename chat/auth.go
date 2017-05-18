package main

import "net/http"

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		//not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	if err != nil {
		//some other err
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//success
	h.next.ServeHTTP(w, r)
}

//MustAuth chain handler after authHandler
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
