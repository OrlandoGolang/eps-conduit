#
# Author:: Scott Crespo <scott@orlandods.com>
# Cookbook Name:: conduit
# Recipe:: golang
#
# Installs Golang 1.5.1 and configures necessary environment variables

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

# Symlink the go command to /bin. This is mostly important for non-interactive
# provisioning of the vagrant machine when PATH variable does not include GOBIN
bash 'symlink go cmd to bin' do
  cwd "/bin"
  code "ln -s /usr/local/go/bin/go"
  not_if { File.exists?("/bin/go") }
end

# Copy custom bash profile and put it in the /etc directory
cookbook_file '/etc/profile_custom' do
  source "profile_custom"
  owner 'vagrant'
  group 'vagrant'
  mode '0600'
  action :create
end

# add Go env vars to bash profiles
# Add a line to the relevant bash profiles that sources our custom bash_profile
# (profile_custom)
profiles = ['/etc/profile', '/root/.bashrc', '/home/vagrant/.bashrc']
profiles.each do |profile|
  ruby_block "Insert lines in bash profiles" do
    block do
      file = Chef::Util::FileEdit.new(profile)
      file.insert_line_if_no_match("^source /etc/profile_custom", "source /etc/profile_custom")
      file.write_file
    end
  end
end

# Vagrant's root user has an annoying line in its bash profile that prevents
# non-interactive sessions from using the profile settings. This resource
# removes that line using sed so that root user's bash profile can be utilized
# during vagrant provisioning
bash 'remove interactive blocking' do
  code "sed -i '/.*z.*PS1.*return/d' /root/.bashrc"
end

# Create lists of directories
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

# Create src directory for. This will become the synced_folder specified in
# the Vagrantfile
srcdirs.each do |path|
  fullpath = File.join(gopath, 'src', path)
  directory fullpath do
    owner 'vagrant'
    group 'vagrant'
    mode '0755'
  end
end

# create github directory in src path
directory '/vagrant/src/github.com' do
  owner 'vagrant'
  group 'vagrant'
  mode '0755'
end
