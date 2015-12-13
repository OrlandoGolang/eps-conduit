#
# Author:: Scott Crespo <scott@orlandods.com>
# Cookbook Name:: conduit
# Recipe:: default

include_recipe "conduit::golang"
include_recipe "conduit::hello_world"
include_recipe "conduit::compile_daemon"
include_recipe "conduit::build"
include_recipe "conduit::supervisor"
