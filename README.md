# Agility Take Home

## Running

This app was first made as a CLI and then turned into a website.

```sh
git clone git@github.com:CalebJohnHunt/agility-take-home.git

cd agility-take-home

# if you want to run the CLI:
git checkout cli # checkout cli branch
go run . Luke

# if you want to run the website
go run . --port <port> # default port: 8080
# it runs on localhost
```

## Iterations

+ [x] Barebones
    + Single name is passed in through CLI arg
    + Output non-pretty to CLI
        + Maybe even just JSON?
+ [x] Website
    + In and out through website
+ [x] Live updating
+ [ ] Make it look… better?
    + Both the code
    + and the website
