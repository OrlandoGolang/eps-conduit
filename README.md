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

`$ vagrant provision`

These vagrant commands should create everything you need to develop and test eps-conduit. In particular, a lightweight "hello world" web application is running in the backgorund on ports 8001 and 8002. It is managed with [supervisor](http://supervisord.org/).

#### 3. SSH in to Vagrant Machine and Run eps-conduit

`$ vagrant ssh`

`$ cd /vagrant/bin`

`$ ./eps-conduit`

After running these commands, eps-conduit should be running in your terminal session.

#### 4. Verify Installation From Host Machine

Open another terminal window on your host machine, and run the following command:

`$ curl localhost:8000`

You should see a message similar to the following:

	Hello world, I'm answering your request from backend port 8001


Subsequent `curl` commands should yield responses from alternating ports (8001 & 8002) which confirms the load balancer is distributing requests, via round-robin, across the two available "hello world" processes

***

##Usage
By default, eps-conduit will bind to port 8000, but any port can be specified.

###Flags
* -b    list of backend hosts
  * ex:  eps-conduit -b "10.2.8.1, 10.2.8.2"
* -bind specify what port to bind to (defaults to 8000)
  * ex:  eps-conduit -bind 80
* -mode specifies what mode to use (http or https)
  * ex:  eps-conduit -mode https
* -cert specify an SSL cert file (for https mode)
  * ex:  eps-conduit -cert mycert.crt
* -key  specify an SSL keyfile (for https mode)
  * ex:  eps-conduit -key mykeyfile.key

## Contributing

### 1. Create an Issue / Proposal

Create a github issue detailing the bug or enhancement you would like to work on.

The project admins will evaluate your proposal and flag as `approved` or `not approved`

**Note**

	Pull requests that do not have a corresponding, approved issue will be rejected.

### 2. Fork, Develop, and Issue Pull Request

When you finished developing your fix or enhancement, issue a pull request against the `develop` branch of the base fork. In your pull request, be sure to reference the corresponding issue via #issue-number (ex. \#24).


Then, be on the lookout for feedback from the contributors, or a merge complete notification!
