First build the program by typing the following command.
```$xslt
    go build main.go
    go run main.go
```

Then run the program by using following command.
```$xslt
     ./main google.com facebook.com hello.com http://golang.org
```

You can use flag "parallel" to set the limit of processing the requests otherwise it is set to 10 by default.
```$xslt
   ./main -parallel 3 google.com facebook.com hello.com http://golang.org
```