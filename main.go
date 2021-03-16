package main

import "github.com/gofiber/fiber/v2"

type BaseResp struct {
	Args    map[string]interface{} `json:"args"`
	Headers map[string]string      `json:"headers"`
	Origin  string                 `json:"origin"`
	URL     string                 `json:"url"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber httpbin 👋!")
	})

	// http methods
	app.Get("/get", GetMethod)
	app.Post("/post", PostMethod)
	app.Put("/put", PutMethod)
	app.Patch("/patch", PatchMethod)
	app.Delete("/delete", DeleteMethod)

	// anything
	app.Delete("/anything", DeleteAnything)
	app.Get("/anything", GetAnything)
	app.Patch("/anything", PatchAnything)
	app.Post("/anything", PostAnything)
	app.Put("/anything", PutAnything)
	app.Delete("/anything/:anything", DeleteAnythingAnything)
	app.Get("/anything/:anything", GetAnythingAnything)
	app.Patch("/anything/:anything", PatchAnythingAnything)
	app.Post("/anything/:anything", PostAnythingAnything)
	app.Put("/anything/:anything", PutAnythingAnything)

	// redirects
	app.Get("/absolute-redirect/:n", GetAbsoluteRedirect)
	app.Delete("/redirect-to", DeleteRedirectTo)
	app.Get("/redirect-to", GetRedirectTo)
	app.Post("/redirect-to", PostRedirectTo)
	app.Patch("/redirect-to", PatchRedirectTo)
	app.Put("/redirect-to", PutRedirectTo)
	app.Get("/redirect/:n", GetRedirect)
	app.Get("/relative-redirect/:n", GetRelativeRedirect)

	// images
	app.Get("/image", GetImage)
	app.Get("/image/jpeg", GetImageJPEG)
	app.Get("/image/jpg", GetImageJPEG)
	app.Get("/image/png", GetImagePNG)
	app.Get("/image/svg", GetImageSVG)
	app.Get("/image/webp", GetImageWebP)

	// cookies

	app.Get("/cookies", GetCookies)
	app.Get("/cookoes/delete", GetCookiesDelete)
	app.Get("/cookies/set", GetCookiesSet)
	app.Get("/cookies/set/:name/:value", GetCookiesSetValue)

	// dynamic data

	app.Get("/base64/:value", GetBase64Value)
	app.Get("/bytes/:n", GetBytesN)
	app.Delete("/delay/:delay", DeleteDelay)
	app.Get("/delay/:delay", GetDelay)
	app.Post("/delay/:delay", PostDelay)
	app.Patch("/delay/:delay", PatchDelay)
	app.Put("/delay/:delay", PutDelay)
	app.Get("/drip", GetDrip)
	app.Get("/links/:n/:offset", GetLinks)
	app.Get("/range", GetRange)
	app.Get("/stream-bytes/:n", GetStreamBytes)
	app.Get("/stream/:n", GetStream)
	app.Get("/uuid", GetUUID)

	// response format
	app.Get("/btotli", GetRespBrotli)
	app.Get("/deflate", GetRespDeflate)
	app.Get("/deny", GetRespDeny)
	app.Get("/encoding/utf8", GetRespEncodingUTF8)
	app.Get("/gzip", GetRespGZip)
	app.Get("/html", GetRespHTML)
	app.Get("/json", GetRespJSON)
	app.Get("/robors.txt", GetRespRobotsTXT)
	app.Get("/xml", GetRespXML)

	// response inspection

	app.Get("/cache", GetRespCache)
	app.Get("/cache/:value", GetRespCacheValue)
	app.Get("/etag/:etag", GetRespEtag)
	app.Get("/response-headers", GetRespHeaders)
	app.Post("/response-headers", PostRespHeaders)

	// request inspection

	app.Get("/headers", GetHeaders)
	app.Get("/ip", GetIP)
	app.Get("/user-agent", GetUserAgent)

	// status code

	app.Delete("/status/:codes", DeleteStatus)
	app.Get("/status/:codes", GetStatus)
	app.Patch("/status/:codes", PatchStatus)
	app.Post("/status/:codes", PostStatus)
	app.Put("/status/:codes", PutStatus)

	// auth

	app.Get("/basic-auth/:user/:passwd", GetBasicAuth)
	app.Get("/bearer", GetBearerAuth)
	app.Get("/digest-auth/:qop/:user/:passwd", GetDigestAuth)
	app.Get("/digest-auth/:qop/:user/:passwd/:algorithm", GetDigestAuthA)
	app.Get("/digest-auth/:qop/:user/:passwd/:algorithm/:stale_after", GetDigestAuthAS)
	app.Get("/hiden-basic-auth/:user/:passwd", GetHidenBasicAuth)

	app.Listen(":3000")
}
