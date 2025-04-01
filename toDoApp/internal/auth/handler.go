package auth

import (
	"fmt"
	"net/http"
	"toDo/configs"
	"toDo/pkg/jwt"
	"toDo/pkg/request"
	"toDo/pkg/response"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/session", handler.Session())
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

// Session handles session creation
// @Summary Create a session
// @Description Generates a session ID for the given phone number
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body SessionRequest true "Session Request"
// @Success 200 {object} SessionResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/session [post]
func (handler *AuthHandler) Session() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Session")

		body, err := request.HandleBody[SessionRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(*body)

		sessionId, err := handler.AuthService.GetSessionId(body.Phone)
		if err != nil {
			response.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Println(sessionId)

		resp := SessionResponse{
			SessionId: sessionId,
		}
		response.Json(w, resp, http.StatusOK)
	}
}

// Login handles user login
// @Summary Login a user
// @Description Authenticates a user using session ID and code, and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Login Request"
// @Success 200 {object} LoginResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/login [post]
func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")

		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(*body)

		sessionId, err := handler.AuthService.Login(body.SessionId, body.Code)
		if err != nil {
			response.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println(sessionId)

		j := jwt.NewJWT(handler.Config.Auth.Secret)
		token, err := j.Create(jwt.JWTData{SessionId: sessionId, Code: body.Code})
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := LoginResponse{
			Token: token,
		}
		response.Json(w, resp, http.StatusOK)
	}
}

// Register handles user registration
// @Summary Register a user
// @Description Registers a user using their phone number
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body RegisterRequest true "Register Request"
// @Success 200 {object} RegisterResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/register [post]
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")

		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(*body)

		email, err := handler.AuthService.Register(body.Phone)
		if err != nil {
			response.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println(email)

		resp := RegisterResponse{
			Message: "Register success",
		}
		response.Json(w, resp, http.StatusOK)
	}
}
