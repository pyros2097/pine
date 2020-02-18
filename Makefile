test:
	go test ./...

install-vim:
	rm -rf ~/.vim/pack/plugins/start/vim-yum/
	cp -R tools/vim-yum/ ~/.vim/pack/plugins/start/vim-yum/