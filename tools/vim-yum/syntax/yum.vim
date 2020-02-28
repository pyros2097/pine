" Quit when a (custom) syntax file was already loaded
if exists("b:current_syntax")
  finish
endif

syn case match

" Keywords
syn keyword     yumKeywords         module import extern return break continue if else match
hi def link     yumKeywords         Keyword

" hi def link     goStatement         Statement
" hi def link     goConditional       Conditional
" hi def link     goLabel             Label
" hi def link     goRepeat            Repeat

" Builtin types
syn keyword     yumType              byte i32 num map bool string error
hi def link     yumType              Type

" Builtin funcs
syn keyword     yumBuiltins                 panic print println
hi def link     yumBuiltins                 Identifier

" Builtin values
syn keyword     yumBoolean                  true false
syn keyword     yumPredefinedIdentifiers    nil
hi def link     yumBoolean                  Boolean
hi def link     yumPredefinedIdentifiers    goBoolean

" Comments; their contents
syn keyword     yumTodo              contained TODO FIXME XXX BUG
syn region      yumComment           start="#" end="$"
hi def link     yumComment           Comment
hi def link     yumTodo              Todo

" String, Numbers
syntax match   yumSpecial           "\\\d\d\d\|\\."
syntax region  yumString            start=+"+  skip=+\\\\\|\\"+  end=+"\|$+	contains=yumSpecial,@htmlPreproc
syntax region  yumTemplate          start=/"/  skip=/\\\\\|\\`\|\n/  end=/"\|$/ contains=yumTemplateSubstitution nextgroup=@yumComments,@yumSymbols skipwhite skipempty
syntax region  yumTemplateSubstitution matchgroup=yumTemplateSB contained start=/\${/ end=/}/ contains=@yumExpression
hi def link    yumString                 String
hi def link    yumTemplate               String
hi def link    yumTemplateSubstitution   Label
hi def link    yumTemplateSB             yumTemplateSubstitution
syntax match   yumSpecialCharacter "'\\.'"
syntax match   yumNumber           "-\=\<\d\+L\=\>\|0[xX][0-9a-fA-F]\+\>"
syntax match   yumFloat            /\<-\=\%(\d\+\.\d\+\|\d\+\.\|\.\d\+\)\%([eE][+-]\=\d\+\)\=\>/
syntax match   yumDollar           "\$"
hi def link    yumNumber                 Number
hi def link    yumFloat                  Number

" Regions
syn region      goParen             start='(' end=')' transparent
if go#config#FoldEnable('block')
  syn region    goBlock             start="{" end="}" transparent fold
else
  syn region    goBlock             start="{" end="}" transparent
endif

