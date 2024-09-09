package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"MessengerV2/api/database"
	"MessengerV2/api/pusherclient"
)

func Auth(c *fiber.Ctx) error{
	userToken, userTokenErr := retrieveTokenFromCookie(c.Cookies("user"))//Retrieve value from cookie and parse token from it.
	anonymousToken, anonymousTokenErr := retrieveTokenFromCookie(c.Cookies("anonymous"))

	//If the user has the anonymous cookie, and but their sign in cookie expired, delete the anonymous user.
	if userTokenErr != nil && anonymousTokenErr == nil{

		//Extract the claims from the anonymous token, and proceed with deletion.
		if claims, ok := anonymousToken.Claims.(jwt.MapClaims); ok && anonymousToken.Valid {
			username := claims["username"].(string)
			
 			if err := database.DeleteAnonymousUser(username); err != nil{
				return c.Status(http.StatusInternalServerError).SendString(err.Error())
			}

			if err := pusherclient.GetPusherClient().Trigger("Public", "anonymous_user_deleted", username); err != nil{
				return c.Status(http.StatusInternalServerError).SendString(err.Error())
			}
		}

		// c.Cookie(&fiber.Cookie{
		// 	Name: "anonymous",
		// 	MaxAge: 0,
		// })
		c.ClearCookie()
		return c.Status(http.StatusUnauthorized).SendString("A valid token is required for entry.")
	}

	//If there is no anonymous cookie, just check to see if the user has a valid token.
	if userTokenErr != nil {
		return c.Status(http.StatusUnauthorized).SendString("A valid token is required for entry.")
	}

	//Parse the token and pull out the username from the token.
	if claims, ok := userToken.Claims.(jwt.MapClaims); ok && userToken.Valid {
		c.SetUserContext(context.WithValue(c.Context(), "token", claims))
	}

	return c.Next()
}

func retrieveTokenFromCookie(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return token, nil
}