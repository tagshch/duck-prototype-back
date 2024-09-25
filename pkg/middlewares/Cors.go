package middlewares

import "net/http"

func Cors(next http.Handler, clientOrigin string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Origin", clientOrigin)

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Encoding, Authorization")
			return
		}

		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Methods, Access-Control-Allow-Headers, Access-Control-Allow-Credentials, Access-Control-Max-Age, Content-Encoding")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Referer, Content-Type, Content-Length, Accept-Encoding, Content-Encoding, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400")

		next.ServeHTTP(w, r)
	})
}
