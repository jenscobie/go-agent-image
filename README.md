# Go Continuous Delivery Agent Image

> [Go CD](http://www.go.cd/) agent machine image

## Motivation

To get the latest version of the Go CD agent setup with minimal effort by providing a machine image of a Go CD agent.

## Requirements

* [ChefDK](https://downloads.chef.io/chef-dk/)
* [Packer](https://www.packer.io/)
* [Vagrant](https://www.vagrantup.com/)
* [VirtualBox](https://www.virtualbox.org/wiki/Downloads)

## Installation

1. Install the requirements listed above
2. Run ```./go build```
3. Run ```./go deploy```

## Usage

    Usage: ./go <command>
    
    Available commands are:
        amazon       Build an Amazon image
        build        Build all images
        deploy       Deploy image to local virtual machine
        destroy      Destroy local virtual machine
        virtualbox   Build a VirtualBox image

## Author

Jen Scobie (jenscobie@gmail.com)