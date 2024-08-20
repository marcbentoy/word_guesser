let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/projects/word_guesser
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +16 cmd/guesser/main.go
badd +1 guesser.go
badd +48 monkey.go
argglobal
%argdel
$argadd NvimTree_1
edit guesser.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
wincmd _ | wincmd |
vsplit
2wincmd h
wincmd w
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 75) / 151)
exe 'vert 2resize ' . ((&columns * 60 + 75) / 151)
exe 'vert 3resize ' . ((&columns * 59 + 75) / 151)
argglobal
enew
file NvimTree_1
balt monkey.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
lcd ~/projects/word_guesser
wincmd w
argglobal
balt ~/projects/word_guesser/monkey.go
setlocal fdm=marker
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 138 - ((37 * winheight(0) + 23) / 47)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 138
normal! 010|
lcd ~/projects/word_guesser
wincmd w
argglobal
if bufexists(fnamemodify("~/projects/word_guesser/monkey.go", ":p")) | buffer ~/projects/word_guesser/monkey.go | else | edit ~/projects/word_guesser/monkey.go | endif
if &buftype ==# 'terminal'
  silent file ~/projects/word_guesser/monkey.go
endif
balt ~/projects/word_guesser/guesser.go
setlocal fdm=marker
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 21 - ((20 * winheight(0) + 23) / 47)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 21
normal! 018|
lcd ~/projects/word_guesser
wincmd w
2wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 75) / 151)
exe 'vert 2resize ' . ((&columns * 60 + 75) / 151)
exe 'vert 3resize ' . ((&columns * 59 + 75) / 151)
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
nohlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
