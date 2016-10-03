package main

import (
  "testing"
  "strings"
)

func TestValid(t *testing.T) {
  cases := []struct {
    argString string
    expected string
  }{
    {"scheme http://example.org", "http"},
    {"scheme ftp://example.org", "ftp"},
    {"scheme jdbc://example.org", "jdbc"},
    {"user http://user@example.org", "user"},
    {"user http://john@example.org", "john"},
    {"user http://john:doe@example.org", "john:doe"},
    {"username http://john:doe@example.org", "john"},
    {"username http://foo:bar@example.org", "foo"},
    {"password http://john:doe@example.org", "doe"},
    {"password http://foo:bar@example.org", "bar"},
    {"password http://user@example.org", ""},
    {"host http://example.org", "example.org"},
    {"host ftp://example.org?some=query", "example.org"},
    {"host http://com.example.org", "com.example.org"},
    {"port http://example.com:8080", "8080"},
    {"port http://example.com:8", "8"},
    {"port http://example.com", ""},
    {"path http://example.com/foo/bar", "/foo/bar"},
    {"path http://example.com/baz?foo=bar", "/baz"},
    {"query http://example.com?foo=bar", "foo=bar"},
    {"query http://example.com?foo=bar&baz=qux", "foo=bar&baz=qux"},
    {"query-param foo http://example.com?foo=bar&baz=qux", "bar"},
    {"query-param baz http://example.com?foo=bar&baz=qux", "qux"},
    {"fragment http://example.com#fragment", "fragment"},
    {"fragment http://example.com", ""},
  }

  for _, c := range cases {
    t.Logf("Running main with args %q", c.argString)
    args := strings.Split(c.argString, " ")

    output, _ := Main(args)

    if output != c.expected {
      t.Errorf("url %q = %q, expected %q", c.argString, output, c.expected)
      break
    } else {
      t.Logf("Successfull running with args %q", c.argString)
    }
  }
}

func TestInvalid(t *testing.T) {
  cases := []string {
    "scheme",
    "",
    "query-param http://example.com",
    "host http://example.com foo",
  }

  for _, c := range cases {
    t.Logf("Running main with args %q", c)
    args := strings.Split(c, " ")
    output, err := Main(args)

    if err == nil {
      t.Errorf("Expected %q to error but returned %q", c, output)
      break
    }

    t.Logf("Errored as expected: %q", c)
  }
}

