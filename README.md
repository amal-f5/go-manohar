# test-eve


## Eve configuration

```json
{
    "app": {
        "name": "myapp",
        "primary": {
            "type": "go",
            "config": {
                "executable": "main.go"
            }
        }
    },
    "deployment": "mydeploy",
    "sources": {
        "name": "myrepo",
        "primary": {
            "name": "myrepo",
            "branch": "main",
            "uri": "https://github.com/amal-f5/go-manohar.git",
            "type": "repository"
        }
    }
}
```

## Changes required

Inside main.go, please import `unit.nginx.org/go` and replace `http.ListenAndServe` with `unit.ListenAndServe`
