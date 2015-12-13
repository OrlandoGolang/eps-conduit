#
# Author:: Scott Crespo <scott@orlandods.com>
# Cookbook Name:: conduit
# Recipe:: supervisor
#
# Installs supervisor - a process management tool
# This installation runs eps-conduit-hello on multiple processes and ports

package 'Python Setuptools' do
  package_name 'python-setuptools'
end

execute 'Install Supervisor' do
  command 'easy_install supervisor'
end

cookbook_file '/etc/supervisord.conf' do
  source 'supervisord.conf'
  owner 'vagrant'
  group 'vagrant'
  mode '0755'
  action :create
end

cookbook_file '/etc/init/supervisord.conf' do
  source 'init_supervisor.sh'
  owner 'vagrant'
  group 'vagrant'
  mode '0755'
  action :create
end

directory '/var/log/supervisord' do
  owner 'root'
  group 'root'
  mode '0755'
  action :create
end

service "supervisord" do
  start_command "supervisord -c /etc/supervisord.conf ; true"
  action [:enable, :start]
end
