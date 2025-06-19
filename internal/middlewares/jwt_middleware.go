package middlewares

import (
	"net/http"
	"nls-go-messaging/internal/utils"
	"strings"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("votre_cle_secrete") // À stocker dans une variable d'env en prod

/*
	 func JWTMiddleware(c *fiber.Ctx) error {
		//application := c.Context().Value(constants.ApplicationCtx).(constants.AppKey)
		 tokenStr := r.URL.Query().Get("token")
		   claims := &Claims{}
		   token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		   	return jwtKey, nil
		   })

		   if err != nil || !token.Valid {

			utils.RespondWithError(w, http.StatusUnauthorized,"Unauthorized")
		   	return
		   }
		   username := claims.Username

		jwtKey := c.Locals(constants.JWTKeyLocals).([]byte)

		if jwtKey != nil {
			c.Locals(constants.RequestJWTLocals, jwtKey)
		} else {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		return c.Next()
	}
*/
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.RespondWithError(w, http.StatusUnauthorized, "Missing or invalid Authorization header")
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		next.ServeHTTP(w, r)
	})
}
