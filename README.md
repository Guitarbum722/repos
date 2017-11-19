# repo-count
Sends a request to the github `Users` API and returns the public repository count for the specified user.

_Quick: get a repo count for a user_
```go
	count, err := repos.RepoCount("username")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("count: ", count)
```

_With a `GithubClient`_
```go
	c := repos.NewGithubClient()
	count, err := c.RepoCount("username")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("count: ", count)
```
