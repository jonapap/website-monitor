# Website Monitor

## Table of Contents
  * [Overview](#overview)
    + [Website Saver](#website-saver)
    + [Website Checker](#website-checker)
    + [Website Deleter](#website-deleter)
  * [Docker Setup](#docker-setup)
    + [Prerequisites](#prerequisites)
    + [Usage](#usage)

## Overview
Set of three utilities that together can monitor if a website, or part of it, gets modified. To obtain information about the websites, Selenium is used to control a Firefox browser. See instructions below in the [Docker Setup](#docker-setup) section on how to setup Website Monitor along with Selenium.

### Website Saver
```shell
websitesaver [-website website] [-selector cssSelector] 
```

Saves the source code of the specified website to a file in the directory *./savedWebsite*. If **cssSelector** is specified, only the part of the website specified by the CSS selector will be saved.  
If **website** is not specified, the program will start in interactive mode and ask for the information.

The information about the website is saved in a JSON format. To prevent issues with characters that can't be in a filename and to prevent name collisions, the filename of the file is obtained by combining the URL of the website with the CSS selector, and transforming this afterward to an MD5 hash.

### Website Checker
```shell
websitechecker
```

Goes through each website in the *./savedWebsite* directory, gets the current source code of that website (or part of the website), and compares the saved copy to the current copy.  
After it compared all the website and if at least one website changed, it will send an email to the user alerting about the list of websites that changed. The email configuration is done as follow:

* If it is running in a Docker container, it will use this list of environment  variables:
    - **MAIL_from**: Email address to send from.
    - **MAIL_to**: Email address to send to.
    - **MAIL_host**: Name of the SMTP server to use. If you are not sure what to use here and have a GMail account, checkout how to use [Google's servers](https://support.google.com/a/answer/176600?hl=en).
    - **MAIL_port**: Port to use for SMTP server. The server need to use TLS. In most cases, this should be **587**.
    - **MAIL_username**: Username for login on the SMTP server
    - **MAIL_password**: Password for login on the SMTP server
* Else, if we are not in a Docker container, it will use a configuration file in *config/mail.ini*. If it does not exist, it will be created with default values and the user will be asked to change them. The parameters are similar to the ones used with Docker. 
    

### Website Deleter

```
websitedeleter
```

Can only run interactively. Lists all the saved websites in *./savedWebsites* and ask the user which one to delete.

## Docker Setup

For a quick and easy setup, it is recommended to use the Docker image along with Docker Compose. The Docker image contains a compiled version of the program, and has a crontab entry to run `websitechecker` daily. The Docker Compose config included in this repo will take care of installing and running Selenium.

### Prerequisites

First, make sure you have Docker along with Docker Compose installed on your computer.

* [Windows](https://docs.docker.com/windows/started)
* [OS X](https://docs.docker.com/mac/started/)
* [Linux](https://docs.docker.com/linux/started/)

### Usage
First, modify the file `docker-compose.yml` to specify the environment variables **MAIL_\***. See the [Website Checker](#website-checker) for the meaning of each variable. Afterward, run this command while you are inside the same directory as `docker-compose.yml`:
```shell
docker-compose up -d
```
This will start two containers, Website Monitor and Selenium, and make sure they can talk to each other.

Now, its time to add websites to monitor. To do so, simply run this:
```
docker exec -it website-monitor ./websitesaver
```
Simply follow the instructions and this will save a copy of the website's source code. Rerun the command for each website to add (or update).

Now, everything is done! Simply wait and once a day, you should be alerted if any of the websites you added changed. An email will be sent only if a website has changed.
