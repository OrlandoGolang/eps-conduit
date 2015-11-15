# download golang 1.5.1 binaries
remote_file '/home/vagrant/go1.5.1.linux-amd64.tar.gz' do
  source 'https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz'
  mode '0655'
  action :create
end

# unpack golang binaries to /usr/local
bash 'install-golang' do
  cwd "/home/vagrant"
  code "sudo tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz"
end

# add Go env vars to bash profile
ruby_block "insert_line" do
  block do
    file = Chef::Util::FileEdit.new("/home/vagrant/.bashrc")
    file.insert_line_if_no_match("^export GOPATH\=.vagrant", "export GOPATH=/vagrant")
    file.insert_line_if_no_match("^export GOROOT\=\/usr\/local\/go", "export GOROOT=/usr/local/go")
    file.insert_line_if_no_match("^export PATH\=.PATH..GOROOT\/bin", "export PATH=$PATH:$GOROOT/bin")
    file.insert_line_if_no_match("^alias projectdir.+", "alias projectdir='cd $GOPATH/src/project'")
    file.write_file
  end
end

# add Go env vars to bash profile
ruby_block "insert_line" do
  block do
    file = Chef::Util::FileEdit.new("/etc/profile")
    file.insert_line_if_no_match("^export GOPATH\=.vagrant", "export GOPATH=/vagrant")
    file.insert_line_if_no_match("^export GOROOT\=\/usr\/local\/go", "export GOROOT=/usr/local/go")
    file.insert_line_if_no_match("^export PATH\=.PATH..GOROOT\/bin", "export PATH=$PATH:$GOROOT/bin")
    file.write_file
  end
end

# Create paths for go files
gopath = '/vagrant'
godirs = ['bin', 'src', 'pkg']
srcdirs = ['eps-conduit',]

# create main golang directories (bin, pkg, src)
godirs.each do |path|
  fullpath = File.join(gopath, path)
  directory fullpath do
    owner 'vagrant'
    group 'vagrant'
    mode '0755'
  end
end

# create src directories
srcdirs.each do |path|
  fullpath = File.join(gopath, 'src', path)
  directory fullpath do
    owner 'vagrant'
    group 'vagrant'
    mode '0755'
  end
end
