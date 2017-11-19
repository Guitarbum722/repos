# repo-count
Sends a request to the github `Users` API and returns the public repository count for the specified user.

```go
	c := repos.NewGithubClient()
	count, err := c.RepoCount("username")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("count: ", count)
```
