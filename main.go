package main

import (
	"launchpad-redirects/lists"
	"net/url"

	"github.com/gin-gonic/gin"
)

type instance struct {
	URL        string
	FaviconURL string
	Name       string
}

func main() {
	app := gin.Default()

	app.Static("/static", "./static")

	app.LoadHTMLGlob("templates/*")

	app.GET("/", func(c *gin.Context) {
		inputUrl := c.Query("url")
		if inputUrl != "" {
			// parse url domain
			parsed, err := url.Parse(inputUrl)
			if err != nil {
				c.HTML(200, "home.html", gin.H{
					"error": err.Error(),
				})
				return
			}

			originalParsedHost := parsed.Host

			// check if domain is in list

			if lists.ContainsString(lists.RedditDomains, parsed.Host) {
				// replace domain with discuss.whatever.social
				parsed.Host = "discuss.whatever.social"
			}

			if lists.ContainsString(lists.TwitterDomains, parsed.Host) {
				// replace domain with read.whatever.social
				parsed.Host = "read.whatever.social"
			}

			if lists.ContainsString(lists.YouTubeDomains, parsed.Host) {
				// replace domain with watch.whatever.social
				parsed.Host = "watch.whatever.social"
			}

			if lists.ContainsString(lists.StackOverflowDomains, parsed.Host) {
				// replace domain with code.whatever.social
				parsed.Host = "code.whatever.social"
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
					URL:        "https://discuss.whatever.social",
					FaviconURL: "https://discuss.whatever.social/favicon.ico",
					Name:       "Libreddit",
				},
				{
					URL:        "https://read.whatever.social",
					FaviconURL: "https://read.whatever.social/favicon.ico",
					Name:       "Nitter",
				},
				{
					URL:        "https://watch.whatever.social",
					FaviconURL: "https://watch.whatever.social/favicon.ico",
					Name:       "Piped",
				},
				{
					URL:        "https://code.whatever.social",
					FaviconURL: "https://code.whatever.social/static/codecircles.png",
					Name:       "AnonymousOverflow",
				},
			},
		})
	})

	app.Run(":8080")
}
