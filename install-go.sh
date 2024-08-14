#!/bin/bash

# Atualiza o sistema
echo "Atualizando o sistema..."
sudo apt update && sudo apt upgrade -y

# Baixa o arquivo de instalação do Go
echo "Baixando o Go..."
curl -O https://dl.google.com/go/go1.23.0.linux-amd64.tar.gz

# Extrai o arquivo para /usr/local
echo "Extraindo o Go..."
sudo tar -C /usr/local -xzvf go1.23.0.linux-amd64.tar.gz

# Adiciona o Go ao PATH, se ainda não estiver no ~/.bashrc
if ! grep -q "export PATH=\$PATH:/usr/local/go/bin" ~/.bashrc; then
    echo "Configurando o PATH..."
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
else
    echo "O PATH já está configurado."
fi

# Atualiza o shell para aplicar as mudanças
echo "Atualizando o shell..."
source ~/.bashrc

# Verifica a instalação
echo "Verificando a instalação do Go..."
go version

echo "Instalação do Go concluída com sucesso!"