" Single-line var, const, and import.
syn match       goSingleDecl        /\%(import\|var\|const\) [^(]\@=/ contains=goImport,goVar,goConst

" Spaces after "[]"
syn match     yumSpaceError display "\%(\[\]\)\@<=\s\+"
hi def link   yumSpaceError        Error

" Space-tab error
if go#config#HighlightSpaceTabError()
  syn match goSpaceError display " \+\t"me=e-1
endif

" Trailing white space error
if go#config#HighlightTrailingWhitespaceError()
  syn match goSpaceError display excludenl "\s\+$"
endif

" Operators
syn match     yumOperator /:=\|||\|<-\|<=\|>=\|==\|!=\|-\|+\|%\|>\|<\|=\|&&/
hi def link   yumOperator          Operator

" Functions;
syn match     yumFunction          /\w\+ = / nextgroup=yumSimpleParams skipwhite skipnl
hi def link   yumFunction          Function

" Function calls;
syn match     yumFunctionCall      /\w\+\ze(/ contains=yumBuiltins,yumDeclaration
hi def link   yumFunctionCall      Type

" Fields;
syn match goField   /\.\w\+\
        \%(\%([\/\-\+*%]\)\|\
        \%([\[\]{}<\>\)]\)\|\
        \%([\!=\^|&]\)\|\
        \%([\n\r\ ]\)\|\
        \%([,\:.]\)\)\@=/hs=s+1
hi def link    goField              Identifier

" Variable Assignments
syn match goVarAssign /\v[_.[:alnum:]]+(,\s*[_.[:alnum:]]+)*\ze(\s*([-^+|^\/%&]|\*|\<\<|\>\>|\&\^)?\=[^=])/
hi def link   goVarAssign         Special

" Variable Declarations
syn match goVarDefs /\v\w+(,\s*\w+)*\ze(\s*:\=)/
hi def link   goVarDefs           Special

" JSX Highlighting
" detect jsx region
syntax region jsxRegion
      \ start=+\%(\%(\_[([,?:=+\-*/>{}]\|<\s\+\|&&\|||\|=>\|\<return\|\<default\|\<await\|\<yield\)\_s*\)\@<=<\_s*\%(>\|\z(\%(script\|T\s*>\s*(\)\@!\<[_$A-Za-z][-:._$A-Za-z0-9]*\>\)\%(\_s*\%([-+*)\]}&|?,]\|/\%([/*]\|\_s*>\)\@!\)\)\@!\)+
      \ end=++
      \ contains=jsxElement

" <tag id="sample">
" ~~~~~~~~~~~~~~~~~
" and self close tag
" <tag id="sample"   />
" ~~~~~~~~~~~~~~~~~~~
syntax region jsxTag
      \ start=+<+
      \ matchgroup=jsxOpenPunct
      \ end=+>+
      \ matchgroup=NONE
      \ end=+\%(/\_s*>\)\@=+
      \ contained
      \ contains=jsxOpenTag,jsxAttrib,jsxExpressionBlock,jsxSpreadOperator,jsComment,@javascriptComments,javaScriptLineComment,javaScriptComment,typescriptLineComment,typescriptComment
      \ keepend
      \ extend
      \ skipwhite
      \ skipempty
      \ nextgroup=jsxCloseString

" <tag></tag>
" ~~~~~~~~~~~
" and fragment
" <></>
" ~~~~~
" and self close tag
" <tag />
" ~~~~~~~
syntax region jsxElement
      \ start=+<\_s*\%(>\|\${\|\z(\<[-:._$A-Za-z0-9]\+\>\)\)+
      \ end=+/\_s*>+
      \ end=+<\_s*/\_s*\z1\_s*>+
      \ contains=jsxElement,jsxTag,jsxExpressionBlock,jsxComment,jsxCloseTag,@Spell
      \ keepend
      \ extend
      \ contained
      \ fold

" <tag key={this.props.key}>
" ~~~~
" and fragment start tag
" <>
" ~~
exe 'syntax region jsxOpenTag
      \ matchgroup=jsxOpenPunct
      \ start=+<+
      \ end=+>+
      \ matchgroup=NONE
      \ end=+\>+
      \ contained
      \ contains=jsxTagName
      \ nextgroup=jsxAttrib
      \ skipwhite
      \ skipempty
      \ ' .''


" <tag key={this.props.key}>
"          ~~~~~~~~~~~~~~~~
syntax region jsxExpressionBlock
      \ matchgroup=jsxBraces
      \ start=+{+
      \ end=+}+
      \ contained
      \ extend
      \ contains=@jsExpression,jsSpreadExpression,@javascriptExpression,javascriptSpreadOp,@javaScriptEmbededExpr,@typescriptExpression,typescriptObjectSpread,jsComment,@javascriptComments,javaScriptLineComment,javaScriptComment,typescriptLineComment,typescriptComment

" <foo.bar>
"     ~
syntax match jsxDot +\.+ contained

" <foo:bar>
"     ~
syntax match jsxNamespace +:+ contained

" <tag id="sample">
"        ~
syntax match jsxEqual +=+ contained skipwhite skipempty nextgroup=jsxString,jsxExpressionBlock,jsxRegion

" <tag />
"      ~~
syntax match jsxCloseString +/\_s*>+ contained

" </tag>
" ~~~~~~
" and fragment close tag
" </>
" ~~~
syntax region jsxCloseTag
      \ matchgroup=jsxClosePunct
      \ start=+<\_s*/+
      \ end=+>+
      \ contained
      \ contains=jsxTagName

" <tag key={this.props.key}>
"      ~~~
syntax match jsxAttrib
      \ +\<[_$A-Za-z][-:_$A-Za-z0-9]*\>+
      \ contained
      \ nextgroup=jsxEqual
      \ skipwhite
      \ skipempty
      \ contains=jsxAttribKeyword,jsxNamespace

" <MyComponent ...>
"  ~~~~~~~~~~~
" NOT
" <someCamel ...>
"      ~~~~~
exe 'syntax match jsxComponentName
      \ +\<[_$]\?[A-Z][-_$A-Za-z0-9]*\>+
      \ contained
      \ ' .''

" <tag key={this.props.key}>
"  ~~~
exe 'syntax match jsxTagName
      \ +\<[-:._$A-Za-z0-9]\+\>+
      \ contained
      \ contains=jsxComponentName,jsxDot,jsxNamespace
      \ nextgroup=jsxAttrib
      \ skipempty
      \ skipwhite
      \ ' .''

" <tag id="sample">
"         ~~~~~~~~
" and
" <tag id='sample'>
"         ~~~~~~~~
syntax region jsxString start=+\z(["']\)+  skip=+\\\\\|\\\z1\|\\\n+  end=+\z1+ contained contains=@Spell

let s:tags = get(g:, 'vim_jsx_pretty_template_tags', ['html', 'jsx'])
let s:enable_tagged_jsx = !empty(s:tags)

" add support to JSX inside the tagged template string
" https://github.com/developit/htm
if s:enable_tagged_jsx
  exe 'syntax match jsxRegion +\%(' . join(s:tags, '\|') . '\)\%(\_s*`\)\@=+ contains=jsTemplateStringTag,jsTaggedTemplate,javascriptTagRef skipwhite skipempty nextgroup=jsxTaggedRegion'

  syntax region jsxTaggedRegion
        \ matchgroup=jsxBackticks
        \ start=+`+
        \ end=+`+
        \ extend
        \ contained
        \ contains=jsxElement,jsxExpressionBlock
        \ transparent

  syntax region jsxExpressionBlock
        \ matchgroup=jsxBraces
        \ start=+\${+
        \ end=+}+
        \ extend
        \ contained
        \ contains=@jsExpression,jsSpreadExpression,@javascriptExpression,javascriptSpreadOp,@javaScriptEmbededExpr,@typescriptExpression,typescriptObjectSpread

  syntax region jsxOpenTag
        \ matchgroup=jsxOpenPunct
        \ start=+<\%(\${\)\@=+
        \ matchgroup=NONE
        \ end=+}\@1<=+
        \ contained
        \ contains=jsxExpressionBlock
        \ skipwhite
        \ skipempty
        \ nextgroup=jsxAttrib,jsxSpreadOperator

  syntax keyword jsxAttribKeyword class contained

  syntax match jsxSpreadOperator +\.\.\.+ contained nextgroup=jsxExpressionBlock skipwhite

  syntax match jsxCloseTag +<//>+ contained

  syntax match jsxComment +<!--\_.\{-}-->+
endif

" Highlight the tag name
highlight def link jsxTag Function
highlight def link jsxTagName Identifier
highlight def link jsxComponentName Function

highlight def link jsxAttrib Type
highlight def link jsxAttribKeyword jsxAttrib
highlight def link jsxString String
highlight def link jsxComment Comment

highlight def link jsxDot Operator
highlight def link jsxNamespace Operator
highlight def link jsxEqual Operator
highlight def link jsxSpreadOperator Operator
highlight def link jsxBraces Special

" Highlight the jsxCloseString (i.e. />), jsxPunct (i.e. <,>) and jsxCloseTag (i.e. <//>)
highlight def link jsxCloseString Comment
highlight def link jsxOpenPunct jsxPunct

highlight def link jsxPunct jsxCloseString
highlight def link jsxClosePunct jsxPunct
highlight def link jsxCloseTag jsxCloseString

let s:vim_jsx_pretty_colorful_config = get(g:, 'vim_jsx_pretty_colorful_config', 0)

if s:vim_jsx_pretty_colorful_config == 1
  highlight def link jsObjectKey Label
  highlight def link jsArrowFuncArgs Type
  highlight def link jsFuncArgs Type
endif


augroup vim-yum-hi
  autocmd!
  autocmd ColorScheme * call s:hi()
augroup end

let b:current_syntax = "yum"
