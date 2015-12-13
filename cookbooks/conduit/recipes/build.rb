#
# Author:: Scott Crespo <scott@orlandods.com>
# Cookbook Name:: conduit
# Recipe:: build
#
# Installs dependencies and eps-conduit

project_path = '/vagrant/src/eps-conduit'

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
