<h1>
  <img alt="ExtLib Logo" height="48" src="https://gist.githubusercontent.com/gtramontina/f3a29963a7aa558d72098f149ebe0e09/raw/ed1c72dae7f341b6ce426c09afa883ce1d6650eb/go-extlib.svg">
  <img alt="ExtLib Gopher Mascot" height="48" src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/9edef573bbbfd880698627fea674fad14bbce477.png" align="right">
</h1>

<sup><b>⚠️ Note:</b> This is still an <em>experiment</em>.</sup>
<a href="https://github.com/gtramontina/go-extlib/actions/workflows/ci.yml"><img alt="Build Badge" src="https://github.com/gtramontina/go-collections/actions/workflows/ci.yml/badge.svg" align="right"></a>

## Motivation

My main motivation was to try out generics in Go. I jumped straight into the usual suspects: `Filter` and `Map`. When testing these, I saw the need for assertion functions: another opportunity to exercise some more generics! It lives as an [internal](https://github.com/gtramontina/go-extlib/blob/main/internal/assert/assert.go) package at the moment.

Later I thought it would be nice to implement `Set`, and then `HashMap`, with which I got carried away and ended up implementing a somewhat naïve [hasher](https://github.com/gtramontina/go-extlib/blob/main/internal/hash/hash.go) (you can read more on the rationale on [this commit](https://github.com/gtramontina/go-extlib/commit/808ac8236c433587c4dc2f85479c1189a5df6010)).

When I realized, I was already having fun with `Maybe`, `Either` and `Result`. Here is where I stumbled upon some limitations of Go generics. Right when I was experimenting with it, I came across a blog post by [@hypirion](https://github.com/hypirion) titled "[Type-Safe HTTP Servers in Go via Generics](https://hypirion.com/musings/type-safe-http-servers-in-go-via-generics)" where the author bumps into these same limitations.

When writing software in Go, I've always wanted slightly higher level constructs, especially to represent business domain concepts. Some may say I haven't fully embraced Go. Perhaps I haven't 🤷. But I sure know people who'd also like to go up in abstraction a tad bit.

Hope this is enough motivation… 😅

⚠️ As the note above says, this is an experiment and no attention was paid to performance or allocations.

##  The ExtLib

<sup>🚧 This section is a work in progress.</sup>

---

<p align="right">
  <sub><sup><i>Gopher: </i><a href="https://gopherize.me/gopher/9edef573bbbfd880698627fea674fad14bbce477">gopherize.me</a></sup></sub>
</p>
