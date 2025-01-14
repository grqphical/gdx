# gdx
![GitHub Release](https://img.shields.io/github/v/release/grqphical/gdx)

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
local client = vim.lsp.start_client({
        name = "gdx",
        cmd = { "PATH_TO_BUILT_GDX_BINARY" },
})

vim.api.nvim_create_autocmd("FileType", {
        pattern = { "*.gd" },
        callback = function()
                vim.lsp.buf_attach_client(0, client)
        end,
})
```

Currently VSCode is unsupported however I plan to create an extension in the future to work with GDX.

## License

gdx is licensed under the MIT License
