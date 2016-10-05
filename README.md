url
===
[![Build Status](https://travis-ci.org/Adracus/url.svg?branch=master)](https://travis-ci.org/Adracus/url)

Ever wanted to easily extract the host, port or something
else in a bash script from a URL? Then tried to google it
and found posts like [this](http://stackoverflow.com/questions/6174220/parse-url-in-shell-script)?

This was the reason for me to harness the go built-in url
parsing to write a small command-line helper that enables
me and my bash scripts to extract these values with ease.

### Usage

By just writing

``
url --help
``

The usage information is printed.

At the moment of writing, the tool supports the following subcommands:

* `scheme`: Extracts the scheme of the given url
* `user`: Extracts the user information of the given url. If for instance
  the url is `http://john:doe@example.com`, the printed string will be
  `john:doe`.
* `username`: Extracts _only_ the username of the given url.
* `password`: Extracts _only_ the password of the given url.
* `host`: Extracts the host part of the given url.
* `port`: Extracts the port of the given url if it was specified.
* `path`: Extracts the path of the given url.
* `query`: Extracts the query of the given url.
* `query-param` <name>: Extracts the query parameter with the given key.
* `fragment`: Extracts the fragment of the given url.

## Issues and Enhancements

If you've got issues, please report them via GitHub issues to this repository.
For enhancements or additional work: Pull requests are always welcome, but
please: __WRITE TESTS__.

## License

License is the MIT license - I anyways only used the go url parsing library.

