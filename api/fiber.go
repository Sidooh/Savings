package api

import (
	fiberHandlers "Savings/api/handlers/fiber"
	"Savings/api/middleware"
	domain "Savings/pkg/domain/personal_account"
	entRepo "Savings/pkg/repositories/ent"
	"Savings/utils/logger"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func FiberServer() *fiber.App {
	// Create a new fiber instance with custom config
	app := fiber.New(fiber.Config{
		Prefork: viper.GetBool("PREFORK"),

		//EnablePrintRoutes: true,

		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			logger.Log.Error(err.Error())
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})

	setMiddleware(app)
	setHealthCheckRoutes(app)
	setHandlers(app)

	//data, _ := json.MarshalIndent(app.GetRoutes(true), "", "  ")
	//fmt.Print(string(data))

	return app
}

func setHealthCheckRoutes(app *fiber.App) {
	app.Get("/200", func(ctx *fiber.Ctx) error {
		return ctx.JSON("200")
	}).Name("health")
}

func setMiddleware(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(cors.New())

	//app.Use(limiter.New(limiter.Config{Max: viper.GetInt("RATE_LIMIT")}))
	app.Use(recover.New())
	app.Use(fiberLogger.New(fiberLogger.Config{
		Output:     logger.GetLogFile("stats.log"),
		TimeFormat: "2006-01-02 15:04:05",
		//Format:     "{'time': '${time}', 'status': ${status}, 'latency': ${latency}, 'ip': ${ip}, 'method': ${method}, 'path': ${path}, 'err': ${error}}\n",
		Format: "[${time}] ${status}: ${latency}, ${ip}, ${method}, ${path}, ${error}\n",
	}))

	app.Use(favicon.New(favicon.Config{
		Next: func(c *fiber.Ctx) bool {
			return true
		},
	}))

	middleware.Validator = validator.New()
}

func setHandlers(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	personalAccountRepository := entRepo.NewEntPersonalAccountRepository()
	personalAccountService := domain.NewPersonalAccountService(personalAccountRepository)
	personalAccountHandler := fiberHandlers.NewPersonalAccountHandler(personalAccountService)

	v1.Get("personal-accounts", personalAccountHandler.Get)
	v1.Get("personal-accounts/:id", personalAccountHandler.GetById)
	v1.Post("personal-accounts", personalAccountHandler.Create)

}
