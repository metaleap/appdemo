module appdemo

go 1.22.3

// there's a non-`require` main dependency: the seperate repo `yo`, which becomes available to `gopls` and
// `go <command>` by having it as a sibling-to-this-repo dir; plus, the following `go.work` file parent to both:

// go 1.22.3
// use ./yo
// use ./appdemo
