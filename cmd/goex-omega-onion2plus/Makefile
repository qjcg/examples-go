out := $(notdir $(PWD))

ssh_user := root
ssh_host := omega-7961
ssh_login := $(ssh_user)@$(ssh_host)
ssh_options := -o 'PubkeyAcceptedAlgorithms +ssh-rsa' -o 'HostkeyAlgorithms +ssh-rsa'
ssh_cmd := ssh $(ssh_options) $(ssh_login)


.PHONY: all
all: clean build scp

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags '-s -w'

.PHONY: scp
scp:
	scp $(ssh_options) $(out) $(ssh_login):/tmp/

.PHONY: ssh
ssh:
	$(ssh_cmd)

.PHONY: clean
clean:
	rm -f $(out)
