#!/bin/bash

read -p "Enter directory name: " dir_name

mkdir "$dir_name"

cd "$dir_name"

cat <<EOF > main.go
package main

/**
*/

func main() {
}
EOF

go mod init "github.com/shylabo/use-of-go/$dir_name"

cd ..

go work use "./$dir_name"
