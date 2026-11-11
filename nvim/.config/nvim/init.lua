vim.o.autowrite = true
vim.o.gdefault = true
vim.o.smartindent = true
vim.o.mouse = 'a'
vim.o.relativenumber = true
vim.o.wrap = false
vim.o.wildmode = 'list:longest'

vim.pack.add({
  'https://github.com/tpope/vim-sleuth',
  'https://github.com/lewis6991/gitsigns.nvim',
  { src = 'https://github.com/catppuccin/nvim', name = 'catppuccin' },
})

require('gitsigns').setup()

-- Latte or Mocha, following the terminal background
vim.cmd.colorscheme('catppuccin')
