arsenal_version=1.0.0

# chaosArsenal-OS
arsenal_os_repo=git@mq.code.sangfor.org:CloudTechnology/dfx/rasfire/arsenal-os.git
arsenal_os_branch=master

# stress-ng
stress_ng_repo=http://mirrors.sangfor.org/github/ColinIanKing/stress-ng.git
stress_ng_branch=master

# chasArsenal-hardware
arsenal_hardware_repo=git@mq.code.sangfor.org:CloudTechnology/dfx/rasfire/arsenal-hardware.git
arsenal_hardware_branch=master

build_dir=build
build_pkg_dir_name=chaosArsenal-$(arsenal_version)
build_pkg_dir=$(build_dir)/$(build_pkg_dir_name)
build_arsenal_package=$(build_pkg_dir).tar.gz

build_pkg_json=configs/arsenal-spec-$(arsenal_version).json

# project build dependents directory
build_pkg_dependents=$(build_dir)/dependents

# build dependents source code
build_arsenal_os_src=$(build_pkg_dependents)/arsenal-os
build_stress_ng_src=$(build_pkg_dependents)/stress-ng
build_arsenal_hardware_src=$(build_pkg_dependents)/arsenal-hardware

# binary store path
build_pkg_binary_dir=$(build_pkg_dir)/bin
build_pkg_third_party_tools_dir=$(build_pkg_binary_dir)/third_party_tools
build_pkg_logs_dir=$(build_pkg_dir)/logs

build_package: build
	tar zcvf $(build_arsenal_package) -C $(build_dir) $(build_pkg_dir_name)

build: build_prepare
	go build -ldflags "-linkmode external -extldflags -static" \
	-o $(build_pkg_dir) cli/arsenal.go
	cp $(build_pkg_json) $(build_pkg_dir)

os:
ifneq ($(build_arsenal_os_src), $(wildcard $(build_arsenal_os_src)))
	git clone -b $(arsenal_os_branch) $(arsenal_os_repo) $(build_arsenal_os_src)
endif
	make -C $(build_arsenal_os_src)
	cp $(build_arsenal_os_src)/arsenal-os $(build_pkg_binary_dir)

stress_ng:
ifneq ($(build_stress_ng_src), $(wildcard $(build_stress_ng_src)))
	git clone -b $(stress_ng_branch) $(stress_ng_repo) $(build_stress_ng_src)
endif
	make -C $(build_stress_ng_src) clean
	make -C $(build_stress_ng_src) STATIC=1 -j$(shell nproc)
	cp $(build_stress_ng_src)/stress-ng $(build_pkg_third_party_tools_dir)

hardware:
ifneq ($(build_arsenal_hardware_src), $(wildcard $(build_arsenal_hardware_src)))
	git clone -b $(arsenal_hardware_branch) $(arsenal_hardware_repo) $(build_arsenal_hardware_src)
endif
	make -C $(build_arsenal_hardware_src)
	cp $(build_arsenal_hardware_src)/arsenal-hardware $(build_pkg_binary_dir)

build_prepare: make_build_dir hardware os stress_ng

make_build_dir:
ifneq ($(build_pkg_dependents), $(wildcard $(build_pkg_dependents)))
	mkdir -p $(build_pkg_dependents)
endif
ifneq ($(build_pkg_third_party_tools_dir), $(wildcard $(build_pkg_third_party_tools_dir)))
	mkdir -p $(build_pkg_third_party_tools_dir)
endif
ifneq ($(build_pkg_logs_dir), $(wildcard $(build_pkg_logs_dir)))
	mkdir -p $(build_pkg_logs_dir)
endif

clean:
ifeq ($(build_pkg_dir), $(wildcard $(build_pkg_dir)))
	rm -rf $(build_pkg_dir)
endif
ifeq ($(build_arsenal_package), $(wildcard $(build_arsenal_package)))
	rm $(build_arsenal_package)
endif
	go clean ./...
	make -C $(build_arsenal_os_src) clean
	make -C $(build_stress_ng_src) clean
	make -C $(build_arsenal_hardware_src) clean