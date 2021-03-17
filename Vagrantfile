# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  config.vm.box = "generic/centos8"

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  #config.vm.network "public_network", :bridge => "enp181s0f3", :mac => "080027a407ba" 
  #config.vm.network "public_network", :bridge => "enp3s0", :mac => "080027a407ba" 
  config.vm.network "public_network", dev: "enp3s0", bridge: "enp3s0", mode: "bridge", mac: "080027a407ba" 
  
  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  config.vm.synced_folder ".", "/vagrant", type: 'virtualbox'

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  config.vm.provider "virtualbox" do |vb|
  #   vb.gui = true
     vb.cpus = 8
     vb.memory = "25480"
     vb.customize ["modifyvm", :id, "--vram","128" ]
     #vb.customize ["modifyvm", :id, "--nictype1", "virtio" ]
  end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Enable provisioning with a shell script. Additional provisioners such as
  # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
  # documentation for more information about their specific syntax and use.
  config.vm.provision "shell", inline: <<-SHELL
      dnf update -y
      # install base packages
      dnf install -y epel-release
      yum groupinstall -y  "Server with GUI"
      yum groupinstall -y  "@Development tools"
      yum install -y  \
                  openssl-devel \
                  ncurses-devel  \
                  elfutils-libelf-devel  \
                  bc \
                  rsync \
                  vim \
                  python2-pip \
                  sudo \
                  tmux \
                  htop\
                  iperf3\
                  lm_sensors \
                  grubby \
                  net-tools \
                  libcurl\
                  bzip2 \
                  bzip2-devel\
                  zlib-devel \
                  readline-devel \
                  sqlite \
                  sqlite-devel \
                  openssl-devel \
                  xz \
                  xz-devel \
                  libffi-devel \
                  findutils\
                  mutt\
                  psmisc\
                  htop\
                  ipmitool\
                  redhat-lsb-core\
                  fail2ban-firewalld\
                  jq\
                  kexec-tools\
                  rasdaemon  \
                  chrony\
                  ntpstat \
                  tigervnc-server \
                  xrdp  \
                  git \
                  zsh
      
      # systemctl set-default graphical
      # install public key
      su -c 'echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC+5Wfw9vKHzBQ0yqzZh4CDHguWfMBCURU30dZRDf0D9NcKPyXt8PEaeJIyTUlA+3a7Evo3phz59T+XtAt6K2iCJcshWPBfoxlA8F1pyfckm01rWCO589KfG+pUkcVt1tyinHoTUAka6IHMYZd3iKtMMgO1Iti7Lj1HIVjbXdAEdjQzFD8ix4qK1+/sRRaTFbPNwGZrGLXyzX2adsahyEwH1I0RH2yDKaHCKYNZS73yP5zF3Ao7uIZeFMPtpqzCoJYKxDqnFyxrA91u761bHasorumgSK2Jn+Or3S35iMQZjszoA9FAOlGlSM3JUxPovOpm0TcQMkZA39uX02L1mbLL vagrant" >> /home/vagrant/.ssh/authorized_keys' vagrant

      # configure remote desktop
      echo "exec gnome-session" >> /etc/xrdp/xrdp.ini
      systemctl restart xrdp
      systemctl enable xrdp --now
      firewall-cmd --permanent --add-port=3389/tcp
      firewall-cmd --reload

      # install docker
      dnf config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo
      dnf install docker-ce --nobest -y
      systemctl start docker
      systemctl enable docker

      # install docker compose
      curl -L "https://github.com/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
      chmod +x /usr/local/bin/docker-compose
      usermod -aG docker vagrant

      # install zshell addons
      su -c 'curl -fsSL https://raw.githubusercontent.com/zimfw/install/master/install.zsh | zsh' vagrant
      echo "zmodule danielporto/wd -s 'wd.plugin.zsh' -f '.'"  >> /home/vagrant/.zimrc
      chown vagrant /home/vagrant/.zimrc
      su -c "zsh -c 'zimfw install' " vagrant

      # install visual studio code
      rpm --import https://packages.microsoft.com/keys/microsoft.asc
      sh -c 'echo -e "[code]\nname=Visual Studio Code\nbaseurl=https://packages.microsoft.com/yumrepos/vscode\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/vscode.repo'
      dnf check-update
      dnf install -y code

      # install gitkraken
      dnf install -y https://release.gitkraken.com/linux/gitkraken-amd64.rpm

      # install wireguard clinet
      yum config-manager --set-enabled PowerTools
      yum copr enable -y jdoss/wireguard
      yum -y install wireguard-dkms wireguard-tools
      
      # disable firewall
      # systemctl disable firewall
      # systemctl stop firewall

  SHELL
end
