# Clones hello world application and builds from source
include_recipe "conduit::supervisor"

ogo_path = '/vagrant/src/github.com/OrlandoGolang'
hello_path = File.join(ogo_path, 'eps-conduit-hello')

package 'Install Git' do
  package_name 'git'
end

directory ogo_path do
  owner 'vagrant'
  group 'vagrant'
  mode '0755'
  action:create
end

cookbook_file '/home/vagrant/.ssh/eps-conduit-hello' do
  source 'eps-conduit-hello'
  owner 'vagrant'
  group 'vagrant'
  mode '0600'
  action :create
end

cookbook_file '/home/vagrant/.ssh/git_wrapper.sh' do
  source 'git_wrapper.sh'
  owner 'vagrant'
  group 'vagrant'
  mode '0755'
  action :create
end

git "Clone eps-conduit-hello" do
  repository 'git@github.com:OrlandoGolang/eps-conduit-hello.git'
  action :sync
  user 'vagrant'
  group 'vagrant'
  ssh_wrapper '/home/vagrant/.ssh/git_wrapper.sh'
  destination hello_path
end

execute 'Build eps-conduit-hello' do
  cwd hello_path
  command 'go get ./... && go install'
end

execute 'Reload Supervisor Process' do
  command <<-EOF
  supervisorctl restart hello-8001 && \
  supervisorctl restart hello-8002
  EOF
end
