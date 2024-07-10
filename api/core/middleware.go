package core

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/envvar"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

// SetupCoreMiddlewares sets up all the core middlewares that apply to all routes
func SetupCoreMiddlewares(app *fiber.App) {
	// cache middleware
	app.Use(cache.New(cache.Config{
		ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
			newCacheTime, _ := strconv.Atoi(c.GetRespHeader("Cache-Time", "600"))
			return time.Second * time.Duration(newCacheTime)
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
	}))

	// compress middleware
	app.Use(compress.New(compress.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/dont_compress"
		},
		Level: compress.LevelBestSpeed, // 1
	}))

	// cors middleware
	app.Use(cors.New())

	// csrf middleware
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUIDv4,
	}))

	// cookie middleware
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "secret-thirty-2-character-string",
	}))

	//envVar middleware
	app.Use("/expose/envvars", envvar.New())

	// etag middleware
	app.Use(etag.New(etag.Config{
		Weak: true,
	}))

	// favicon middleware
	app.Use(favicon.New())

	// healthcheck middleware
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessEndpoint: "/ready",
	}))

	// helmet middleware
	app.Use(helmet.New())

	// idempotency middleware
	app.Use(idempotency.New())

	// rate limit middleware
	app.Use(limiter.New())

	// request id middleware
	app.Use(requestid.New())

	// logger middleware
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	// metrics monitoring endpoint
	app.Get("/metrics", monitor.New())

	// recover middleware
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
}
