## Overview

Companion code for the Gopher Academy blog post:

http://www.gopheracademy.com/writing-a-distributed-system-library


## Running

To run the code, simply `go get` it:

```sh
$ go get https://github.com/benbjohnson/writing-a-distributed-systems-library
```

And then you can run the executable:

```
$ writing-a-distributed-systems-library
```

From another terminal window, you can add values to your system using something fancy like curl:

```
$ curl http://localhost:8000/add?value=5
5

$ curl http://localhost:8000/add?value=10
15

$ curl http://localhost:8000/add?value=-2
13
```

You can stop and restart the program (from the same directory) and your data will be persisted.

Cluster configuration is not included in this simple example.
For a more complete example, please check out the [raftd](https://github.com/goraft/raftd) reference implementation.
