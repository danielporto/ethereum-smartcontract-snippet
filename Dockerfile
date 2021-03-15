FROM ubuntu:20.04

# install needed packages
RUN apt update && DEBIAN_FRONTEND=noninteractive apt install -y \
               #   ansible require python3
               python3-pip \ 
               sudo \
               openssl \
               # required for build asdf python
               libssl-dev \ 
               curl \
               vim \
               zsh \
               openssh-client \
               rsync \
               jq \
               ansible \
               git \
               gnupg \
               unzip \
               # required for compile ethereum devtools
               software-properties-common \
               protobuf-compiler \
               && rm -rf /var/cache/apt/archives


# install solidity compiler, dependency to compile go tool - abigen
RUN add-apt-repository ppa:ethereum/ethereum -y && \
            apt update && \
            apt install -y solc 

# install packages
RUN pip3 install --upgrade pathlib2 \
                          enum34 \
                          plumbum \
                          eventlet \
                          # sensors
                          futures \
                          requests \
                          numpy \
                          pandas \
                          plumbum \
                          serial \
                          ioctl_opt \
                          pyserial \
                          flask \
                          # pysensors \ not yet compatible with alpine version
                          eventlet \
                          # zookeeper
                          kazoo \
                          ;

# ARG UID=1097
# RUN addgroup -S dporto && adduser --uid $UID -g dporto dporto && echo 'dporto ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
ARG UID=501
RUN useradd -m --uid=$UID -U dporto && echo 'dporto ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# make default alias for python
# RUN  ln -s /usr/local/lib/pyenv/versions/3.7.3/bin/python /usr/bin/python 
RUN sudo update-alternatives --install /usr/bin/python python /usr/bin/python3 10


USER dporto
# docker commands that run from now works as user $USER above inside the container
# speedup ansible with mitogen
# RUN curl -o /tmp/mitogen-0.2.9.tar.gz https://codeload.github.com/dw/mitogen/tar.gz/v0.2.9 \
#     && tar xvf /tmp/mitogen-0.2.9.tar.gz -C /home/dporto \
#     && mv /home/dporto/mitogen-0.2.9 /home/dporto/mitogen \
#     && chown -R dporto.dporto /home/dporto/mitogen 

# install ansible dependencies
# RUN ansible-galaxy install git+https://github.com/danielporto/ansible-sdkman.git
# RUN ansible-galaxy collection install crivetimihai.virtualization
# RUN ansible-galaxy install markosamuli.asdf

# install asdf
RUN git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.8.0 

# install JAVA plugin for asdf
# RUN PATH="$HOME/.asdf/bin:$HOME/.asdf/shims:$PATH" && \
    # asdf plugin-add java https://github.com/halcyon/asdf-java.git && \
    # asdf install java zulu-11.43.55 && \
    # asdf global java zulu-11.43.55

# install NODEJS plugin for asdf
# RUN PATH="$HOME/.asdf/bin:$HOME/.asdf/shims:$PATH" && \
    # asdf plugin-add nodejs https://github.com/asdf-vm/asdf-nodejs.git && \
    # bash -c '${ASDF_DATA_DIR:=$HOME/.asdf}/plugins/nodejs/bin/import-release-team-keyring' && \
    # asdf install nodejs 14.15.4 && \
    # asdf global nodejs 14.15.4

# install quorum-wizard to generate network configurations for quorum
# RUN PATH="$HOME/.asdf/bin:$HOME/.asdf/shims:$PATH" && \
    # npm install -g quorum-wizard@1.3.1 

# install GOLANG plugin for asdf
RUN PATH="$HOME/.asdf/bin:$HOME/.asdf/shims:$PATH" && \
    asdf plugin-add golang https://github.com/kennyp/asdf-golang.git && \
    asdf install golang 1.14.13 && \
    asdf global golang 1.14.13 && \
# install abigen (to generate contracts)
    go get -u github.com/ethereum/go-ethereum && \
    cd $(go env GOPATH)/src/github.com/ethereum/go-ethereum && make && make devtools



# install istanbul-tools for generating keys and templates
# RUN git clone https://github.com/ConsenSys/istanbul-tools.git $HOME/istanbul-tools &&\
#     PATH="$HOME/.asdf/bin:$HOME/.asdf/shims:$PATH" && \
#     cd $HOME/istanbul-tools && \
#     make

# # install geth (quorum ethereum client) to initialize the nodes, tag v20.10.0 
# RUN git clone https://github.com/ConsenSys/quorum.git $HOME/goquorum &&\
#     PATH="$HOME/.asdf/bin:$HOME/.asdf/shims:$PATH" && \
#     cd $HOME/goquorum && \
#     git checkout af7525189f2cee801ef6673d438b8577c8c5aa34 && \ 
#     make 

# install zsh
RUN curl -fsSL https://raw.githubusercontent.com/zimfw/install/master/install.zsh | zsh

# configure zsh
RUN PATH="$HOME/.asdf/bin:$HOME/.asdf/shims:$PATH" && \
    # echo "alias play='/usr/bin/ansible-playbook -i inventory.cfg'" > ~/.profile ; \
    echo "alias grep='egrep'" >> ~/.profile ; \
    # echo "alias ssh='ssh -F ~/.ssh/config_gsd_kaioken_dporto'" >> ~/.profile ; \
    echo ". ~/.profile" >> ~/.zshrc ; \
    echo "export PATH=$HOME/.asdf/bin:$HOME/.asdf/shims:$(go env GOPATH)/bin:$HOME/istanbul-tools/build/bin:$HOME/goquorum/build/bin:$PATH" >> ~/.zshrc ; \

    echo "Done"

WORKDIR /code

CMD /bin/zsh
