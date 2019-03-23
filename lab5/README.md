# Lab work N 5

## Content info

Based on Lab4. Added Encoder interface (`encoding/iface.go`) that should be implemented in your plugins. By default `Default` encoder is added (does nothing).
In plugins folder you can find AES & base64 encoder plugins.


## Installation

- Install glide (https://github.com/Masterminds/glide)
- Run `glide install` to fetch vendors.
- Run `make plugin NAME=FOLDER_UNDER_plugins/src/_NAME/FILENAME` to build plugin code
- Run `make build` command - it'll compile to `bin` folder.
- Run ./bin/app --plugins FOLDER_UNDER_plugins/src/_NAME/FILENAME --plugins FOLDER_UNDER_plugins/src/_NAME/FILENAME_2

