# eps-conduit
HTTP/HTTPS Load Balancer written in GO

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

When you finished developing your fix or enhancement, issue a pull request against the `develop` branch. In your pull request, be sure to reference the corresponding issue via #issue-number (ex. \#24).


Then, be on the lookout for feedback from the contributors, or a merge complete notification!
