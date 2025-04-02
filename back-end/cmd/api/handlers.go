package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (app *application) Home(c *gin.Context) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Food4Student",
		Version: "1.0.0",
	}

	app.writeJSON(c, http.StatusOK, payload)
}

func (app *application) AllUsers(c *gin.Context) {
	users, err := app.DB.GetAllUsers()
	if err != nil {
		app.errorJSON(c, err)
		return
	}
	app.writeJSON(c, http.StatusOK, users)
}

func (app *application) AllRestaurants(c *gin.Context) {
	restaurants, err := app.DB.GetAllRestaurants()
	if err != nil {
		app.errorJSON(c, err)
		return
	}
	app.writeJSON(c, http.StatusOK, restaurants)
}

func (app *application) RestaurantDetails(c *gin.Context) {
	// Get id from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		app.errorJSON(c, errors.New("invalid restaurant ID"), http.StatusBadRequest)
		return
	}

	restaurant, err := app.DB.GetRestaurant(uint(id))
	if err != nil {
		app.errorJSON(c, err)
		return
	}

	app.writeJSON(c, http.StatusOK, restaurant)
}

func (app *application) authenticate(c *gin.Context) {
	// read json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.readJSON(c, &requestPayload); err != nil {
		app.errorJSON(c, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(c, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(c, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// create a jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(c, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	c.SetCookie(
		refreshCookie.Name,
		refreshCookie.Value,
		refreshCookie.MaxAge,
		refreshCookie.Path,
		refreshCookie.Domain,
		refreshCookie.Secure,
		refreshCookie.HttpOnly,
	)

	app.writeJSON(c, http.StatusAccepted, tokens)
}
