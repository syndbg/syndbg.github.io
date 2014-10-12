Date: 2014-10-12
Title: Linux: Python(Django) Provisioning and Deployment
Tags:
Slug: linux-python-provisioning-and-deployment


In the following tutorial/guide/article I'll explain how to provision and deploy a Django app.


Before all you should have:

* A domain (GoDaddy, Namecheap, Hover, ICN.bg, superhosting and etc)
* VPS/Cloud host with SSH (digitalocean.com is my pick)



## Domain, VPS, Cloud setup

**If you bought a domain, but didn't configure it yet:**

You should change the DNS Name Servers (NS in short) to DigitalOcean's or your VPS/Cloud provider's (if they have such).

To do so:

* Sign in your domain provider's site
* Change the NS to ns1.digitalocean.com, ns2.digitalocean.com, ns3.digitalocean.com or your VPS/Cloud provider's.
* Save and forget the domain provider's site, hopefully.


**Now in your VPS/Cloud provider's website, DigitalOcean in my case:**

* Login
* Go to the DNS page and add your domain name there.
* (If you don't have one yet) Create a container with latest Ubuntu/Debian and copy the container IP address
* Create an A record with your |domain name you bought| and |container IP address that you copied|
* (Optional) you can also point every sub-domain to your container, to do so create a CNAME | * | and | @ |.


In the end you should have something like

Record Type | Name |Value
--- | --- | ---
A | @ | Your Container IP
CNAME| * | @
NS| ns1.digitalocean.com |
NS| ns1.digitalocean.com |
NS| ns1.digitalocean.com |


* CNAME is optional!


**Now a zone file should be generated for your domain. While the file is generated, we can move to the provisioning part.**


## Provisioning

**Requirements:**

* Python 2/3 (Python 3 is preferred nowadays)
* nginx (reverse proxy server you should use to serve static files *fast*)
* virtualenv (the python tool for isolated virtual environments)
* pip (the python package manager)
* gunicorn (python HTTP server you should be using in production)


**1. Connect through SSH to your container**

To do so - `ssh root@containerIP`

You should be asked for password. Enter the password you set during container creation or the password you had automatically generated for your container.

After the first successful login, you should use SSH keys.


* To generate a SSH keys pair, follow [Github's tutorial](https://help.github.com/articles/generating-ssh-keys/) until Step 3.
* Then just copy your public key with `ssh-copy-id root@containerIP`
* You should see a prompt that says your identity is unknown, do you want to continue: type yes.
* Logout from your SSH and try login again with `ssh root@containerIP`. Now you shouldn't be asked for a password at all!


**2. Install the above requirements. In Debian/Ubuntu it should look like this. (else, replace apt-get,
with your distro's package manager):**

```shell
sudo apt-get update
sudo apt-get install python3, nginx, python3-pip
sudo pip3 install virtualenv
```

**3. Create a sudo user:**

You should never ever use root to deploy your apps. Instead create new sudo users for each app you deploy.

To create an user:
```shell
useradd -m -s /bin/bash USERNAME
```
* `-m` creates a home folder for the user, ex. /home/USERNAME, `-s` sets the new user's preferred shell.

To make him a sudo user:
```shell
usermod -a -G sudo USERNAME
```
* `-a` tells to append, `-G` specifies a group, in our case that's group `sudo`.



**4. Create the folder structure:**

This is what I recommend. If you know what you're doing, you can use your own, but the DEPLOYMENT steps will differ a lot!

* A sites folder that contains all our domain names
* A database, source, static and virtualenv folder for each domain name.


`database` - our database usage instructions and files will be there

`source` - our app code for that domain code will be there. In our case, our Django app will be there

`static` - the HTML/CSS/JS files that our app needs. These are the files that nginx will serve.

`virtualenv` - our python2/3 virtualenv with the packages we need, one of them is gunicorn, the others are what you have in your PIP requirements.txt file.


**Our folder structure should look like this:**

```
/home/USERNAME
├── sites
│   └── your-domain.com
│       ├── database
│       ├── source
│       ├── static
│       └── virtualenv
```

So create the above folders.
```shell
mkdir -p sites/your-domain.com/database
cd sites/your-domain.com
mkdir source
mkdir static
mkdir virtualenv
```


## Deployment:

If you disconnected from your container, connect again.

By now your domain name should be correctly pointing to your container's IP address.

Try now to connect with `ssh yourNewUsername@domain-name.com`.
If it still fails, consider contacting the domain provider and VPS/Cloud hosting provider supports.
It's probably a problem on their ends.


If everything is good now, continue.


**1. Clone your code:**


*If you have your code hosted in Github:*

* Install git on your container -`sudo apt-get install git`
* Clone your app's github repo in `sites/your-domain.com/source` - `git clone git@github.com:yourGitUsername/yourGithubRepoName.git`


*If you've downloaded your app to your local PC from git, SVN, hg or wherever:*

* In your local PC, go to your app's folder, ex. `cd ~/apps/django-app`
* From your local PC, copy the folder contents to your container's `sites/your-domain.com/source` - `scp -r * yourNewUsername@domain-name.com:~/sites/domain-name.com/source/`


**2. Create a Python virtual environment:**

Replace `python_version` with python2.7 or python3, depending on what you need for your app.

* Go to `sites/domain-name.com/`

* Create a virtualenv - `virtualenv --python=python_version virtualenv/`

You should have your virtualenv folder populated with the virtualenv you just created.

At last, activate the virtualenv `source virtualenv/bin/activate`


**3. Install your app's requirements**

I assume you're still in the virtualenv.

*If you have a requirements file* - `pip install -r requirements.txt`

*If you don't, manually install every app like* -`pip install app-name`



**install gunicorn, the HTTP server we'll use:**

`pip install gunicorn`


*Finally, see if everything you need, incl. gunicorn is installed * - `pip freeze`

For me that is:
```shell
% pip freeze
Django==1.7
gunicorn==19.1.1
```


**3. Adjust our Django app settings.py to our deploy folder structure**

* Go to `sites/domain-name.com/source/app_name/`
* Modify `settings.py` - `vim settings.py`

You should modify the following lines

```python
import os
BASE_DIR = os.path.abspath(os.path.dirname(os.path.dirname(__file__)))

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = False
TEMPLATE_DEBUG = DEBUG

# Needed when DEBUG=False
ALLOWED_HOSTS = ['domain-name.com']
## ^ However for me it doesn't work all the time,
## so it's ok to just replace it with ['*']

DATABASES = {
    'default': {
        # For other DBs 'django.db.backends.X', X = 'postgresql_psycopg2', 'django_mongodb_engine', 'mysql', 'sqlite3' or 'oracle'.
        'ENGINE': 'django.db.backends.sqlite3',
        'NAME': os.path.join(BASE_DIR, '../database/db.sqlite3'), # or DB name
        # If using sqlite3, delete the following lines
        'USER': 'db_user',
        'PASSWORD': 'db_password',
        # Empty for localhost through domain sockets or '127.0.0.1' for localhost through TCP.
        'HOST': '',
        'PORT': '',  # Empty for default
    }
}

STATIC_ROOT = os.path.join(BASE_DIR, '../static')
```

**4. Configure nginx**

* Create an nginx settings file for our domain. - `vim /etc/nginx/sites-available/domain-name.com`

```shell
server {
    listen 80;
    server_name domain-name.com;

    location /static {
        alias /home/USERNAME/sites/domain-name.com/static;
    }

    # We'll use UNIX domain socket's that can be used by nginx and gunicorn
    location / {
        proxy_set_header Host $host;
        proxy_pass http://unix:/tmp/domain-name.com.socket/
    }

}
```

** 5. Create an Upstart file**

This will make sure our nginx and gunicorn servers are running automatically.

So let's create it - `vim /etc/init/domain-name.com.conf`
```shell
description "Our gunicorn server for domain-name.com"

start on net-device-up
stop on shutdown

respawn

setuid yourNewUsername
chdir /home/yourNewUsername/sites/domain-name.com/source

exec ../virtualenv/bin/gunicorn --bind unix:/tmp/domain-name.com.socket app-name.wsgi:application
```


What did we just write?

* `description` - is a description, ok.
* `start on net-device-up` - starts when we have internet.
* `stop on shutdown` - stops on shutdown, as we want to.
* `respawn` - makes sure to keep it alive. When the process dies it will automatically try to start it again.
* `setuid` - makes the process run as our username we use to deploy our app.
* `chdir` - sets the working dir to our domain-name.com's source folder.
* `exec` - lauches our gunicorn HTTP server and binds it to the Unix domain socket we used in our nginx config.


**Important!** If you have a Django app, you would have a `wsgi.py` file, that is Python's primary web deployment platform.
Flask and other web frameworks have that too.


To elaborate more, WSGI - Webserver Gateway Interface. And the `wsgi.py` file exposes the WSGI callable as a module-level variable named `application`.
We use that as `app-name.wsgi:application`


**6. Last touches before everything is running**

So assuming we wrote everything correctly, last thing we need is:

* Source in your domain-name.com's virtualenv - `source virtualenv/bin/activate`
* To migrate our database - for Django earlier than 1.7 - `python manage.py syncdb -all;python manage.py migrate --fake`.
For Django 1.7 - `python manage.py migrate`
* Generate staticfiles `python manage.py collectstatic`


**7. Run everything!**

* Start/restart nginx - `sudo service nginx restart`
* Start/restart the gunicorn job - `sudo restart gunicorn-domain-name.com`
* Open a browser on your local PC and try to open `Your container IP address` and `domain-name.com`.
Both should show your app that you just deployed.


Congrats!
