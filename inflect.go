package inflect

import (
  "fmt"
  "regexp"
  "strings"
)

func Pluralize(str string) string {
  if inflector, ok := LANGUAGES[LANGUAGE]; ok {
    return inflector.Pluralize(str)
  }

  return str
}

func Singularize(str string) string {
  if inflector, ok := LANGUAGES[LANGUAGE]; ok {
    return inflector.Singularize(str)
  }

  return str
}

// Split's a string so that it can be converted to a different casing.
// Splits on underscores, hyphens, spaces and camel casing.
func split(str string) []string {
  return strings.Split(regexp.MustCompile(`-|_|([a-z])([A-Z])`).ReplaceAllString(strings.Trim(str, `-|_| `), `$1 $2`), ` `)
}

// Converts a string to it's upper camel case version.
func UpperCamelCase(str string) string {
  pieces := split(str)

  for index, s := range pieces {
    pieces[index] = fmt.Sprintf(`%v%v`, strings.ToUpper(string(s[0])), strings.ToLower(s[1:]))
  }

  return strings.Join(pieces, ``)
}

// Converts a string to it's lower camel case version.
func LowerCamelCase(str string) string {
  pieces := split(str)

  pieces[0] = strings.ToLower(pieces[0])

  for i := 1; i < len(pieces); i++ {
    pieces[i] = fmt.Sprintf(`%v%v`, strings.ToUpper(string(pieces[i][0])), strings.ToLower(pieces[i][1:]))
  }

  return strings.Join(pieces, ``)
}

// Converts a string to it's underscored version.
func Underscore(str string) string {
  pieces := split(str)

  for index, piece := range pieces {
    pieces[index] = strings.ToLower(piece)
  }

  return strings.Join(pieces, `_`)
}

// Converts a string to it's underscored version.
func Hyphenate(str string) string {
  pieces := split(str)

  for index, piece := range pieces {
    pieces[index] = strings.ToLower(piece)
  }

  return strings.Join(pieces, `-`)
}

// Converts a string to it's constantized version.
func Constantize(str string) string {
  pieces := split(str)

  for index, piece := range pieces {
    pieces[index] = strings.ToUpper(piece)
  }

  return strings.Join(pieces, `_`)
}

// Converts a string to it's humanized version.
func Humanize(str string) string {
  pieces := split(str)

  pieces[0] = fmt.Sprintf(`%v%v`, strings.ToUpper(string(pieces[0][0])), pieces[0][1:])

  for i := 1; i < len(pieces); i++ {
    pieces[i] = fmt.Sprintf(`%v`, strings.ToLower(pieces[i]))
  }

  return strings.Join(pieces, ` `)
}

// Converts a string to it's titleized version.
func Titleize(str string) string {
  pieces := split(str)

  for i := 0; i < len(pieces); i++ {
    pieces[i] = fmt.Sprintf(`%v%v`, strings.ToUpper(string(pieces[i][0])), strings.ToLower(pieces[i][1:]))
  }

  return strings.Join(pieces, ` `)
}
