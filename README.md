# imgtobase64

Small Go package for converting image to base64 string

## Installation

You can install this package by using go modules

```
go get github.com/kirillshevch/imgtobase64
```

## Usage

```go
package main

func main() {
  base64, err := Imgtobase64("/path/to/image")

  if err != nil {
    fmt.Print(err)
  }

  fmt.Print(base64)
}
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/kirillshevch/imgtobase64.

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
