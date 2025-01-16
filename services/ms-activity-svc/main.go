package main

import (
	"fmt"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/config"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/controllers"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/models"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/repositories"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase(env config.Environment) *gorm.DB {
	fmt.Println("Connect to database ....")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
		env.DATABASE_HOST,
		env.DATABASE_USER,
		env.DATABASE_PASSWORD,
		env.DATABASE_NAME,
		env.DATABASE_PORT,
		env.DATABASE_SCHEMA,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Test query
	var result string
	db.Raw("SELECT 1;").Scan(&result)

	fmt.Printf("Connection successfull. Result from test SQL: %s\n", result)

	// Migrations
	db.AutoMigrate(&models.Activity{})

	return db
}

func setupRouter(activityController *controllers.ActivityController) *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/v1")
	// v1.Use(middlewares.JWTAuthMiddleware())
	{
		departments := v1.Group("/activity")
		{
			// departments.POST("/", departmentController.Create)
			departments.GET("/", activityController.GetAllDepartments)
			// departments.GET("/:id", departmentController.GetOneDepartment)
			// departments.PATCH("/:id", departmentController.UpdateOneDepartment)
			// departments.DELETE("/:id", departmentController.DeleteOneDepartment)
			// departments.PATCH("/update-department-id", departmentController.UpdateDepartmentId)
		}
	}

	return r
}

func main() {
	// Load .env
	wd, _ := os.Getwd()
	envpath := path.Join(wd, "services", "ms-activity-svc", ".env")
	godotenv.Load(envpath)
	cfg := config.EnvironmentConfig()

	// connect databases
	db := connectDatabase(cfg)

	activityRepo := repositories.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepo)
	activityController := controllers.NewActivityController(activityService)

	r := setupRouter(activityController)
	r.Run(":8080")
}
