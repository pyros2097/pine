test:
	go test ./...

vim:
	rm -rf ~/.vim/pack/plugins/start/vim-pine/
	cp -R tools/vim-pine/ ~/.vim/pack/plugins/start/vim-pine/


# wat2wasm -d -v ww.wat
# wasm2wat code_gen/.snapshots/main.wasm
# wasm-objdump -x code_gen/.snapshots/main.wasm
# vim -b ww.wasm

# https://github.com/medikoo/es6-template-strings
# https://github.com/souldreamer/string-template-parser