# eps-conduit Recipes

Vagrant installation of eps-conduit is somewhat complicated, mostly due to the following requirements:

- Run latest version of golang (1.5.1)
- Configure a web app, running on multiple ports, behind eps-conduit.

## Golang

Installs Golang 1.5.1. This version is currently not available via linux package distributions, so this recipe downloads binaries and configures environment variables.

## Hello World

Downloads and installs [eps-conduit-hello](https://github.com/OrlandoGolang/eps-conduit-hello) - a lightweight web application that reports the port number it's listening on.

This information is helpful for debugging eps-conduit to confirm if it's distributing load properly across multiple backends.

## Build

Downloads dependencies via `$ go get` and compiles eps-conduit via `$ go install` each time vagrant machine is provisioned.

## Supervisor

Installs [supervisor](http://supervisord.org/) and is configured to run and manage `eps-conduit-hello` on multiple ports.

Supervisor will automatically start when the vagrant machine is started or provisioned.

If you need to modify supervisor or some of the process it manages, consult supervisor's [documentation](http://supervisord.org)
