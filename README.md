# undock

Install servers by using a Dockerfile



## Requirements

You will need `rsync` installed and in your path.

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


## Usage

Navigate yourself to a folder which contains a `Dockerfile`.

~~~
# undock > undock.sh
~~~

The command will generate a bash script for you, which you can run with
`bash undock.sh` in order to run the relevant commands.

## Example output

~~~
#!/bin/bash
# FROM debian:jessie -- not implemented (from)
# ARG DEBIAN_FRONTEND=noninteractive -- not implemented (arg)
# RUN apt-get -qq update && apt-get -qq -y install krb5-user dnsutils curl wget
apt-get -qq update && apt-get -qq -y install krb5-user dnsutils curl wget
# ARG GITVERSION=development -- not implemented (arg)
# ENV GITVERSION ${GITVERSION} -- not implemented (env)
# COPY conf/krb5.conf /etc/krb5.conf
rsync -ia [dockerfile path]conf/krb5.conf /etc/krb5.conf
# ADD conf/run.sh /run.sh
cp [dockerfile path]conf/run.sh /run.sh
# RUN chmod +x /run.sh ; sync; sleep 1
chmod +x /run.sh ; sync; sleep 1
# WORKDIR /app -- not implemented (workdir)
# ENTRYPOINT ["/run.sh"] -- not implemented (entrypoint)
~~~

