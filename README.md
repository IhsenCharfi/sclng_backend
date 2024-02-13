

# Backend Technical Test at Scalingo - sclng_backend

## Overview

This projects aims to return a list of 100 github repositories with some practical information about them. 

## Instructions

* To run you need to configure MYAPP_PORT on Dockerfile  (Default 3000).
```bash
  ENV MYAPP_PORT=3000

```
If you chnage it , you need to update Makefile docker command too
```bash
  docker run --rm -it -p 3000:3000 sclng_backend:latest
```
* You need to have generated a github token to make the calls to endpoints.


## Installation

Use git clone to install repo

```bash
  git clone https://github.com/IhsenCharfi/sclng_backend.git
```

## Execution

sclng_backend is a Go lang REST API for fetching github repositories data and make some stats.
To execute it you can run docker commands using Makefile
```bash
  make docker
```
This command will run a local docker container and will make the HTTP server available on configured PORT.

Start making requests calls on endpoints : 
* /repos
* /stats

Both endpoints may accept filter language



#### Request Example: /repos

List 100 from latest github public repositories. It can be filtered by "language" to show only repositories using this language.

#### Without filter
```bash
  localhost:3000/repos
```
#### With filter
```bash
  localhost:3000/repos?language=PHP
```
###### Response Example: /repos
```json
// GET /repos?language=PHP
{   
    "id": 717216661,
    "name": "TrabajoWebApi",
    "full_name": "facchin21/TrabajoWebApi",
     "owner": {
        "login": "facchin21"
    },
    "languages_url": "https://api.github.com/repos/facchin21/TrabajoWebApi/languages",
    "languages": {
        "CSS": 382,
        "HTML": 14561,
        "PHP": 45649
    }
},
{
    ..
}
...
```

#### Request Example: /stats

Show statistics about a list of 100 from latest github public repositories. 
It shows number of repositories and the repo with largest bytes.

It can be filtered by "language" to show only stats about repositories using this language.
#### Without filter
```bash
  localhost:3000/stats
```
#### Response Example: /stats
```json
// GET /stats
{
    "repo_with_largest_bytes": {
        "id": 717912938,
        "name": "PistonGame",
        "full_name": "ert2855/PistonGame",
        "owner": {
            "login": "ert2855"
        },
        "languages_url": "https://api.github.com/repos/ert2855/PistonGame/languages",
        "languages": {
            "ASP.NET": 181725,
            "C#": 14113082,
            "C++": 17765,
            "CMake": 426,
            "GAP": 9254,
            "HLSL": 7157,
            "Objective-C++": 8863,
            "ShaderLab": 7667
        }
    },
    "total_repos": 100
}
```
#### With filter
```bash
  localhost:3000/stats?language=PHP
```
#### Response Example: /stats
```json
// GET /stats?language=PHP
{   
    {
    "repo_with_largest_bytes": {
        "id": 718101925,
        "name": "FInal-proj",
        "full_name": "Franz1001/FInal-proj",
        "owner": {
            "login": "Franz1001"
        },
        "languages_url": "https://api.github.com/repos/Franz1001/FInal-proj/languages",
        "languages": {
            "CSS": 13266,
            "Hack": 11159,
            "JavaScript": 2493,
            "PHP": 319428
        }
    },
    "total_repos": 5
}
}

