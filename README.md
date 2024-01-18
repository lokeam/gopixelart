# GoPixelArt
Pixel Art Editor made in Golang.

Define your image size and create your own custom sprites, or edit existing images.

Uses the [Fyne](https://fyne.io/) toolkit for Widget and Canvas manipulation.

## Features
- Create new files, save and load existing work
- Create and use custom Swatches (Palletes)
- Color picker
- Pan / Zoom image using the mouse scroll wheel
- Pixel brush indicator

## How To Use
In order to run this application, you'll need the following:
- Git
- Go

### Installation and Setup
```bash
# Clone the repo
$  git clone git@github.com:lokeam/gopixelart.git

# Navigate to the gopixelart directory
$  cd gopixelart

# Install all app dependencies
$  go mody tidy

# Boot up the app
$  go run -v ./pixel
```

### Notes:
`macOS users`
The `File` menu used for creating, loading and saving files will not appear within the Pixel Editor application window as fyne toolkit integrates natively with the macOS top menu bar.

`Windows users`
The `File` menu will render at the top of the app window.

## License
This project is licensed under the terms of the MIT license.