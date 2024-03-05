## Install golang on raspi

1. Download golang package
```bash
wget https://dl.google.com/go/go1.13.7.linux-armv6l.tar.gz -O go.tar.gz
```

2. Extract
```bash
sudo tar -C /usr/local -xzf go.tar.gz
```

3. Add paths
```bash
nano ~/.bashrc
```

4. Add on end of .bashrc file
```bash
export GOPATH=$HOME/go
export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
```
