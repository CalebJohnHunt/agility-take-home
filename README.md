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

## Why

I used Go, [Go's templating library](https://pkg.go.dev/text/template), and [HTMX](https://htmx.org/).

I've used Go quite a bit on my hobby projects, so I felt pretty confident in using it for the API and logic portion of this challenge.
I thought I would be doing a TUI, which I've also done using Go before, but then I recalled Go's templating library.
A few online creators I follow have been using it recently for their hobby projects, so I figured it would be worth a go, and it was great!
If I needed to get this project done as quickly as possible, I probably would have skipped the templating for the CLI version, but since I knew I could be a bit creative I went for it.
And then when I decided I wanted to give this project more of an "app-like" experience, I realized I could re-use my templating for the website!
As for HTMX, I've read about it a few times in the past, and my coworker used it for a hackathon, so while I wasn't familiar with it exactly, I knew it would work for what I wanted here.

All-in-all, I'm pretty happy with my tech stack (if you can even call it that haha).
I've been looking for a reason to use templating and HTMX, so I'm glad I finally got it with this project.
