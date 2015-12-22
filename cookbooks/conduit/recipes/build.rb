#
# Author:: Scott Crespo <scott@orlandods.com>
# Cookbook Name:: conduit
# Recipe:: build
#
# Installs dependencies and eps-conduit

project_path = '/vagrant/src/eps-conduit'
config_path = File.join(project_path, 'conduit.conf')

# downloads, installs, and builds dependencies and eps-conduit
bash 'Get and Install' do
  cwd project_path
  code <<-EOF
  go get ./... &&\
  go install &&\
  cp conduit.conf /vagrant/bin
  EOF
  environment 'GOPATH' => '/vagrant/'
end

bash 'Symlink conduit.conf to /etc/' do
  cwd '/etc'
  code "ln -s #{config_path}"
  not_if { File.exists?(config_path) }
end
