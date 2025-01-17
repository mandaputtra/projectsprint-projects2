package handlers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-users-svc/config"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-users-svc/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Type
type UserCreateOrLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	Action   string `json:"action" binding:"required,oneof=create login"`
}

type UserUpdateRequest struct {
	Name            string `json:"name" binding:"min=4,max=52"`
	Email           string `json:"email" binding:"required,email"`
	UserImageUri    string `json:"userImageUri" binding:"required,uri"`
	CompanyName     string `json:"companyName" binding:"min=4,max=52"`
	CompanyImageUri string `json:"companyImageUri" binding:"required,uri"`
}

type APIEnv struct {
	DB *gorm.DB
}

// Utils
func validateURIWithTLD(uri string) bool {
	parsedURI, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}
	return strings.Contains(parsedURI.Host, ".")
}

// Services
func (a *APIEnv) CreateOrLogin(c *gin.Context) {
	db := a.DB
	cfg := config.EnvironmentConfig()
	var userRequest UserCreateOrLoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create or log in the user based on the action
	switch userRequest.Action {
	case "create":
		user := database.User{Email: userRequest.Email, Password: userRequest.Password}
		// Encrypt the password before saving it to the database
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
			return
		}
		user.Password = string(hashedPassword)
		result := db.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": result.Error,
			})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": user.Email,
		})

		tokenString, err := token.SignedString([]byte(cfg.JWT_SECRET))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"email": user.Email,
			"token": tokenString,
		})
		return
	case "login":
		var user database.User

		db.Where("email =?", userRequest.Email).First(&user)
		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email or password"})
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": user.Email,
		})
		tokenString, err := token.SignedString([]byte(
			cfg.JWT_SECRET,
		))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"email": user.Email,
			"token": tokenString,
		})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Action"})
		return
	}
}

func (a *APIEnv) GetUsers(c *gin.Context) {
	db := a.DB
	email := c.GetString("email")

	var user database.User
	db.Where("email = ?", email).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"email":           user.Email,
		"name":            user.Name,
		"userImageUri":    user.UserImageUri,
		"companyName":     user.CompanyName,
		"companyImageUri": user.CompanyImageUri,
	})
}

func (a *APIEnv) UpdateUser(c *gin.Context) {
	db := a.DB

	var userRequest UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !validateURIWithTLD(userRequest.UserImageUri) || !validateURIWithTLD(userRequest.CompanyImageUri) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad url"})
		return
	}

	email := c.GetString("email")

	var user database.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Name = userRequest.Name
	user.UserImageUri = userRequest.UserImageUri
	user.CompanyName = userRequest.CompanyName
	user.CompanyImageUri = userRequest.CompanyImageUri
	user.Email = userRequest.Email

	if err := db.Save(&user).Error; err != nil {
		if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": "Failed to update user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email":           user.Email,
		"name":            user.Name,
		"userImageUri":    user.UserImageUri,
		"companyName":     user.CompanyName,
		"companyImageUri": user.CompanyImageUri,
	})
}
