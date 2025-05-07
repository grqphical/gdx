# gdx

A language server for Godot/GDScript written in Golang.

The goal of this project is to make an improved and faster language server for working with external editors in your Godot projects.

> **⚠️NOTE:** This project is in early alpha. It is very unstable and definitely not suitable for use (yet)

## Planned Features

- [x] Communication through STDIN/STDOUT (not through localhost)

- [ ] Autocomplete

- [ ] Godot documentation lookups

- [ ] Support for static type checking via type annotations

## Installation

Clone this repository and run `go build .`

If you want to use this with Neovim you can add this script to your Neovim config:
```lua
local cwd = vim.fn.getcwd()
local util = require("lspconfig.util")

function get_root_dir()
    return util.root_pattern("project.godot")(cwd)
end

vim.lsp.config['gdx'] = {
    cmd = { os.getenv("HOME") .. "/dev/go/gdx/gdx" },
    root_dir = get_root_dir(),
    filetypes = { "gdscript" }
}

vim.lsp.enable("gdx")
```
Note this config requires Neovim 0.11+

Currently VSCode is unsupported however I plan to create an extension in the future to work with GDX.

## License

gdx is licensed under the MIT License
