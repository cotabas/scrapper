package main

import (
  "fmt"
  "log"
  //"net/http"
  //"net/url"

  "github.com/gocolly/colly"
)

func main() {
  // Create a new collector
  c := colly.NewCollector(
    // Attach a debugger to the collector
//    colly.Debugger(&colly.DebugLog{}),
    colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"),
    colly.AllowedDomains("www.177milkstreet.com"),
  )

  // The URL for the login action
  loginURL := "https://www.177milkstreet.com/login"

  // The data you need to send for a login
  loginData := map[string]string{
    "loginName": "cotabas@gmail.com",
    "password": "Qwu824r2!",
  }

  // Handle login
  err := c.Post(loginURL, loginData)
  if err != nil {
    log.Fatal("Login failed:", err)
    fmt.Println("Login failed:", err)
  }

  // After logging in, visit the page you want to scrape
  c.OnHTML("html selector", func(e *colly.HTMLElement) {
    // Scrape information
    fmt.Println("First name:", e.ChildText("#first-name"))
  })

  // Visit a page that requires authentication
  err = c.Visit("https://example.com/protected-page")
  if err != nil {
    log.Fatal("Visiting failed:", err)
  }

}
