#!/bin/bash

set -e -u

VBoxManage -v >/dev/null 2>&1 || { echo >&2 "VirtualBox is required. Please install the latest version."; exit 1; }
vagrant -v >/dev/null 2>&1 || { echo >&2 "Vagrant is required. Please install the latest version."; exit 1; }
chef -v >/dev/null 2>&1 || { echo >&2 "Chef Development Kit is required. Please install the latest version."; exit 1; }
packer version >/dev/null 2>&1 || { echo >&2 "Packer is required. Please install the latest version."; exit 1; }

[[ $(vagrant plugin list) == *vagrant-vbguest* ]] || { vagrant plugin install vagrant-vbguest; }

[[ -d "provisioners/chef/berks-cookbooks" ]] || {
    chef exec berks vendor;
    mv berks-cookbooks provisioners/chef;
}

function help_text {
    echo "Usage: ./go <command>"
    echo ""
    echo "Available commands are:"
    echo "    amazon       Build an Amazon image"
    echo "    build        Build all images"
    echo "    deploy       Deploy image to local virtual machine"
    echo "    destroy      Destroy local virtual machine"
    echo "    virtualbox   Build a VirtualBox image"
}

function build_all {
    build_virtualbox
    build_amazon
}

function build_amazon {
    packer build -only amazon-ebs goagent.json
}

function build_virtualbox {
    rm go-agent.box || true
    packer build -only virtualbox-iso goagent.json
    vagrant box remove go-agent || true
    vagrant box add go-agent go-agent.box
}

[[ $@ ]] || { help_text; exit 1; }

case "$1" in
    build) build_all
    ;;
    amazon) build_amazon
    ;;
    virtualbox) build_virtualbox
    ;;
    deploy) vagrant up --no-provision
    ;;
    destroy) vagrant destroy -f
    ;;
    help) help_text
    ;;
esac