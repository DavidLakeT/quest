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
