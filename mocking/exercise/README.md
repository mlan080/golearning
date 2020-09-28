# HTTP Downloader Exercise

This exercise is a bit backward: the implementation is already complete and your task is to write
tests.

The `downloader.go` file defines a Downloader type that downloads files over HTTP and uses some kind
of cache. The cache could, for example, be implemented with Redis. We want to test the Downloader
without a “real” Cache implementation and without making HTTP requests to some third-party website.

We’ll do that in three steps.


## Step 1: Test the case where a URL is already cached

To get started, write a fake implementation of the Cache interface that returns a fixed string as
the result. Use it to implement TestDownloadCached.

We don’t need httptest here -- not yet.

You can use testify/assert or not, whichever style you prefer.


## Step 2: Test the case where the URL is not cached

Let’s test the case where it actually downloads the URL.

Change your FakeCache implementation so it returns `false` for this test. You could add a `bool`
field to the struct that determines what it does, or make it return different values depending on
the URL. 

In TestDownloadNotCached, set up httptest and test the actual downloading.


## Step 3: Make sure the cache is updated

There’s one part of the function that we haven’t tested so far: updating the cache after downloading
a file. We can do that with mocking.

Write a MockCache using the testify/mock library.

In TestDownloadCacheUpdate, use httptest and MockCache to write the test. Make sure you call the
`AssertExpectations` method at the end; that’s where the library checks that the methods have been
called as expected.
