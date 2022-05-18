goiio
======

goiio is a [Go](http://golang.org/) client to interact with IIO (industrial I/O).

Introduction
-------

go-iio was written to access remote iiod instanced, since libiio does not seem to have golang bindings by itself. It's main use case is connecting with remote IIOd instances: local IIO access is out of scope, you'll need to run iiod.

Example
-------------


    // Create TCP connection
    c, err := goiio.New("my-sensor.home.lan:30431")
    if err != nil {
      panic(err)
    }

    // Populate the values in attributes
    if err = c.FetchAttributes(); err != nil {
      panic(err)
    }

    // Loop through all devices
    for _, dev := range c.Context.Devices {
      log.Infof("Device: id=%s, name=%s", dev.ID, dev.Name)
      for _, ch := range dev.Channels {
        log.Infof("  Channel: id=%s", ch.ID)
        for _, attr := range ch.Attributes {
          log.Infof("    Attribute: %s, value: %0.3f", attr.Name, attr.Value)
        }
      }
    }

Installation
------------

Install goiio using the "go get" command:

    go get github.com/brendan-ta/goiio


Minimum Golang dependency: 1.16

Dependencies
------------

- [logrus](https://pkg.go.dev/github.com/sirupsen/logrus) - Structured logger

Standard Go library imports:

- [bufio](https://pkg.go.dev/bufio) - buffered I/O
- [encoding/xml](https://pkg.go.dev/encoding/xml) - XML parser
- [fmt](https://pkg.go.dev/fmt) - formatted I/O
- [net](https://pkg.go.dev/net) - network I/O
- [strconv](https://pkg.go.dev/strconv) - to and from string conversion
- [strings](https://pkg.go.dev/strings) - maipulate UTF-8 encoded strings

Related Projects
----------------

- [libiio](https://github.com/analogdevicesinc/libiio) - libiio


Resources
--------------

This project was forked from the original project [goiio](https://pkg.go.dev/github.com/jonkerj/goiio) however due
to go modules forking a project meant collisions with module names.  Therefore this is simply an evolution of
the original project:

- [goiio](https://github.com/jonkerj/goiio) - Original goiio

License
--------

goiio is available under the [GNU AGPL Licence](https://www.gnu.org/licenses/agpl-3.0.en.html).
