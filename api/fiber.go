package api

import (
	fiberHandlers "Savings/api/handlers/fiber"
	"Savings/api/middleware"
	"Savings/api/middleware/jwt"
	jobDomain "Savings/pkg/domain/job"
	paDomain "Savings/pkg/domain/personal_account"
	patDomain "Savings/pkg/domain/personal_account_transaction"
	entRepo "Savings/pkg/repositories/ent"
	"Savings/utils/logger"
	"crypto/sha256"
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
	"time"
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
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
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

	api.Use(jwt.New(jwt.Config{
		Secret: viper.GetString("JWT_KEY"),
		Expiry: time.Duration(15) * time.Minute,
	}))

	personalAccountRepository := entRepo.NewEntPersonalAccountRepository()
	personalAccountService := paDomain.NewPersonalAccountService(personalAccountRepository)
	personalAccountHandler := fiberHandlers.NewPersonalAccountHandler(personalAccountService)

	v1.Get("personal-accounts", personalAccountHandler.Get)
	v1.Get("personal-accounts/:id", personalAccountHandler.GetById)
	v1.Post("personal-accounts", personalAccountHandler.Create)

	v1.Post("personal-accounts/:id/deposit", personalAccountHandler.Deposit)
	v1.Post("personal-accounts/:id/withdraw", personalAccountHandler.Withdraw)

	personalAccountTransactionRepository := entRepo.NewEntPersonalAccountTransactionRepository()
	personalAccountTransactionService := patDomain.NewPersonalAccountTransactionService(personalAccountTransactionRepository)
	personalAccountTransactionHandler := fiberHandlers.NewPersonalAccountTransactionHandler(personalAccountTransactionService)

	v1.Get("personal-account-transactions", personalAccountTransactionHandler.Get)
	v1.Get("personal-accounts-transactions/:id", personalAccountTransactionHandler.GetById)

	jobRepository := entRepo.NewEntJobRepository()
	jobService := jobDomain.NewJobService(jobRepository, personalAccountRepository)
	jobHandler := fiberHandlers.NewJobHandler(jobService)

	v1.Get("jobs", jobHandler.Get)

	jobRouter := app.Group("/jobs").
		Use(func(c *fiber.Ctx) error {

			h := sha256.New()
			h.Write([]byte(viper.GetString("JWT_KEY")))

			if fmt.Sprintf("%x", h.Sum(nil)) == c.Query("key") {
				return c.Next()
			}

			return c.SendStatus(fiber.StatusUnauthorized)
		})

	jobRouter.Post("interest/calculate", jobHandler.CalculateInterest)
	jobRouter.Post("interest/allocate", jobHandler.AllocateInterest)

}
