package routes

import (
	"github.com/gofiber/fiber/v2"
	"testapi/internal/middleware"
	"testapi/internal/models"
	"testapi/internal/services"
)

func SetupAuthRoutes(app *fiber.App) {
	routes := app.Group("/auth")

	routes.Post("/login", middleware.ParseBodyWithValidation(loginHandler))
	routes.Post("/register", middleware.ParseBodyWithValidation(registerHandler))

	routes.Post("/change-password",
		middleware.JWTMiddleware(),
		middleware.ParseBodyWithValidation(changePasswordHandler))

	routes.Post("/refresh-token", middleware.ParseBodyWithValidation(refreshTokenHandler))

}

// LoginHandler Login godoc
//
//	@Summary		User login
//	@Description	Authenticates a user and returns a token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.LoginRequest	true	"Login Request Body"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		401		{object}	map[string]string
//	@Router			/auth/login [post]
func loginHandler(c *fiber.Ctx, req *models.LoginRequest) error {
	token, err := services.Authenticate(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(token)
}

// RegisterHandler Register godoc
//
//	@Summary		User registration
//	@Description	Registers a new user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.RegisterRequest	true	"User Request Body"
//	@Success		200		{object}	models.UserDto
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/auth/register [post]
func registerHandler(c *fiber.Ctx, req *models.RegisterRequest) error {
	user, err := services.Register(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

// ChangePasswordHandler ChangePassword godoc
//
//	@Summary		Change password
//	@Description	Changes the password of the authenticated user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body	models.ChangePasswordRequest	true	"Change Password Request Body"
//	@Security		BearerAuth
//	@Success		200	{object}	map[string]bool
//	@Failure		400	{object}	map[string]string
//	@Router			/auth/change-password [post]
func changePasswordHandler(c *fiber.Ctx, req *models.ChangePasswordRequest) error {
	uid := uint(c.Locals("user_id").(float64))
	userID := uint(uid)
	success, err := services.ChangePassword(userID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": success})
}

// RefreshTokenHandler RefreshToken godoc
//
//	@Summary		Refresh token
//	@Description	Refreshes the token of the authenticated user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.RefreshTokenRequest	true	"Refresh"
//
// Token Request Body"
//
//	@Success		200		{object}	models.TokenResponse
//	@Failure		400		{object}	map[string]string
//	@Router			/auth/refresh-token [post]
func refreshTokenHandler(c *fiber.Ctx, req *models.RefreshTokenRequest) error {
	response, err := services.RefreshToken(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(response)
}
