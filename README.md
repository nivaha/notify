Notify
==

The goal of this project is to create a *simple* server with a small, well-defined API to support a notification system across several applications.

The project is set up to use [gb](github.com/constabulary/gb) for building and vendoring:

```sh
# First make sure you have gb installed
$ go get github.com/constabulary/gb/...

# Now clone the repo and build
$ git clone git@github.com:nivaha/notify.git
$ cd notify
$ gb build
```

You can get help by typing `./bin/notify -h`

Project Status
==
Doesn't really do anything yet - just set up for talking to the db and for handling routing.
