import os
from time import sleep
from fabric.api import env, run, sudo
from fabric.context_managers import cd
from fabric.contrib.files import append

env.use_ssh_config = True
env.password = "barc3lona1"


def add_key(key_file_path):
    # authorized_key_path = '~/.ssh/authorized_keys'
    key_file_path = os.path.expanduser(key_file_path)

    if not key_file_path.endswith('pub'):
        raise Exception("Not valid public key")

    with open(key_file_path) as f:
        # run('mkdir -p ~/.ssh && chmod 700 ~/.ssh')
        key_file_text = f.read()
        append('~/.ssh/id_rsa.pub', key_file_text)
        # append(authorized_key_path, key_file_text)


def install_deps_one():
	sudo('apt-get install -y \
            vim \
            build-essential \
            curl \
            autoconf \
            automake \
            wget \
            libtool \
            git \
            openjdk-7-jdk \
            pkg-config')


def install_deps_two():
	sudo('apt-get install -y \
            openssl \
            ssh \
            zlib1g-dev \
            libssl-dev \
            libreadline-dev \
            libyaml-dev \
            libsqlite3-dev \
            sqlite3 \
            libxml2-dev \
            libxslt1-dev \
            libcurl4-openssl-dev \
            python-software-properties \
            libffi-dev \
            libtool \
            libgdbm-dev \
            libncurses5-dev \
            bison \
            libmp3lame-dev yasm libx264-dev')


def clean():
    sudo('apt-get clean')


def install_youtube_dl():
    sudo('apt-get install -y youtube-dl')


def install_ffmpeg():
    """This will take a while"""
    run("curl http://ffmpeg.org/releases/ffmpeg-2.6.3.tar.bz2 > ffmpeg-2.6.3.tar.bz2 ")
    run("tar -xvf ffmpeg-2.6.3.tar.bz2 ")
    with cd("ffmpeg-2.6.3"):
        run("ls -l")
        run("./configure --enable-gpl --enable-libx264")
        run("make")
        sudo("make install")
        sudo("ldconfig")


def start_redis():
    sudo('redis-server /etc/redis.conf --daemonize yes')


def install_redis():
    run("wget http://download.redis.io/redis-stable.tar.gz")
    run("tar xvzf redis-stable.tar.gz")
    with cd("redis-stable"):
        run("make")
        sudo("make install")

    start_redis()


def install_rethinkdb():
    sudo('echo "deb http://download.rethinkdb.com/apt $DISTRIB_CODENAME main" | tee /etc/apt/sources.list.d/rethinkdb.list')
    sudo('wget -qO- http://download.rethinkdb.com/apt/pubkey.gpg | apt-key add -')
    sudo('apt-get update')
    sudo('apt-get y install rethinkdb=2.\*')
    sudo('cp /etc/rethinkdb/default.conf.sample /etc/rethinkdb/instances.d/instance1.conf')
    sudo('/etc/init.d/rethinkdb restart')


def install_rvm():
    """Switching to golang will be simpler :)"""
    run("gpg --keyserver hkp://keys.gnupg.net --recv-keys D39DC0E3 BF04FF17")
    run("curl -sSL https://get.rvm.io | bash -s stable")
    run("rvm install 2.2.0")
    run("rvm use 2.2.0 --default")
    run("gem install bundler --no-ri --no-rdoc")


def install_supervisor():
    sudo("apt-get install -y supervisor")


def export_foreman():
    with cd('~/pelvis'):
        sudo('foreman export supervisord /etc/supervisor/conf.d -a pelvis -u adrian')


def clone_app():
    run("git clone git@github.com:ammoses89/pelvis.git")


def bundle_install():
    with cd('~/pelvis'):
        run("bundle install")


def setup_db():
    with cd('~/pelvis'):
        run("rake rethinkdb:setup_db")


def start_app():
    sudo("service supervisor stop")
    sleep(3)
    sudo("service supervisor start")


def pull_app():
    with cd('~/pelvis'):
        run("git pull")


def setup_server():
    install_deps_one()
    install_deps_two()
    clean()
    install_youtube_dl()
    install_ffmpeg()
    install_redis()
    install_rethinkdb()
    install_rvm()
    clone_app()


def deploy():
    """ Will deploy and restart server"""
    pull_app()
    bundle_install()
    start_app()
