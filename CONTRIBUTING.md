# <a name="contributing">Contributing Overview</a>
Please do! Thanks for your help improving the project! :balloon:

All contributors are welcome. Please see the [newcomers welcome guide](https://layer5.io/community/newcomers) for how, where and why to contribute. This project is community-built and welcomes collaboration. Contributors are expected to adhere to our [Code of Conduct](.CODE_OF_CONDUCT.md).

Not sure where to start? First, see the [newcomers welcome guide](https://layer5.io/community/newcomers). Grab an open issue with the [help-wanted label](../../labels/help%20wanted) and jump in. Join the [Slack account](http://slack.layer5.io) and engage in conversation. Create a [new issue](/../../issues/new/choose) if needed.  All [pull requests](/../../pulls) should reference an open [issue](/../../issues). Include keywords in your pull request descriptions, as well as commit messages, to [automatically close issues in GitHub](https://help.github.com/en/github/managing-your-work-on-github/closing-issues-using-keywords).

**Sections**
- <a name="contributing">General Contribution Flow</a>
  - <a href="#commit-signing">Developer Certificate of Origin</a>

Relevant coding style guidelines are the Go Code Review Comments and the Formatting and style section of Peter Bourgon's Go: Best Practices for Production Environments.

# <a name="contributing">General Contribution Flow</a>

In order to contribute to Meshery, please follow the fork-and-pull request workflow described [here](./CONTRIBUTING-gitflow.md).

## <a name="commit-signing">Signing-off on Commits (Developer Certificate of Origin)</a>

To contribute to this project, you must agree to the Developer Certificate of
Origin (DCO) for each commit you make. The DCO is a simple statement that you,
as a contributor, have the legal right to make the contribution.

See the [DCO](https://developercertificate.org) file for the full text of what you must agree to
and how it works [here](https://github.com/probot/dco#how-it-works).
To signify that you agree to the DCO for contributions, you simply add a line to each of your
git commit messages:

```
Signed-off-by: Jane Smith <jane.smith@example.com>
```

In most cases, you can add this signoff to your commit automatically with the
`-s` or `--signoff` flag to `git commit`. You must use your real name and a reachable email
address (sorry, no pseudonyms or anonymous contributions). An example of signing off on a commit:
```
$ commit -s -m “my commit message w/signoff”
```

To ensure all your commits are signed, you may choose to add this alias to your global ```.gitconfig```:

*~/.gitconfig*
```
[alias]
  amend = commit -s --amend
  cm = commit -s -m
  commit = commit -s
```
Or you may configure your IDE, for example, Visual Studio Code to automatically sign-off commits for you:

<a href="https://user-images.githubusercontent.com/7570704/64490167-98906400-d25a-11e9-8b8a-5f465b854d49.png" ><img src="https://user-images.githubusercontent.com/7570704/64490167-98906400-d25a-11e9-8b8a-5f465b854d49.png" width="50%"><a>

# <a name="getnighthawk">About GetNightHawk?</a>
GetNighthawk aims to provide generally-available distributions of Nighthawk under different architectures and platforms and easy-to-use tooling for installation and operation. This will include creating distributions of Nighthawk as well as augmenting existing tooling, Meshery, to retrieve these arch-specific packages and update their deployments.


# <a name="contribute">How to Contribute?</a>
You can contribute to the GetNightHawk main project and documentation or you could help in building and contributing to the site which is running on Jekyll, a static site generator(SSG).

## Steps to contribute to GetNightHawk project
1. Visit the [Project documentation](https://docs.google.com/document/d/1lHfMo4iIx2WXFZIspfHyxTsPR1T63_2IV5NUkgxoo0w/edit#) 
2. Suggest requests, changes, or features by creating an issue on the repository
3. Look at existing issues to tackle and get involved with


## Steps to build GetNightHawk site

1. Fork this repository
2. Clone your forked copy to your local computer
3. `cd` into the getnighthawk repo
4. `cd docs` to access the GetNighthawk Jekyll website codebase
5. Run the following code in your terminal `bundle install`.
6. Run `make site` in order to serve the jekyll website on your machine at `http://127.0.0.1:4000/`
( `bundle exec jekyll serve` can also be used for spinning up the Jekyll website at `http://127.0.0.1:4000/`)


### UI Lint Rules

Layer5 uses ES-Lint to maintain code quality & consistency in our UI Code.

# <a name="maintaining"> Reviews</a>
All contributors are invited to review pull requests. See this short video on [how to review a pull request](https://www.youtube.com/watch?v=isLfo7jfE6g&feature=youtu.be).

# New to Git?
Resources: https://lab.github.com and https://try.github.com/

# Notes related to Nighthawk [Transform function](https://github.com/layer5io/getnighthawk/blob/d13923e181c6bd8daddacee74d8a8d571658d839/pkg/client/transform.go)
1. The format of results generated by Nighthawk is different than format of Fortio's result, to convert Nighthawk's results into Fortio compatible format we use [transform function](https://github.com/layer5io/getnighthawk/blob/d13923e181c6bd8daddacee74d8a8d571658d839/pkg/client/transform.go#L17), based on the [proto file](https://github.com/envoyproxy/nighthawk/blob/f42add9bb5c71a5652db2f9471dda357b96d5ee2/api/client/transform/fortio.proto#L31).

2. The proto file has discrepancies and is not completely compatible with fortio. 
Check the [issue](https://github.com/envoyproxy/nighthawk/issues/932) and go through comments for clarity.

3. To handle this mismatched types, we have re-implemented [Unmarshall function](https://github.com/meshery/meshery/blob/b3199637ec7a5a25cac476cf1d0d2c474c978c37/server/helpers/load_test_interface.go#L447) to facilitate compatibility.

In case of incompatibility in future, we need to alter the [transorm function](https://github.com/layer5io/getnighthawk/blob/d13923e181c6bd8daddacee74d8a8d571658d839/pkg/client/transform.go) or [Unmarshall function](https://github.com/meshery/meshery/blob/b3199637ec7a5a25cac476cf1d0d2c474c978c37/server/helpers/load_test_interface.go#L447).

*The above Unmarshall function needs to be removed once the [issue](https://github.com/envoyproxy/nighthawk/issues/932) is fixed.*

### License

This repository and site are available as open source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).

### About Layer5

**Community First**
<p>The <a href="https://layer5.io">Layer5</a> community represents the largest collection of service mesh projects and their maintainers in the world.</p>

**Open Source First**
<p>Our projects establish industry standards and enable service developers, owners, and operators with repeatable patterns and best practices for managing all aspects of distributed services. Our shared commitment to the open source spirit push the Layer5 community and its projects forward.</p>
