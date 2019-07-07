package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Todd McLeod"
	str := fmt.Sprint(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hello World!</title>
</head>
<body>
<h1>` +
		name +
		`</h1>
</body>
</html>
	`)

	nf, err := os.Create("index.html")

	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	// Using Copy form io package
	io.Copy(nf, strings.NewReader(str))

	// Using WriteString from bufio package
	w := bufio.NewWriter(nf)
	w.WriteString(str)

}