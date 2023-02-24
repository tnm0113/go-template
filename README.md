# go-template

**go-template** is a template project for VQ2 Backend services written in Golang
## SET UP DEV ENVIRONMENT
### Install Go lang
```bash
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf setup/go1.19.4.linux-amd64.tar.gz
```
### Set variable environment on bash profile
```bash
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bashrc
echo "export GOPROXY=http://172.31.252.188:8081/repository/go-proxy/" >> ~/.bashrc
echo "export GOSUMDB=off" >> ~/.bashrc
echo "export GOPRIVATE=192.168.205.151" >> ~/.bashrc
echo "export GOINSECURE=192.168.205.151" >> ~/.bashrc
source ~/.bashrc
```

### Create `.netrc` file for installing local module
```bash
touch ~/.netrc
chmod 600 ~/.netrc
```
- Create Gitlab private access token: Go to Gitlab → User Profile → Access Tokens → Create new token with full access
- Set content for `.netrc` file with following `machine 192.168.205.151 login USERNAME_HERE password TOKEN_HERE`
### Install go tools
```bash
go install -v golang.org/x/tools/gopls@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install -v github.com/josharian/impl@latest
go install -v honnef.co/go/tools/staticcheck@latest
```
### Install protoc
```bash
sudo unzip setup/protoc-3.15.8-linux-x86_64.zip -d /usr/local/ && sudo chmod +x /usr/local/bin/protoc
cp setup/protoc-gen-go $(go env GOPATH)/bin/protoc-gen-go
```
### Install Go extension for VSCode
Install Go extension golang.Go-0.36.0.vsix in setup folder, then restart VSCode

## DEVELOPMENT GUIDELINE

See [`here`](docs/DevelopmentGuideLine.md)

