# FyneCamera

Build steps:

1. Clone down the fork of tools [here](https://github.com/bounce-chat/tools) and build the `fyne` binary using the `hkparker-camerax` branch
2. Clone down the fork of fyne [here](https://github.com/bounce-chat/fyne), check out the `hkparker-camerax` branch, and update this repo's go.mod `replace` directive to point to it
3. Use the binary from 1 to run `fyne package --app-id camera.demo -icon Icon.png -name FyneCamera -os android -tags migrated_fynedo`
4. `adb install FyneCamera.apk`
