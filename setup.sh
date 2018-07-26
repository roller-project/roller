sudo apt-get -y update
sudo apt-get install build-essential libssl-dev git curl screen
sudo curl -O https://storage.googleapis.com/golang/go1.9.7.linux-amd64.tar.gz
sudo tar -xvf go1.9.7.linux-amd64.tar.gz
sudo mv go /usr/local
sudo echo 'export PATH=$PATH:/usr/local/go/bin' > ~/.profile
source ~/.profile
git clone https://github.com/roller-project/roller.git roller
cd ./roller
sudo make geth