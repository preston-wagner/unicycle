# fetch
This submodule contains functions meant for common types of network requests.

## Highlights:

### fetch.Fetch
The most basic form of a request facilitated by unicycle.
Accepts a url string and an instance of `FetchOptions`, which specifies the request method, body contents, headers, and so on in an easy-to-use struct.
For unprovided values, sane defaults such as `Method = "GET"` and `timeout := time.Minute` are used.

### fetch.FetchJson
Wraps `fetch.Fetch`
