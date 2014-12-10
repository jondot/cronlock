# Cronlock

A tiny utility to help make sure only one cron job (or process) of the same kind is running at any given time.

### The problem

As an example, if you have a cron job that depends on a database and hangs when the database is down, next cron jobs will hang as well when they're launched.
This can possibly snowball into consuming all of the available resources.

### The solution

A good design is to have a circuitbreaker in place.

In our case - fail the next cron job when an existing cron job is already running.

## Quick start

Download a binary release here:

[https://github.com/jondot/cronlock/releases/tag/1.0.2](https://github.com/jondot/cronlock/releases/tag/1.0.2)


## Usage

Cronlock is Unixy in that it does one thing and one thing well. Cronlock will

* Fail (exit(1)) when it cannot acquire an exclusive lock
* Succeed (exit(0)) when it acquired an exclusive lock

You can exploit these two properties in order to make sure your Cron job
doesn't step on itself:

```
$ cronlock your required command && signal success command
```

you can also define your own lock with the `CRONLOCK_LOCKFILE` environment variable:

```
$ CRONLOCK_LOCKFILE=foobar.lock cronlock your required command && your
success command
```

Cronlock is build around `flock` internally and uses just the `syscall` package to stay lean and mean, which should be stable and reliable for this use case. That's basically it.


## Usage (Extended)

You got me. Cronlock isn't only useful for cron, it could have just been
called `flock`.

You can use it anywhere you want to make sure just one unix process is
running, in the same way.




# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2014 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.






