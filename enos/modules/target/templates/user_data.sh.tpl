#! /bin/bash

sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/' /etc/ssh/sshd_config
service sshd restart
useradd -md /home/${username} -p $(perl -e 'print crypt("${password}", "pw")' $pw) ${username}
