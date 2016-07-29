# luminati

A go library to use luminati.io proxy service. Personal project just to do something meaningful while trying to learn
a little bit of golang

[![CircleCI](https://circleci.com/gh/Elph/luminati/tree/master.svg?style=svg&circle-token=42e3139e3ff89bcfbbe4f62d31a30cb8753340db)](https://circleci.com/gh/Elph/luminati/tree/master) master branch

[![CircleCI](https://circleci.com/gh/Elph/luminati/tree/dev.svg?style=svg&circle-token=42e3139e3ff89bcfbbe4f62d31a30cb8753340db)](https://circleci.com/gh/Elph/luminati/tree/dev) dev branch

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisities

Just import like any other go library

```
import "github.com/elph/luminati"
```

### Using it

``` 
c := luminati.NewClient("luminati-user", "luminati-pwd", "luminati-host", 989898)
request, _ := http.NewRequest("GET", "http://www.google.com", nil)
resp, err := c.Do(request)
```

## Running the tests

```
go test
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
