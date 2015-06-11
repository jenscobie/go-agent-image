# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.network :private_network, ip: "192.168.0.3"

  config.vm.provider :virtualbox do |vb|
    vb.customize [ 'modifyvm', :id, '--memory', '2048' ]
  end

  config.vm.provision "chef_solo" do |chef|
    chef.cookbooks_path = "provisioners/chef/berks-cookbooks"
    chef.roles_path = "provisioners/chef/roles"
    chef.add_role("agent")
  end
end