test:
	go test ./...
	wasmer run code_gen/.snapshots/main.wasm 1 2

install-vim:
	rm -rf ~/.vim/pack/plugins/start/vim-yum/
	cp -R tools/vim-yum/ ~/.vim/pack/plugins/start/vim-yum/

