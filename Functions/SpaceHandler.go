package reloaded

import (
  "strings"
)

func SpaceHandler(input string) string {
  words := strings.Fields(input)
  var result string
  for _, word := range words {
    if word != "" {
      if result != "" {
        result += " "
      }
      result += word
    }
  }
  return result
}