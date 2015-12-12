project_path = '/vagrant/src/eps-conduit'

bash 'Get and Install' do
  cwd project_path
  code <<-EOF
  go get ./... &&\
  go install &&\
  cp conduit.conf /vagrant/bin
  EOF
  environment 'GOPATH' => '/vagrant/'
end
