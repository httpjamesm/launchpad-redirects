package main

import (
	"launchpad-redirects/lists"
	"net/url"

	"github.com/gin-gonic/gin"
)

type instance struct {
	URL         string
	FaviconURL  string
	Name        string
	Description string
}

func main() {
	app := gin.Default()

	app.Static("/static", "./static")

	app.LoadHTMLGlob("templates/*")

	app.Use(func(c *gin.Context) {
		// add strict CSP

		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'none'; img-src 'self'; style-src 'self';")

		c.Next()
	})

	app.GET("/", func(c *gin.Context) {
		inputUrl := c.Query("url")
		if inputUrl != "" {
			// parse url domain
			parsed, err := url.Parse(inputUrl)
			if err != nil {
				c.HTML(400, "home.html", gin.H{
					"Error": "Invalid URL",
				})
				return
			}

			originalParsedHost := parsed.Host

			// check if domain is in list

			if lists.ContainsString(lists.RedditDomains, parsed.Host) {
				// replace domain with discuss.whatever.social
				parsed.Host = "discuss.whatever.social"
			} else if lists.ContainsString(lists.YouTubeDomains, parsed.Host) {
				// replace domain with watch.whatever.social
				parsed.Host = "watch.whatever.social"
			} else if lists.ContainsString(lists.StackOverflowDomains, parsed.Host) {
				// replace domain with code.whatever.social
				parsed.Host = "code.whatever.social"
			} else if lists.ContainsString(lists.TikTokDomains, parsed.Host) {
				// replace domain with cringe.whatever.social
				parsed.Host = "cringe.whatever.social"
			} else if lists.ContainsString(lists.YouTubeMusicDomains, parsed.Host) {
				// replace domain with listen.whatever.social
				parsed.Host = "listen.whatever.social"
			} else if lists.ContainsString(lists.GeniusDomains, parsed.Host) {
				// replace domain with sing.whatever.social
				parsed.Host = "sing.whatever.social"
			} else if lists.ContainsString(lists.ImgurDomains, parsed.Host) {
				// replace domain with images.whatever.social
				parsed.Host = "images.whatever.social"
			} else if lists.ContainsString(lists.ImdbDomains, parsed.Host) {
				// replace domain with binge.whatever.social
				parsed.Host = "binge.whatever.social"
			} else {
				// send error
				c.HTML(400, "home.html", gin.H{
					"Error": "Unsupported domain",
				})
				return
			}

			if parsed.Host != originalParsedHost {
				// redirect
				c.Redirect(302, parsed.String())
				return
			}

		}

		c.HTML(200, "home.html", gin.H{
			"Instances": []instance{
				{
					URL:         "https://code.whatever.social",
					FaviconURL:  "/static/assets/apps/code.png",
					Name:        "AnonymousOverflow",
					Description: "Whatever's own frontend for StackOverflow.",
				},
				{
					URL:         "https://discuss.whatever.social",
					FaviconURL:  "/static/assets/apps/libreddit.png",
					Name:        "Libreddit",
					Description: "Alternative frontend for Reddit.",
				},
				{
					URL:         "https://watch.whatever.social",
					FaviconURL:  "/static/assets/apps/piped.png",
					Name:        "Piped",
					Description: "Alternative frontend for YouTube.",
				},
				{
					URL:         "https://cringe.whatever.social",
					FaviconURL:  "/static/assets/apps/proxitok.png",
					Name:        "ProxiTok",
					Description: "Alternative frontend for TikTok.",
				},
				{
					URL:         "https://listen.whatever.social",
					FaviconURL:  "/static/assets/apps/hyperpipe.png",
					Name:        "Hyperpipe",
					Description: "Alternative frontend for YouTube Music.",
				},
				{
					URL:         "https://sing.whatever.social",
					FaviconURL:  "/static/assets/apps/dumb.png",
					Name:        "Dumb",
					Description: "Alternative frontend for Genius.",
				},
				{
					URL:         "https://images.whatever.social",
					FaviconURL:  "/static/assets/apps/rimgo.png",
					Name:        "Rimgo",
					Description: "Alternative frontend for Imgur.",
				},
				{
					URL:         "https://nerd.whatever.social",
					FaviconURL:  "/static/assets/apps/breezewiki.png",
					Name:        "Breezewiki",
					Description: "Alternative frontend for Fandom.",
				},
				{
					URL:         "https://binge.whatever.social",
					FaviconURL:  "/static/assets/apps/libremdb.png",
					Name:        "Libremdb",
					Description: "Alternative frontend for IMDb.",
				},
				{
					URL:         "https://notavault.com",
					FaviconURL:  "/static/assets/apps/vaultwarden.png",
					Name:        "Vaultwarden",
					Description: "Whatever's self-hosted instance of Vaultwarden, a Bitwarden fork.",
				},
			},
		})
	})

	app.Run(":8080")
}
