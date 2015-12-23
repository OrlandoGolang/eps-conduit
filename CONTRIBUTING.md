# Contributing

Thanks a lot for showing interest in making **eps-conduit** better! We have a few guidelines that we expect from contributors to help make triaging code and testing new features easier.

## 1. Create an Issue / Proposal

Create a github issue detailing the bug or enhancement you would like to work on. The project admins will evaluate your proposal and flag as `approved` or `not approved`.

**Note**

	Pull requests that do not have a corresponding, approved issue will be rejected.

Before we can triage your changes, please make sure that you:

- Test that the Vagrant box builds and processes requests based on the [README](https://github.com/OrlandoGolang/eps-conduit/blob/master/README.md)
- Create new test cases for your added features
- Run the testing suite and ensure that all tests pass
- Properly document new functions and variables according to Golang [best practices](http://blog.golang.org/godoc-documenting-go-code)
- Run your code through `golint` and `go fmt`

## 2. Fork, Develop, and Issue Pull Request

When you finished developing your fix or enhancement, issue a pull request against the `develop` branch of the base fork. In your pull request, be sure to reference the corresponding issue via #issue-number (ex. \#24).

Then, be on the lookout for feedback from the contributors, or a merge complete notification!
