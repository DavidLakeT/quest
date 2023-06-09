#!/bin/bash

set -xeuf -o pipefail

if [[ -v CI ]]; then
    SUDO=""
    apt-get update
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends tzdata
    echo "en_US UTF-8" > /etc/locale.gen
else
    SUDO="sudo"
fi

# Install Go
${SUDO} apt update
${SUDO} apt-get install -y curl tar
${SUDO} apt-get update

${SUDO} rm -rf /usr/local/go
wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
${SUDO} tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
echo 'export GOPATH=$HOME/go' >> /etc/profile
echo 'export PATH=$PATH:$GOPATH/bin' >> /etc/profile
source /etc/profile
rm go1.20.4.linux-amd64.tar.gz

sudo apt-get install -y postgresql-client

# Install docker
if ! type "docker" > /dev/null; then
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | ${SUDO} gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | ${SUDO} tee /etc/apt/sources.list.d/docker.list > /dev/null

    ${SUDO} apt update

    apt-cache policy docker-ce

    if [[ -v CI ]]; then
        echo "Skipping creation of docker group"
    else
        ${SUDO} apt install -y docker-ce
        ${SUDO} usermod -aG docker ${USER}
    fi
else
    echo 'docker already installed'
fi

# Install docker-compose
if ! type "docker-compose" > /dev/null; then
    ${SUDO} curl -SL https://github.com/docker/compose/releases/download/v2.18.1/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose
    ${SUDO} chmod +x /usr/local/bin/docker-compose
else
    echo 'docker-compose already installed'
fi

# Install NodeJS

curl -sL https://deb.nodesource.com/setup_18.x -o nodesource_setup.sh
${SUDO} bash nodesource_setup.sh
${SUDO} apt-get install nodejs
rm -f nodesource_setup.sh
