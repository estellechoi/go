# Go

This repo is for getting to know Go.

<br />

### Install Go

```zsh
# installing path would be /usr/local/Cellar/go ($GOROOT)
# workspace would be /Users/username/go ($GOPATH) by default
# just in case you install it using brew
brew install go
```

<br />

### Go to your workspace defined by `$GOPATH` and make dir named `src`

```zsh
cd /Users/username/go
mkdir src
cd src
```

### Clone this repo under your `$GOPATH`

```zsh
# make path /github.com/estellechoi under the $GOPATH to use go command
mkdir -p github.com/estellechoi
cd github.com/estellechoi

# git clone
git clone url
cd go
```

<br />

### Install packages

```zsh
# if GO111MODULE=off, use $GOPATH and vendor/ path when installing modules
# if GO111MODULE=on, ignore $GOPATH
go env -w GO111MODULE=off
go get
```

<br />

### Run main package

```zsh
go run main.go
```

<br />

---

### References

- [Tutorial: Get started with Go | Go](https://go.dev/doc/tutorial/getting-started)
