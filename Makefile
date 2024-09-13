.PHONY:	bench
bench:
	@go test -bench=. -benchmem  ./...

.PHONY:	ut
ut:
	@go test -tags=goexperiment.arenas -race ./...

.PHONY:	setup
setup:
	@sh ./.script/setup.sh

.PHONY:	fmt
fmt:
	@sh ./.script/goimports.sh

.PHONY:	lint
lint:
	@golangci-lint run -c .golangci.yml

.PHONY: tidy
tidy:
	@go mod tidy -v

.PHONY: check
check:
	@$(MAKE) fmt
	@$(MAKE) tidy

.PHONY: install-user
install-user:
	@mkdir -p ~/.config/systemd/user
	@cp testservice.service ~/.config/systemd/user
	@systemctl --user daemon-reload

.PHONY: install-root
install-root:
	@if [ $$(id -u) -eq 0 ]; then \
		cp testservice.service /etc/systemd/system/; \
		systemctl daemon-reload; \
	else \
		echo "错误：需要 root 权限来安装服务"; \
		exit 1; \
	fi