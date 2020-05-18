if exists("current_compiler")
  finish
endif

let current_compiler = "pine"

if exists(":CompilerSet") != 2 " older Vim always used :setlocal
  command -nargs=* CompilerSet setlocal <args>
endif

let s:cpo_save = &cpo
set cpo-=C

CompilerSet makeprg=pine\ c\ --verbosity:0\ --listfullpaths\ $*\ %:p

CompilerSet errorformat=
  \%-GHint:\ %m,
  \%E%f(%l\\,\ %c)\ Error:\ %m,
  \%W%f(%l\\,\ %c)\ Hint:\ %m

let &cpo = s:cpo_save
unlet s:cpo_save

let g:syntastic_pine_checkers = ['pine']