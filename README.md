tcp2stdout
==========

Listen on a TCP socket and send to stdout.

Why?
====

This was written to handle wanting to write to stdout from inside a container from a PHP
application. However, [php-fpm](https://bugs.php.net/bug.php?id=71880) prefixes each line with a preamble.  So, this allows one to configure PHP loggers - such as [monolog](https://github.com/Seldaek/monolog) - to log
to a tcp socket and still [capture stdout](https://12factor.net/logs).

Usage
=====

```
$ ./tcp2stdout -h
Usage of ./tcp2stdout:
  -addr string
    	listen address (default "127.0.0.1:1313")
  -buffer int
    	number of lines to buffer (default 512)
exit status 2
```

An example using monolog to log to this socket:

```php
$logger = new Monolog\Logger("example");
$handler = new Monolog\Handler\SocketHandler('tcp://127.0.0.1:1313');
$formatter = new Monolog\Formatter\JsonFormatter();
$handler->setFormatter($formatter);
$logger->pushHandler($handler);
```

Building
========

```
go build .
````

Also availible at https://quay.io/repository/bakins/tcp2stdout
