# Octogopher

## About

This is my experimental Go wrapper for the GitHub API.

For my current approach, I'm trying to make the library's interface as nice as possible (according to my own definition) even if this make the implementation more complex. Basically, I'm thinking that a lot more time is spent (hopefully) using a library than writing it, so let's optimize for that.

For example, I'm hard-coding each resource, such as a `Repo`, as its own `struct` type rather than representing them as an `interface`. This makes the library less resilient to changes in the GitHub API, but it makes code that uses Octogopher cleaner and easier to debug.

In the future, I want to generate the data models directly from the GitHub API documentation.

## Docs

Check out the automatically generated docs here:

http://godoc.org/github.com/jamesadney/octogopher

Or, you can always use `godoc` to view them in your terminal.

`godoc github.com/jamesadney/octogopher`
