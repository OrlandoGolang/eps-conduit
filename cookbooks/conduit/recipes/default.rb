#
# Author:: Scott Crespo <scott@orlandods.com>
# Cookbook Name:: conduit
# Recipe:: default


# The order of recipes it important
include_recipe "conduit::golang"
include_recipe "conduit::hello_world"
include_recipe "conduit::build"
include_recipe "conduit::supervisor"
