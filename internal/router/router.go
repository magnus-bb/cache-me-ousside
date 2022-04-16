package router

import (
	"github.com/NormalReedus/cache-me-ousside/cache"
	"github.com/NormalReedus/cache-me-ousside/internal/config"
	"github.com/NormalReedus/cache-me-ousside/internal/logger"
	"github.com/gofiber/fiber/v2"
)

//TODO: tests for router https://dev.to/koddr/go-fiber-by-examples-testing-the-application-1ldf
func Start(conf *config.Config, port string, cache *cache.LRUCache) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true, // has own HiMom message
		Immutable:             true, // muy importante - makes sure that OriginalUrl() cannot mutate cached endpoints somehow
	})

	// Make cache available in all handlers with ctx.Locals("cache").(*cache.LRUCache)
	app.Use(injectCtxCache(cache))

	// Will loop through methods, endpoints, and patterns and set a middleware for each that removes cache entries when patterns are matched
	setBustingEndpoints(app, conf)

	// Will loop through cachable endpoints in config and set route handlers + middleware to handle caching on those routes
	setCachingEndpoints(app, conf)

	// Any non-cache / non-cache-busting requests should just proxy directly to the original API
	app.Use("*", createProxyHandler(conf.ApiUrl)) // default behavior

	logger.HiMom(conf.ApiUrl, port)

	logger.Panic(app.Listen("localhost:" + port))
}

// Will loop through methods, endpoints, and patterns and set a middleware for each that removes cache entries when patterns are matched.
// e.g. POST to /users could remove all cache entries that match the pattern ^/users or ^/users/:id etc.
func setBustingEndpoints(app *fiber.App, conf *config.Config) {
	for method, endpointMap := range conf.Bust {
		for endpoint, patterns := range endpointMap {
			app.Add(method, endpoint, createBustMiddleware(patterns))
		}
	}
}

// Sets route handlers and middleware that:
// 1. uses cache if anything is cached, if not, then
// 2. proxies the request to apiUrl and gets a response, then
// 3. sets the response in the cache
func setCachingEndpoints(app *fiber.App, conf *config.Config) {
	for _, endpoint := range conf.Cache {
		app.Get(
			endpoint,

			readCacheMiddleware,

			createProxyMiddleware(conf.ApiUrl), // needs ApiUrl to proxy request

			writeCacheMiddleware,
		)
	}
}

func injectCtxCache(cache *cache.LRUCache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("cache", cache)
		ctx.Next()
		return nil
	}
}
