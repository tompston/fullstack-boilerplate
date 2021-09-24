# NGINX setup guide

Current setup doesn't use nginx as a docker container, because i could't figure out how to add ssl certificates.
These are the steps i took to set up and connect the domain name and nginx.

1.  Buy a domain

2.  After buying the domain you also need to change some other stuff. If using Linode, [this](https://youtu.be/bOLNq4tkwj8) explanation wil help.

3.  change all of the YOUR_DOMAIN placeholders in nginx.conf to... your domain. duh

4.  After the tutorial, check [this](https://youtu.be/sPosn1XjbMs) guide for setting automatic ssl certificates.

## steps mentioned in the certbot tutorial + more

    sudo apt install python3-certbot-nginx  # install certbot
    cd /etc/nginx                           # cd into the nginx config
    sudo cp nginx.conf nginx.txt            # create a backup for the main nginx file
    sudo nano nginx.conf                    # copy the stuff that is written in the current ./nginx.conf file
    sudo nginx -t                                           # test to check if the nginx is working fine + sytax errors
    sudo certbot --nginx -d YOUR_DOMAIN -d www.YOUR_DOMAIN  # generate ssl certs
    sudo cp nginx.conf nginx-cert-backup.txt                # backup the new generated nginx.conf file again

## Some common nginx commands

    sudo systemctl status nginx     # check status
    sudo nginx -t                   # test to check if the nginx is working fine + sytax errors
    sudo systemctl enable nginx     # start nginx on boot
    sudo systemctl stop nginx       # stop
    sudo systemctl restart nginx    # restart
    sudo systemctl reload nginx     # reload
