## Lets Go is a book to learn Go programming language by building some projects 


<img src="./letsGoBook.png" alt="lets go book logo" width="200"/>



## Route Patterns
##### When a route pattern ends with a trailing slash — like "/" or "/static/" — it is known as a
##### subtree path pattern. Subtree path patterns are matched (and the corresponding handler
##### called) whenever the start of a request URL path matches the subtree path. If it helps your
##### understanding, you can think of subtree paths as acting a bit like they have a wildcard at
##### the end, like "/**" or "/static/**" .
##### This helps explain why the "/" route pattern acts like a catch-all. The pattern essentially
##### means match a single slash, followed by anything (or nothing at all) .
##### Restricting subtree paths

##  Restricting subtree paths
##### To prevent subtree path patterns from acting like they have a wildcard at the end, you can
##### append the special character sequence {$} to the end of the pattern — like "/{$}" or
##### "/static/{$}" .

### test your api from terminal using curl command

```
```
```sh

curl -i  "" localhost:8080/

curl: (3) URL rejected: Malformed input to a URL function
HTTP/1.1 200 OK
Date: Sat, 31 May 2025 22:07:37 GMT
Content-Length: 25
Content-Type: text/plain; charset=utf-8

Welcome to the Home Page
               
curl -i  "" localhost:8080/item/view/100

curl: (3) URL rejected: Malformed input to a URL function
HTTP/1.1 200 OK
Date: Sat, 31 May 2025 22:07:48 GMT
Content-Length: 16
Content-Type: text/plain; charset=utf-8

Viewing Item 100

```






