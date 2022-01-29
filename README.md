# goldmark-pikchr

Goldmark extension for rendering [pikchr](https://pikchr.org)
diagrams. Heavily copied from
[github.com/abhinav/goldmark-mermaid](https://github.com/abhinav/goldmark-mermaid). Uses
[gopikchr](https://github.com/gopikchr/gopikchr) for translating
pikchr diagrams to svg.

## Syntax

A simple pikchr diagram:

    ```pikchr
    box color red wid 2.6in \
        "Click on any diagram on this page" big \
        "to see the Pikchr source text" big
    ```

### Attributes

The pikchr plugin supports the following attributes:

- `toggle` (default `false`): if `true`, add javascript to the page to
  toggle between the diagram and its source when it is clicked.
- `limitwidth` (default `true`): if `true`, add a `max-width` style
  limiting the maximum width of the `<div>` containing the `<svg>`
  element to the width of the diagram. This prevents it from getting
  scaled up.

## Configuration

Add a `&pikchr.Extender{}` to the list of goldmark extensions:

```go
gm := goldmark.New(goldmark.WithExtensions(&pikchr.Extender{}))
```

If you wish to use the `toggle` attribute, add the following snippet to your `css`:

```css
.hidden {
   position: absolute !important;
   opacity: 0 !important;
   pointer-events: none !important;
   display: none !important;
}
```

## TODOs

- [x] Enable click-to-toggle-source as on pikchr.org examples
- [x] Add a `toggle` parameter to the pikchr blocks, so toggling
      source can be turned on or off per block
- [x] Add a `limitwidth` parameter so that `max-width` styling can be
      turned off.
