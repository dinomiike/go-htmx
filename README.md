# Go and HTMX

This is a first pass at using HTMX based on a project from
[Frontend Masters](https://github.com/ThePrimeagen/fem-htmx-proj). It's using Go on the
backend, which is also new to me. They say it's easy to get started with. Let's see if
I've been lied to.

## Goal

Create a use case for HTMX and explore what it's capable of. At the end of this experiment,
I'd like to have an opinion about HTMX. (Spoilers, it impressed me in minutes)

## Findings

Half way through this concept of a simple contact list, I realize that half of the concern in
this code is in the backend. Unfortunately, I'm learning Go and Echo at the same time as HTMX,
so it's compounded unknowns.

HTMX is mature and thought through every use case. If you want to swap out content in other
places in the dom, you can do that with [hx-swap-oob](https://htmx.org/attributes/hx-swap-oob/),
which even has properties like `afterbegin`, which will insert the response html before the
first child of the target element.

Something I enjoy about ui development is being able to inspect and debug. It's hard to get lost
when I can follow line by line and see the browser do each step in order. HTMX makes it easy to
see what's been done by inspecting the dom as you go. You can also hit `htmx.logAll()` in the
browser console to see a detailed output log of what HTMX has done. This is extremely useful
for me at this point in the process. By using this and inspecting the dom, you can see anything
HTMX touches as part of its request, gets some kind of class added to it. That makes it easy
to identify later in case you want to get fancy with things like hx-indicators.

This seems to highlight the importance of fullstack engineering, where we simplify some of the
magic of the frontend and get everyone talking on the backend. Let's use the backend for things
like state and validation. Whether or not the backend is better or faster at those things, it
simplifies your program because you don't have to duplicate code. Many professional projects I've
worked on had state and validation on both client and server side. What if we didn't need to do
things that way any more?

Lot of fun working with this concept. I want to experiment with this in Express so I can move
faster and spend more time on HTMX itself.
