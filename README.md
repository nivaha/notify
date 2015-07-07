Notify
==

The goal of this project is to create a *simple* server with a small, well-defined API to support a notification system across several applications.

The project uses go 1.4.x is set up to use [gb](https://github.com/constabulary/gb) for building and vendoring:

```sh
# First make sure you have go installed
$ brew install go
```

Next, make sure that your **GOPATH** environment variable is set in your **.bash_profile** and that its bin directory is added to your **PATH**:

```sh
# add to your .bash_profile
export GOPATH="${HOME}/devdt/go"
export PATH="${GOPATH}/bin:${PATH}"
```

```sh
# Make sure you have gb installed (it will build and end up in ${GOPATH}/bin)
$ go get github.com/constabulary/gb/...

# Now clone the repo and build
$ git clone git@github.com:nivaha/notify.git
$ cd notify
$ gb build
```

You can get help by typing `./bin/notify -h`

Editor
==
If you are using the Atom editor, we recommend using `go-plus` package
```sh
$ apm install go-plus
```

Project Status
==
Doesn't really do anything yet - just set up for talking to the db and for handling routing.
