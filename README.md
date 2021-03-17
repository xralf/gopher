# Notes

## Greetings library and Hello main program

    cd greetings
    go mod init github.com/xralf/gopher/greetings


    cd greetings
    git init -q
    git remote add origin https://github.com/xralf/greetings.git
    git add greetings.go go.mod
    git commit -q -m 'Initial commit'
    git branch -M main
    git push -u origin main


    mkdir hello
    cd hello
    go get github.com/xralf/greetings
    # go: downloading github.com/xralf/greetings v0.0.0-20210317233144-88626fb63681
    go mod init github.com/xralf/greetings

    go list -m -f {{.Version}} github.com/xralf/greetings

    cd ../hello
    go mod init github.com/xralf/hello
    touch hello.go
    go build
    ./hello