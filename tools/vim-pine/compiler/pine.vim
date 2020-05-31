" Vim compiler file
" Language:     Pine
" Maintainer:   pyros2097
" URL:          https://github.com/pyros2097/vim-pine

" if exists("current_compiler")
"   finish
" endif

" let current_compiler = "pine"

" if exists(":CompilerSet") != 2 " older Vim always used :setlocal
"   command -nargs=* CompilerSet setlocal <args>
" endif

" let s:cpo_save = &cpo

" CompilerSet makeprg=pine\ c\ --verbosity:0\ --listfullpaths\ $*\ %:p

" CompilerSet errorformat=
"   \%-GHint:\ %m,
"   \%E%f(%l\\,\ %c)\ Error:\ %m,
"   \%W%f(%l\\,\ %c)\ Hint:\ %m

" let &cpo = s:cpo_save
" unlet s:cpo_save

" let g:syntastic_pine_checkers = ['pine']

" if exists("current_compiler")
"   finish
" endif
" let current_compiler = "eslint"

" if exists(":CompilerSet") != 2
"   command! -nargs=* CompilerSet setlocal <args>
" endif

" CompilerSet makeprg=eslint\ -f\ compact\ %
" CompilerSet errorformat=%f:\ line\ %l\\,\ col\ %c\\,\ %m
