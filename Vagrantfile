# -*- mode: ruby -*-
# vi: set ft=ruby :
Vagrant.configure(2) do |config|

  config.vm.box = "ubuntu/trusty64"
  # port forwarding
  config.vm.network "forwarded_port", guest: 80, host: 8085
  # disabled default sync location
  config.vm.synced_folder './', '/vagrant', disabled: true
  # sync according to standard golang environment & directory structure

  # update packages
  config.vm.provision :shell, :inline => "sudo apt-get update --fix-missing"

  # run chef scripts to additionally provision the virtual machine
  config.vm.provision :chef_solo do |chef|
    chef.cookbooks_path = ['./cookbooks',]
    chef.add_recipe "conduit"
  end
end
