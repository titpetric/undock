# undock

Install servers by using a Dockerfile

## Requirements

The Dockerfile which you want to use as the installation script on your
server should be based on the same OS base image which you have installed.

For example:

- OK: `FROM debian:stretch` for a Debian Stretch install,
- POSSIBLY OK: `FROM debian:jessie` for a Debian Stretch install,
- FAILURE: `FROM golang:1.8` or other custom images,
- FAILURE: `FROM alpine:latest` on a Debian Stretch install

So, there's not a lot of mixing and matching that's enabled here. Basically
what the project does is executes all the `RUN` commands that are listed
in the Dockerfile, as if you were building a Docker image but instead runs
them on a supposedly clean install of your OS.

## Caveats

Aside from the requirements, there are several parts of the Dockerfile which
are just not supported, at all:

- `FROM` (parent images and multi-stage builds),
- `ENTRYPOINT` and `CMD`
- `EXPOSE`
- `HEALTHCHECK`

and pretty much anything except `RUN`. If you think that some additional
coverage can be made (ie, converting a `HEALTHCHECK` into a cron job), I
will be fielding PR's to improve functionality.
