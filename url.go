package main

import (
  "net/url"
  "os"
  "fmt"
  "strings"
  "errors"
)

const usage =
  "url - A command line utility for extracting parts of urls.\n\n" +
  "url SUBCOMMAND-AND-ARGS urlString\n" +
  "\tSUBCOMMAND-AND-ARGS:\n" +
  "\t\tscheme: Returns the scheme of the given url.\n\n" +
  "\t\tuser: Returns the user part of the given url.\n\n" +
  "\t\tusername: Returns the username of the user part of the given url.\n\n" +
  "\t\tpassword: Returns the password of the user part of the given url.\n\n" +
  "\t\thost: Returns the host of the given url.\n\n" +
  "\t\tport: Returns the port of the given url.\n\n" +
  "\t\tpath: Returns the path of the given url.\n\n" +
  "\t\tquery: Returns the query of the given url.\n\n" +
  "\t\tquery-param paramName: Returns the value of the specified query parameter.\n\n" +
  "\t\tfragment: Returns the fragment of the given query."

func urlSubCommand(fn func(*url.URL)(string, error))func([]string)(string, error) {
  return func(args []string)(string, error) {
    if len(args) != 1 {
      return "", errors.New(appendUsage("No url specified."))
    }
    parsedUrl, err := url.Parse(args[0])
    if err != nil {
      return "", err
    }
    return fn(parsedUrl)
  }
}

var subCommands = map[string]func([]string)(string, error){
  "scheme": urlSubCommand(func(u *url.URL)(string, error) { return u.Scheme, nil }),
  "user": urlSubCommand(func(u *url.URL)(string, error) { return u.User.String(), nil }),
  "username": urlSubCommand(func(u *url.URL)(string, error) { return u.User.Username(), nil }),
  "password": urlSubCommand(func(u *url.URL)(string, error) {
    password, _ := u.User.Password()
    return password, nil
  }),
  "host": urlSubCommand(func(u *url.URL)(string, error) { return u.Host, nil }),
  "port": urlSubCommand(func(u *url.URL)(string, error) {
    host := u.Host
    idx := strings.Index(host, ":")
    if idx != -1 && len(host) - 1 > idx {
      return host[(idx + 1):], nil
    } else {
      return "", nil
    }
  }),
  "path": urlSubCommand(func(u *url.URL)(string, error) { return u.Path, nil }),
  "query": urlSubCommand(func(u *url.URL)(string, error) { return u.RawQuery, nil }),
  "query-param": func(args []string)(string, error) {
    if len(args) != 2 {
      return "", errors.New(appendUsage("No query parameter name specified"))
    } else {
      paramName := args[0]
      return urlSubCommand(func(u *url.URL)(string, error) {
        return strings.Join(u.Query()[paramName], ","), nil
      })(args[1:])
    }
  },
  "fragment": urlSubCommand(func(u *url.URL)(string, error) { return u.Fragment, nil }),
}

func Main(args []string) (string, error) {
  if len(args) < 2 {
    return "", errors.New(appendUsage("Too few arguments given."))
  }

  subCommand, prs := subCommands[args[0]]

  if !prs {
    return "", errors.New(appendUsage("Unknown subcommand: " + args[0]))
  }

  return subCommand(args[1:])
}

func main() {
  args := os.Args[1:]
  result, err := Main(args)
  
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  fmt.Print(result)
}

func appendUsage(s string)string {
  return s + "\n\n" + usage
}

