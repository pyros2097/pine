test:
	go test ./...

vim:
	rm -rf ~/.vim/pack/plugins/start/vim-yum/
	cp -R tools/vim-yum/ ~/.vim/pack/plugins/start/vim-yum/


# wat2wasm -d -v ww.wat
# wasm2wat code_gen/.snapshots/main.wasm
# wasm-objdump -x code_gen/.snapshots/main.wasm
# vim -b ww.wasm