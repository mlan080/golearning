# Training 3: Testing in Go

This session is about writing unit tests with Go.

I’ve started writing up more explanation now that we’re going beyond the basics of Go programming.
For each of the topics explained below, there’s an example in [examples/](examples/), plus there’s
an exercise in [exercise/](exercise/).

I copied some of the examples from a session on testing we did back in March.


## Unit tests

Unit testing means testing the smallest “units” of code that we can test: often a single function or
the methods for a single type. Go has built-in support for unit tests. There’s two parts to this:

* The [testing](https://golang.org/pkg/testing/) package in the standard library.
* The [`go test`](https://golang.org/cmd/go/#hdr-Testing_flags) command, which is included in the Go
  tooling.

Usually, for a Go source file `something.go`, you’ll have a file `something_test.go` with unit
tests.

1. To see an example, take a look at [examples/reverse.go](examples/reverse.go) and
  [examples/reverse_test.go](examples/reverse_test.go).
2. Run the test on the command line: go into the `examples` directory and run `go test -run TestReverse`.
3. Run TestReverse directly in your editor/IDE. If you don’t know how to do that, find out!


## Table-driven tests

A good way to test different inputs to the same function is to put the inputs in a slice, then use a
`for` loop to test each of them (described in [*How to Write Go
Code*](https://golang.org/doc/code.html#Testing)).

1. Take a look at this example: [examples/palindrome.go](examples/palindrome.go) and
   [examples/palindrome_test.go](examples/palindrome_test.go).
2. Add another test case: `トマト` is a palindrome, so we should get `true`.
3. Comment out the first line in `Palindrome` (`word = strings.ToLower(word)`), then run the test
   again -- it should fail this time -- and uncomment the line again. This may seem silly, but it’s
   often useful to break your own code to make sure your test does really test it.


## Using testify/assert

The [testify](https://github.com/stretchr/testify/) library is an open-source project that offers
some functionality to help write tests. Its `assert` package lets you write tests with a little less
code.

1. Check out [examples/palindrome_assert_test.go](examples/palindrome_assert_test.go). It’s almost
   the same as the previous example, except the “if” statement is replaced by a call to
   `assert.Equal`.
2. Comment out the first line in `Palindrome` again, then run both `TestPalindrome` and
   `TestAssertPalindrome`. Compare the two error messages -- the one from `TestPalindrome` tells you
   exactly what went wrong, while the one from `TestAssertPalindrome` is a bit generic.

This is generally the trade-off with the `assert` package: your test code becomes more concise,
especially when you have a lot of checks, but the messages you get for test failures may be less
helpful. ([Subtests](https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks) can be used to
improve the messages, but that’s a bit beyond this exercise.)

There’s more functions than just `assert.Equal`. See [this
introdution](https://github.com/stretchr/testify/#assert-package) or [the reference
documentation](https://godoc.org/github.com/stretchr/testify/assert) for the others.


## Providing files for tests

If your code reads a file from the file system, your test somehow needs to provide that file. In Go,
any files that are only needed for tests are put in a directory called `testdata`.

Example: [examples/translations.go](examples/translations.go) and
[examples/translations_test.go](examples/translations_test.go).


## Using httptest

Code that makes HTTP requests to some API can be hard to unit-test. For example, our Kobold service
uses exchangeratesapi.io to get currency conversion rates. We want to test that code, but we don’t
want it to actually make a request to exchangeratesapi.io during test -- unit tests should *just*
test our code, not rely on some third-party service.

The solution is that during the test you run a web server that works like the API, but with
hard-coded reponses. Then you pass the URL to that server to your code, so it’ll use this “fake”
version of the API instead of the real one. The
[httptest](https://golang.org/pkg/net/http/httptest/) package helps us do that.

1. Take a look at [examples/rates.go](examples/rates.go) for the code we want to test.
2. In [examples/rates_test.go](examples/rates_test.go), look at `TestConversionRate1`. It looks like
   a straightforward test, but it depends on the real exchangeratesapi.io service. If you run it,
   it’ll probably fail because the conversion rate changed since I wrote the test.
3. Take a look at `TestConversionRate2` for a better test. It has a sample response as a string,
   and it uses httptest to set up a server that always returns that response.

The JSON response here is just a few lines, so I put it directly in the source code. When the
reponse is larger, it’s common to put it in a file in `testdata` instead.


## Writing fakes

In Go we often use interfaces to decouple different packages. This enables two techniques for
testing called “fakes” and “mocks.”

It’s easiest to explain with an example. Take a loot at [examples/convert.go](examples/convert.go).
It defines a function that uses the `ExchangeRates` type from the previous example, so we could test
it by again using httptest, creating an `ExchangeRates` object, and using that to test the
function... but we don’t need to.

Instead, in our test code, we’ll define a new type that implements the `Rates` interface. This is
what I’ve done in [examples/convert_fake_test.go](examples/convert_fake_test.go): `FakeRates` always
returns `2` as the rate, which makes it easy to write the actual test.


## Mocking with testify/mock

A mock is similar to a fake, but it lets you do one more thing: it lets you check, in your test,
which methods were called on the mock object. Mocks are usually implemented with the help of a
library. The [testify/mock](https://github.com/stretchr/testify/#mock-package) library is the one
we’re using.

I’ll explain it with an example again. Take a look at
[examples/convert_mock_test.go](examples/convert_mock_test.go). It also tests the `Convert`
function, and it also defines an implementation of `Rates` to be used in the test. However, here the
the `mock.Mock` type is “embedded” in the new struct:

    type MockRates struct {
            mock.Mock
    }

This means you can call methods defined on `mock.Mock` directly on `MockRates`. This is used in two
places:

* The test function uses the `On` method to set up expectations: `rates.On("Get", "EUR",
  "USD").Return(nil, float64(2))` means “I expect a call to `Get` with arguments `"EUR"` and
  `"USD"`, and it should return `nil` and `2`.”
* The mock implementation of `Get` calls the mock.Mock methods to tell it that it was called and
  which arguments it got and to find out what it should return.

To see it in action, go into `convert.go` and change `r.Get(from, to)` to `r.Get(to, from)`. Run
`TestConvertWithFake` and `TestConvertWithMock` -- the one using a fake won’t catch this bug,
because it doesn’t check the method arguments, but the one using a mock will.


## Fakes vs mocks

So if mocks can catch bugs that a fake wouldn’t catch, does that mean mocks are simply better? Well,
it’s a trade-off.

One the one hand, mock objects let you test some things that would be difficult to test with fakes.
Mocks are also very flexible: once you’ve defined a mock object with mock.Mock, you can call the
`On` method any number of times to set different expectations and return values.

On the other hand, mocking code is often a bit hard to understand. The testify/mock library makes
heavy use of some advanced language features (so do other mocking libraries) which can make the code
look a bit like magic.

Mocking can also make it easy to write tests that rely on implementation details of the code under
test because they check exactly which methods are called, with which arguments, in which order. This
kind of test can then break when you refactor the code under test or make other small changes, even
if the code is still correct.


## Exercise

The [exercise/](exercise/) directory contains an exercise that ties together httptest, fakes and
mocks.
