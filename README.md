# Go GRPC Server Quick Start Guide

This guide will walk you through deploying a Go [grpc][grpc] server on [Deis Workflow][]. The server is made non-routable outside of the cluster and a Go [grpc client][grpcclient] needs to be installed to talk with the server.

## Usage

```console
$ git clone https://github.com/deis/example-grpc-server.git
$ cd example-grpc-server
$ deis create
Creating Application... done, created finest-rabbitry
Git remote deis added for app finest-rabbitry
$ deis routing:disable
Disabling routing for finest-rabbitry... done
$ git push deis master
Counting objects: 11, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (8/8), done.
Writing objects: 100% (11/11), 3.60 KiB | 0 bytes/s, done.
Total 11 (delta 0), reused 0 (delta 0)
Starting build... but first, coffee!
-----> Restoring cache...
       No cache file found. If this is the first deploy, it will be created now.
-----> Go app detected
        !!    
        !!    'GOVERSION' isn't set, defaulting to 'go1.7.1'
        !!    
        !!    Run 'heroku config:set GOVERSION=goX.Y' to set the Go version to use
        !!    for future builds
        !!    
-----> Installing go1.7.1... done
-----> Installing glide v0.12.2... done
-----> Installing hg 3.9... done
        !!    Installing package '.' (default)
        !!    
        !!    To install a different package spec for the next build run:
        !!    
        !!    'heroku config:set GO_INSTALL_PACKAGE_SPEC="<pkg spec>"'
        !!    
        !!    For more details see: https://devcenter.heroku.com/articles/go-dependencies-via-glide
        !!    
-----> Fetching any unsaved dependencies (glide install)
       [INFO]	Downloading dependencies. Please wait...
       [INFO]	--> Fetching github.com/golang/protobuf.
       [INFO]	--> Fetching google.golang.org/grpc.
       [INFO]	--> Fetching golang.org/x/net.
       [INFO]	Setting references.
       [INFO]	--> Setting version for github.com/golang/protobuf to 8616e8ee5e20a1704615e6c8d7afcdac06087a67.
       [INFO]	--> Setting version for golang.org/x/net to fb93926129b8ec0056f2f458b1f519654814edf0.
       [INFO]	--> Setting version for google.golang.org/grpc to b7f1379d3cbbbeb2ca3405852012e237aa05459e.
       [INFO]	Exporting resolved dependencies...
       [INFO]	--> Exporting github.com/golang/protobuf
       [INFO]	--> Exporting golang.org/x/net
       [INFO]	--> Exporting google.golang.org/grpc
       [INFO]	Replacing existing vendor dependencies
-----> Running: go install -v -tags heroku .
       github.com/deis/example-grpc-server/vendor/github.com/golang/protobuf/proto
       github.com/deis/example-grpc-server/vendor/golang.org/x/net/context
       github.com/deis/example-grpc-server/vendor/golang.org/x/net/http2/hpack
       github.com/deis/example-grpc-server/vendor/golang.org/x/net/http2
       github.com/deis/example-grpc-server/vendor/golang.org/x/net/internal/timeseries
       github.com/deis/example-grpc-server/vendor/golang.org/x/net/trace
       github.com/deis/example-grpc-server/vendor/google.golang.org/grpc/codes
       github.com/deis/example-grpc-server/vendor/google.golang.org/grpc/credentials
       github.com/deis/example-grpc-server/vendor/google.golang.org/grpc/grpclog
       github.com/deis/example-grpc-server/vendor/google.golang.org/grpc/internal
       github.com/deis/example-grpc-server/vendor/google.golang.org/grpc/metadata
       github.com/deis/example-grpc-server/vendor/google.golang.org/grpc/naming
       github.com/deis/example-grpc-server/vendor/google.golang.org/grpc/peer
       github.com/deis/example-grpc-server/vendor/google.golang.org/grpc/transport
       github.com/deis/example-grpc-server/vendor/google.golang.org/grpc
       github.com/deis/example-grpc-server/_proto
       github.com/deis/example-grpc-server
-----> Discovering process types
       Procfile declares types -> web
-----> Checking for changes inside the cache directory...
       Files inside cache folder changed, uploading new cache...
       Done: Uploaded cache (85M)
-----> Compiled slug size is 2.9M
Build complete.
Launching App...
Done, finest-rabbitry:v2 deployed to Workflow

Use 'deis open' to view this application in your browser

To learn more, use 'deis help' or visit https://deis.com/

To ssh://git@deis-builder.deis.rocks:2222/finest-rabbitry.git
 * [new branch]      master -> master
```

## Additional Resources

* [GitHub Project](https://github.com/deis/workflow)
* [Documentation](https://deis.com/docs/workflow/)
* [Blog](https://deis.com/blog/)

[Deis Workflow]: https://github.com/deis/workflow#readme
[grpc]: http://www.grpc.io/docs/quickstart/go.html
[grpcclient]: https://github.com/deis/example-grpc-client
