# This project is no longer maintained. Please use [StackStats](https://github.com/kubo550/edge-functions/blob/main/supabase/functions/stack-stats/README.md)

<div align="center">
  <a href="https://github.com/kubo550/stack-stats">
     <img src="https://user-images.githubusercontent.com/43968748/168917115-9587fc8f-2648-43da-b10f-39743f78295e.png" alt="stack stats logo" />
  </a>

  
  <h2 align="center">Stack Overflow Stats SVG Generator</h2>

  <p align="center">
     Dynamically generated stack overflow stats for your github readmes
    <br />
  </p>
</div>


<div align="center">

<a href="https://github.com/kubo550/stack-stats/stargazers"><img src="https://img.shields.io/github/stars/kubo550/stack-stats" alt="Stars Badge"/></a>
<a href="https://github.com/kubo550/stack-stats/network/members"><img src="https://img.shields.io/github/forks/kubo550/stack-stats" alt="Forks Badge"/></a>
<a href="https://github.com/kubo550/stack-stats/pulls"><img src="https://img.shields.io/github/issues-pr/kubo550/stack-stats" alt="Pull Requests Badge"/></a>
<a href="https://github.com/kubo550/stack-stats/issues"><img src="https://img.shields.io/github/issues/kubo550/stack-stats" alt="Issues Badge"/></a>
<a href="https://github.com/kubo550/stack-stats/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/kubo550/stack-stats?color=2b9348"></a>
<a href="https://github.com/kubo550/stack-stats/blob/master/LICENSE"><img src="https://img.shields.io/github/license/kubo550/stack-stats?color=2b9348" alt="License Badge"/></a>

</div>


## Overview

<!-- HERE YOU GO!  -->

<img src="https://stack-stats.herokuapp.com/stats?id=5798347" alt="stack stats" />&nbsp;
<img src="https://stack-stats.herokuapp.com/stats?id=14513625" alt="stack stats" />&nbsp;
<img src="https://stack-stats.herokuapp.com/stats?id=6904888" alt="stack stats" />&nbsp;
<img src="https://stack-stats.herokuapp.com/stats?id=3397217" alt="stack stats" />&nbsp;




Temporary not working

## How do I use it?

The only thing you need to do is to add the following code to your page:

```md
![stack stats](https://stack-stats.herokuapp.com/stats?id=<id>)
```

Important note: the `id` is the id of the stack you want to display.

In my case, I have a stack with id 14513625.

```md
![stack stats](https://stack-stats.herokuapp.com/stats?id=14513625)
```


Also, you can wrap the above in an anchor tag to make it easier to move on to your stack profile:

```
<a href="https://stack-stats.herokuapp.com/stats?id=14513625" target="_blank" rel="noopener noreferrer">
    <img src="https://stack-stats.herokuapp.com/stats?id=14513625" alt="stack stats" />&nbsp;
</a>
```

<a href="https://stackoverflow.com/users/14513625/jakub-kurdziel" target="_blank" rel="noopener noreferrer" title="My Stack Overflow Profile">
    <img src="https://stack-stats.herokuapp.com/stats?id=14513625" alt="stack stats" />
</a>





## Why my stats are not updating?

It is because GitHub uses a cache to store the data. The data is updated every 2 hours, but you can manually update the cache:

Only you need to do is to right-click on the svg and select copy link location then type this command:


```bash
$ curl -X PURGE https://camo.githubusercontent.com/4d04abe0044d94fefcf9af2133223....
> {"status": "ok", "id": "216-8675309-1008701"}

```
change the url to the one you want to purge.


## Run it locally

To run it locally, you just need to run the following commands:

```go
go run ./src/server.go 8080
```

## Tech Stack

* ![golang](https://img.shields.io/badge/GO-05122A?style=flat&logo=go)&nbsp; v1.18

* ![fiber](https://img.shields.io/badge/Fiber-05122A?style=flat&logo=go)&nbsp; v2.32.0

* ![gock](https://img.shields.io/badge/gock-05122A?style=flat&logo=go)&nbsp; v1.0

* ![stackexchange](https://img.shields.io/badge/stackexchange-05122A?style=flat&logo=stackexchange)&nbsp; API v2.3

* ![heroku](https://img.shields.io/badge/heroku-05122A?style=flat&logo=heroku)&nbsp;



