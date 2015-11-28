execute 'Get Compile Daemon' do
  command 'go get github.com/githubnemo/compiledaemon'
end

execute 'Install Compile Daemon' do
  cwd "/vagrant/src/github.com/githubnemo/compiledaemon"
  command "go install"
end

ruby_block "insert_line" do
  block do
    file = Chef::Util::FileEdit.new("/home/vagrant/.bashrc")
    file.insert_line_if_no_match("^alias conduitd",
      "alias conduitd=\"CompileDaemon -directory=\'$GOPATH/src/eps-conduit\' -build=\'go install\' -command=\'$GOPATH/bin/eps-conduit\'\"")
    file.write_file
  end
end
