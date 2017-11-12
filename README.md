go-gopathless (part II)
=======================

**gopathless** demonstrates:

- a project-local $GOPATH
- with no fuss.

That's it.


What?  Really?
--------------

Yeah.  This is way easy.

Check out [the Part 1 repo](https://github.com/warpfork/go-gopathless)
for the majority of the story.


Part 2
------

The Go code in this repo depends on the [github.com/warpfork/go-gopathless](https://github.com/warpfork/go-gopathless) package,
which *also* uses a project-local GOPATH.

This just works.  It's completely fine with normal tools.

Run this:

```bash
./go get
./go test
# it worked, didn't it?  I bet it did.
```


Did anything get *weird*?
-------------------------

Nope.

Try `./go test` again.  Did it do anything silly like run the tests from
your dependencies for no reason?  Nope.  It did the right thing:
just your tests run.

Try `./go list -f '{{join .Deps "\n"}}'`.  Is it 100% correct and sane?
Yes, yes it is.

It's just GOPATH.
