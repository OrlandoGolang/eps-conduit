#
# Author:: Scott Crespo <scott@orlandods.com>
# Cookbook Name:: conduit
# Recipe:: supervisor
#
# Installs supervisor - a process management tool
# This installation runs eps-conduit-hello on multiple processes and ports

# Ensure setuptools is installed beause we'll need to use the easy_install
# package manager
package 'Python Setuptools' do
  package_name 'python-setuptools'
end

# Install supervisor
execute 'Install Supervisor' do
  command 'easy_install supervisor'
end

# Upload custom supervisor configuration file from cookbook
cookbook_file '/etc/supervisord.conf' do
  source 'supervisord.conf'
  owner 'vagrant'
  group 'vagrant'
  mode '0755'
  action :create
end

# Init script to have supervisor start and run application when VM booted
cookbook_file '/etc/init/supervisord.conf' do
  source 'init_supervisor.sh'
  owner 'vagrant'
  group 'vagrant'
  mode '0755'
  action :create
end

# Ensure log directory is created
directory '/var/log/supervisord' do
  owner 'root'
  group 'root'
  mode '0755'
  action :create
end

# start supervisord. the "; true" snippet ensures this command always returns true
# Othwersie, chef will abort if supervisord is already running.
service "supervisord" do
  start_command "supervisord -c /etc/supervisord.conf ; true"
  action [:enable, :start]
end
