package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

func main() {
  getFirstPage()
}

func getLinx() {
  linx := []string{}
  outFile, err := os.Create("output.txt")
  if err != nil {
    fmt.Println(" file open error" , err)
  }

  for i := 2; i < 181; i++ {
    recipeLink := "https://www.177milkstreet.com/recipes/p" + fmt.Sprint(i)
    fmt.Println(i)
    fmt.Println(recipeLink)
    curl := exec.Command("curl", recipeLink) 
    out, err := curl.Output()
    if err != nil {
      fmt.Println("curl output error" , err)
    }

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

func getFirstPage() {
  linx := []string{}

  recipeLink := "https://www.177milkstreet.com/recipes"
  curl := exec.Command("curl", recipeLink) 
  out, err := curl.Output()
  if err != nil {
    fmt.Println("curl output error" , err)
  }

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
}
