test:
	go test ./...

vim:
	rm -rf ~/.vim/pack/plugins/start/vim-yum/
	cp -R tools/vim-yum/ ~/.vim/pack/plugins/start/vim-yum/


# wat2wasm -d -v ww.wat
# vim -b ww.wasm