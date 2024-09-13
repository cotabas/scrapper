package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

func main() {
  //get the links to all the recipes
  //getLinx(getFirstPage())
  inFile, _ := os.ReadFile("output.txt")
  linx := strings.Split(string(inFile), "\n")

  for i, line := range linx[:len(linx) - 1] {
    curl := exec.Command("curl", line)
    out, err := curl.Output()
    if err != nil { fmt.Println("curl output error" , err) }

    recipe := strings.Split(string(out), `<script type="application/ld+json">`)[1]
    recipe = strings.Split(recipe, `</script>`)[0]

    outFile, err := os.Create("../recipes/ms" + fmt.Sprint(i) + ".json")
    if err != nil { fmt.Println("file open error" , err) }
    outFile.WriteString(recipe)
    outFile.Sync()

    fmt.Println(i, line)
  }
}

func getLinx(linx []string) {
  outFile, err := os.Create("output.txt")
  if err != nil { fmt.Println("file open error" , err) }

  for i := 2; i < 181; i++ {
    recipeLink := "https://www.177milkstreet.com/recipes/p" + fmt.Sprint(i)
    fmt.Println(recipeLink)
    curl := exec.Command("curl", recipeLink) 
    out, err := curl.Output()
    if err != nil { fmt.Println("curl output error" , err) }

    lines := strings.Split(string(out), "\n")
    for _, line := range lines {
      if strings.Contains(line, `class="recipe-card__image-link"`) {
        line = strings.Split(line, `"`)[1]
        linx = append(linx, line)
      }
    }
  }
  for _, line := range linx {
    outFile.WriteString(line + "\n")
    outFile.Sync()
    fmt.Println(line)
  }
}

func getFirstPage() []string {
  linx := []string{}

  recipeLink := "https://www.177milkstreet.com/recipes"
  curl := exec.Command("curl", recipeLink) 
  out, err := curl.Output()
  if err != nil { fmt.Println("curl output error" , err) }

  lines := strings.Split(string(out), "\n")
  for _, line := range lines {
    if strings.Contains(line, `class="recipe-card__image-link"`) {
      line = strings.Split(line, `"`)[1]
      linx = append(linx, line)
    }
  }
  for _, line := range linx {
    fmt.Println(line)
  }

  return linx
}
