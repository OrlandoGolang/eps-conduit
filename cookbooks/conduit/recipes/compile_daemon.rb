#
# Author:: Scott Crespo <scott@orlandods.com>
# Cookbook Name:: conduit
# Recipe:: compile_daemon
#
# Downloads and installs Compile Daemon
# Compile Daemon watches your project files and rebuilds+re-starts the project
#
# Project url: https://github.com/githubnemo/CompileDaemon

bash 'Get Compile Daemon' do
  code "go get github.com/githubnemo/compiledaemon"
  environment 'GOPATH' => '/vagrant/'
end

bash'Install Compile Daemon' do
  cwd "/vagrant/src/github.com/githubnemo/compiledaemon"
  code "go install"
  environment 'GOPATH' => '/vagrant/'
end
