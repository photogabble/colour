<h1 align="center">Colour</h1>
<p align="center"><em>A lightweight toolset for colour manipulation, ported from <a href="https://github.com/styled-components/polished/tree/master/src/color">polished.js</a></em></p>

<p align="center">
  <a href="LICENSE"><img src="https://img.shields.io/github/license/photogabble/colour.svg" alt="License"></a>
  <a href="https://gitmoji.carloscuesta.me/"><img src="https://img.shields.io/badge/gitmoji-%20😜%20😍-FFDD67.svg" alt="Gitmoji"></a>
</p>

## About

While writing [Go Pixels Fight!](https://github.com/photogabble/go-pixel-fight) I needed to be able to mix two colours to produce an average of the two to show which colour was winning; during my research I stumbled upon the _color_ helpers in [polished.js](https://github.com/styled-components/polished) and while they didn't work for what I needed I thought it would be useful to myself in the future and others if they were ported to Go.

Primarily as a learning exercise in how to write tests with Golang I release this port to the wild in the hope that it may be of some use to someone.

## Building

...

## Todo

* [ ] adjust hue
* [ ] complement
* [ ] darken
* [ ] desaturate
* [x] get luminance
* [ ] grayscale
* [ ] HSL
* [ ] HSL to Colour String
* [ ] HSLA
* [x] invert
* [ ] lighten
* [x] mix
* [ ] opacify
* [ ] parse to HSL
* [ ] parse to RGB
* [ ] readable colour
* [ ] RGB
* [ ] RGB to Colour String
* [ ] RGBA
* [ ] saturate
* [ ] set hue
* [ ] set lightness
* [ ] set saturation
* [x] shade
* [x] tint
* [ ] to Colour String
* [ ] transparentise

## License

[MIT](LICENSE)