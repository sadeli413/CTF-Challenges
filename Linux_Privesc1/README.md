# Privesc lab

## Setup Instructions
1. Generate SSH keys
```
$ mkdir keys
$ ssh-keygen -q -t rsa -N '' -f keys/sadeli_rsa
```

2. Build the docker image (and give the image the name "devilsec"):
```
$ docker build . -t devilsec
```

3. Run the image in detatched mode:
```
$ docker run -h devilsec --name devilsec -dt devilsec
```

4. Get the IP address of the container.
```
$ docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' devilsec
```

5. SSH into the user sadeli
```
$ ssh sadeli@172.17.0.2 -i keys/sadeli_rsa
```

6. The order of flags goes:
    - sadeli.flag
    - 7StringsOvPasta.flag
    - tayari.flag
    - root.flag
