# Load Test Web App with Golang
This is web app designed to help you to benchmark your projects.
It allows you to sit and relax while your service is running on Heroku,OpenShit, AppEngine, etc.

[AngularJS](https://angularjs.org/), [Bower](http://bower.io/), [Firebase](http://firebaseio.com), [Golang](golang.org) and [Vegeta](https://github.com/tsenart/vegeta)

It nicely wraps Vegeta HTTP load testing with Golang net/http package, firebase and AngularJS.

Beutify is in progress, so do not expect to find awesome graphics and metrics just yet.


Use it via the webapp: http://localhost:3000/

Or

cUrl it manually:
```
curl 'http://localhost:3000/benchtest' -H 'Accept: application/json, text/plain, */*' -H 'Referer: http://localhost:3000/' -H 'Origin: http://localhost:3000' -H 'X-Requested-With: XMLHttpRequest' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.118 Safari/537.36' -H 'Content-Type: application/json;charset=UTF-8' --data-binary '{"name":"google","site":"http://www.google.com","method":"GET","body":"","duration":"4"}' --compressed
```

returns back:
```
{"latencies":{"mean":468452083,"50th":469447624,"95th":597150580,"99th":633816447,"max":690458106},"bytes_in":{"total":1041072,"mean":2602.68},"bytes_out":{"total":0,"mean":0},"duration":3980000000,"wait":554108796,"requests":400,"rate":100.50251256281408,"success":0,"status_codes":{"503":400},"errors":["503 Service Unavailable"]}
```

#AngularJS
![AngularJS](http://upload.wikimedia.org/wikipedia/commons/thumb/c/ca/AngularJS_logo.svg/250px-AngularJS_logo.svg.png)

AngularJS, commonly referred to as Angular, is an open-source web application framework maintained by Google and a community of individual developers ...
[Wikipedia](http://en.wikipedia.org/wiki/AngularJS)

#Firebase
![Firebase](http://upload.wikimedia.org/wikipedia/en/b/bd/Firebase_Logo.png)

Firebase is a cloud services provider and backend as a service company based in San Francisco, California. The company makes a number of products for software developers building mobile or web applications.
[Wikipedia](http://en.wikipedia.org/wiki/Firebase)

#Golang
![Golang](http://upload.wikimedia.org/wikipedia/commons/thumb/2/23/Golang.png/300px-Golang.png)

Go, also commonly referred to as golang, is a programming language initially developed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson.
[Wikipedia](http://en.wikipedia.org/wiki/Go_%28programming_language%29)


# Vegeta
![Vegeta](http://fc09.deviantart.net/fs49/i/2009/198/c/c/ssj2_vegeta_by_trunks24.jpg)

Vegeta is a versatile HTTP load testing tool built out of a need to drill
HTTP services with a constant request rate. It's over 9000!

## Install
You need go installed and `GOPATH` must be configured. Once that is done, run the
command:
```shell
$ go get github.com/tsenart/vegeta
$ go install github.com/tsenart/vegeta
```

Install css/js dependencies, require [Bower](http://bower.io/)
```shell
cd assets
bower install
```
Outputs:
```
bower angularfire#~0.2          cached git://github.com/firebase/angularFire.git#0.2.0
bower angularfire#~0.2        validate 0.2.0 against git://github.com/firebase/angularFire.git#~0.2
bower bootstrap#~3.3.4          cached git://github.com/twbs/bootstrap.git#3.3.4
bower bootstrap#~3.3.4        validate 3.3.4 against git://github.com/twbs/bootstrap.git#~3.3.4
bower angular-resource#~1.0     cached git://github.com/angular/bower-angular-resource.git#1.0.8
bower angular-resource#~1.0   validate 1.0.8 against git://github.com/angular/bower-angular-resource.git#~1.0
bower angular#~1.0              cached git://github.com/angular/bower-angular.git#1.0.8
bower angular#~1.0            validate 1.0.8 against git://github.com/angular/bower-angular.git#~1.0
bower angular#1.0.7             cached git://github.com/angular/bower-angular.git#1.0.7
bower angular#1.0.7           validate 1.0.7 against git://github.com/angular/bower-angular.git#1.0.7
bower jquery#>= 1.9.1           cached git://github.com/jquery/jquery.git#2.1.3
bower jquery#>= 1.9.1         validate 2.1.3 against git://github.com/jquery/jquery.git#>= 1.9.1

Unable to find a suitable version for angular, please choose one:
    1) angular#1.0.7 which resolved to 1.0.7 and is required by angularfire#0.2.0
    2) angular#~1.0 which resolved to 1.0.8 and is required by assets
    3) angular#1.0.8 which resolved to 1.0.8 and is required by angular-resource#1.0.8

Prefix the choice with ! to persist it to bower.json

? Answer: 1
bower angularfire#~0.2         install angularfire#0.2.0
bower angular-resource#~1.0    install angular-resource#1.0.8
bower bootstrap#~3.3.4         install bootstrap#3.3.4
bower angular#1.0.7            install angular#1.0.7
bower jquery#>= 1.9.1          install jquery#2.1.3

angularfire#0.2.0 bower_components/angularfire
└── angular#1.0.7

angular-resource#1.0.8 bower_components/angular-resource
└── angular#1.0.7

bootstrap#3.3.4 bower_components/bootstrap
└── jquery#2.1.3

angular#1.0.7 bower_components/angular

jquery#2.1.3 bower_components/jquery
```

## Usage
```shell
go run app.go
```

#### Limitations
On a UNIX set soft-limit values.
```shell
$ ulimit -n # file descriptors
60560
$ ulimit -u # processes / threads
850
```
Just pass a new number as the argument to change it.


## Author
Thomas Modeneis
[StackOverflow](https://careers.stackoverflow.com/thomasmodeneis)
[LinkedIn](https://uk.linkedin.com/in/thomasmodeneis)

## License
```
The MIT License (MIT)

Copyright (c) 2015-2015 Gibraltar Golang Community

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
