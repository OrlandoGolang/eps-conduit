# -*- mode: ruby -*-
# vi: set ft=ruby :
Vagrant.configure(2) do |config|

  config.vm.box = "ubuntu/trusty64"
  # port forwarding
  config.vm.network "forwarded_port", guest: 8000, host: 8000
  # disabled default sync location until project directories are created.
  config.vm.synced_folder './', '/vagrant', disabled: true
  # update packages
  config.vm.provision :shell, :inline => "sudo apt-get update --fix-missing"

  # run chef scripts to additionally provision the virtual machine
  config.vm.provision :chef_solo do |chef|
    chef.cookbooks_path = ['./cookbooks',]
    chef.add_recipe "conduit"
  end
  # now sync the project folder after chef scrips complete.
  config.vm.synced_folder '.', '/vagrant/src/eps-conduit'
end
