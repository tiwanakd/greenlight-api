package main

import (
	"net/http"
)

func (app *application) homeHanlder(w http.ResponseWriter, r *http.Request) {
	type homeProperties []struct {
		AllowedMedthod string `json:"allowed_method"`
		Endpoint       string `json:"endpoint"`
		Description    string `json:"description"`
	}

	homeview := map[string]any{
		"healthcheck": map[string]string{
			"endpoint":    "/v1/healthcheck",
			"description": "check the health of the api",
		},
		"user": homeProperties{
			{
				AllowedMedthod: http.MethodPost,
				Endpoint:       "/v1/users",
				Description:    "register a new user",
			},
			{
				AllowedMedthod: http.MethodPut,
				Endpoint:       "/v1/users/activated",
				Description:    "activate the registered user",
			},
			{
				AllowedMedthod: http.MethodPut,
				Endpoint:       "/v1/users/password",
				Description:    "change password with passoword reset token",
			},
		},
		"movie": homeProperties{
			{
				AllowedMedthod: http.MethodGet,
				Endpoint:       "/v1/movies",
				Description:    "get all the movies",
			},
			{
				AllowedMedthod: http.MethodPost,
				Endpoint:       "/v1/movies",
				Description:    "create new movie",
			},
			{
				AllowedMedthod: http.MethodGet,
				Endpoint:       "/v1/movies/:id",
				Description:    "get movie with given id",
			},
			{
				AllowedMedthod: http.MethodPatch,
				Endpoint:       "/v1/movies/:id",
				Description:    "update movie with given id",
			},
			{
				AllowedMedthod: http.MethodDelete,
				Endpoint:       "/v1/movies/:id",
				Description:    "delete the movie with given id",
			},
		},
		"token": homeProperties{
			{
				AllowedMedthod: http.MethodPost,
				Endpoint:       "/v1/tokens/activation",
				Description:    "recveive a new token for activation",
			},
			{
				AllowedMedthod: http.MethodPost,
				Endpoint:       "/v1/tokens/authentication",
				Description:    "autheticate to get authentication token",
			},
			{
				AllowedMedthod: http.MethodPost,
				Endpoint:       "/v1/tokens/password-reset",
				Description:    "get a password reset token",
			},
		},
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"greenlight_api": homeview}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
