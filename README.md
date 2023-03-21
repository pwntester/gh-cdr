# gh-cdr

Manage all your GitHub.com repos locally

## Installation

You need to set up a function in your bash `rc` file to CD into the repo directory. eg:

```
function cdr() {cd $(gh cdr $@)}
```

## Usage

Once the above function is created you can use either `cdr pwntester/gh-cdr` or `cdr pwntester gh-cdr` to CD into your local copy of the repo or clone it if it was not downloaded previously
