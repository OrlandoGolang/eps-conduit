bash 'Get Compile Daemon' do
  code "go get github.com/githubnemo/compiledaemon"
  environment 'GOPATH' => '/vagrant/'
end

bash'Install Compile Daemon' do
  cwd "/vagrant/src/github.com/githubnemo/compiledaemon"
  code "go install"
  environment 'GOPATH' => '/vagrant/'
end
