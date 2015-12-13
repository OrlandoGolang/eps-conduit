#
# Author:: Scott Crespo <scott@orlandods.com>
# Cookbook Name:: conduit
# Recipe:: hello_world
#
# Installs eps-conduit-hello - a lightweight web app thats's useful for testing
# and debugging eps-conduit
#
# Clones hello world application and builds from source

include_recipe "conduit::supervisor"

ogo_path = '/vagrant/src/github.com/OrlandoGolang'
hello_path = File.join(ogo_path, 'eps-conduit-hello')

# ensure git client is installed in vagrant machine
package 'Install Git' do
  package_name 'git'
end


directory ogo_path do
  owner 'vagrant'
  group 'vagrant'
  mode '0755'
  action:create
end

# Copy ssh deploy key to .ssh directory on VM
cookbook_file '/home/vagrant/.ssh/eps-conduit-hello' do
  source 'eps-conduit-hello'
  owner 'vagrant'
  group 'vagrant'
  mode '0600'
  action :create
end

# Vagrant doesn't handle manage ssh keys very well, especially when making
# git requests. We have to add this git_wrapper shell script to assist in ssh key
# management.
cookbook_file '/home/vagrant/.ssh/git_wrapper.sh' do
  source 'git_wrapper.sh'
  owner 'vagrant'
  group 'vagrant'
  mode '0755'
  action :create
end

# Clone eps-conduit-hello and apply git_wrapper for request
git "Clone eps-conduit-hello" do
  repository 'git@github.com:OrlandoGolang/eps-conduit-hello.git'
  action :sync
  user 'vagrant'
  group 'vagrant'
  ssh_wrapper '/home/vagrant/.ssh/git_wrapper.sh'
  destination hello_path
end

# Run go get and go install
bash 'Build eps-conduit-hello' do
  cwd hello_path
  code <<-EOF
  go get ./... && \
  go install
  EOF
  environment 'GOPATH' => '/vagrant/'
end

# Restart hello world processes managed by supervisor
bash 'Reload Supervisor Process' do
  code <<-EOF
  supervisorctl restart hello1 && \
  supervisorctl restart hello2
  EOF
end
