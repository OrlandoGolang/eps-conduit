# eps-conduit
HTTP/HTTPS Load Balancer written in GO

## Installation - Vagrant
The easiest way to have a local development installation of eps-conduit is to use vagrant.

### Dependencies

- Virtualbox
- Vagrant

### Installation

#### 1. Clone the repo

`$ git clone git@github.com:orlandogolang/eps-conduit.git`

#### 2. Run Vagrant Commands

`$ vagrant up`

These vagrant commands should create everything you need to develop and test eps-conduit. In particular, a lightweight "hello world" web application is running in the backgorund on ports 8001 and 8002. It is managed with [supervisor](http://supervisord.org/).

#### 3. SSH in to Vagrant Machine and Run eps-conduit

`$ vagrant ssh`

`$ cd /vagrant/bin`

`$ sudo ./eps-conduit`

**Note**

	You must run eps-conduit as sudo to use the default logpath. Otherwise, specify a logpath
	accessible by a non-root user.

After running these commands, eps-conduit should be running in your terminal session.

#### 4. Verify Installation From Host Machine

Open another terminal window on your host machine, and run the following command:

`$ curl localhost:8000`

You should see a message similar to the following:

	Hello world, I'm answering your request from backend port 8001


Subsequent `curl` commands should yield responses from alternating ports (8001 & 8002) which confirms the load balancer is distributing requests, via round-robin, across the two available "hello world" processes

***

## Usage

By default, eps-conduit will bind to port 8000, but any port can be specified.

### Flags

| Flag  | Description                                       | Example Usage                             |
| ----- | ------------------------------------------------- | ----------------------------------------- |
| -b    | specify a list of backend hosts                   | eps-conduit -b "10.2.8.1, 10.2.8.2"       |
| -bind | specify what port to bind to (defaults to 8000)   | eps-conduit -bind 80                      |
| -mode | specifies what mode to use (http or https)        | eps-conduit -mode https                   |
| -cert | specify an SSL cert file (for https mode)         | eps-conduit -cert mycert.crt              |
| -key  | specify an SSL keyfile (for https mode)           | eps-conduit -key mykeyfile.key            |
| -log  | specify a path to store access logs               | eps-conduit -log /var/log/eps-conduit.log |
