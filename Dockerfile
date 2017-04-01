FROM ubuntu:14.04

RUN apt-get update

RUN apt-get install -y \
 vim \
 build-essential \
 curl \
 autoconf \
 automake \
 wget \
 libtool \
 git \
 openjdk-7-jdk \
 pkg-config

RUN apt-get install -y \
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
  libmp3lame-dev \ 
  yasm \ 
  libx264-dev

RUN apt-get clean

# youtube-dl
RUN apt-get install -y youtube-dl 

# ffmpeg
RUN curl http://ffmpeg.org/releases/ffmpeg-2.6.3.tar.bz2 > /ffmpeg-2.6.3.tar.bz2 
RUN tar -xvf /ffmpeg-2.6.3.tar.bz2 
WORKDIR /ffmpeg-2.6.3
RUN ./configure --enable-gpl --enable-libx264
RUN make
RUN make install
RUN ldconfig

WORKDIR /

# RVM
RUN gpg --keyserver hkp://keys.gnupg.net --recv-keys D39DC0E3 BF04FF17
RUN curl -sSL https://get.rvm.io | bash -s stable
RUN /bin/bash -l -c "rvm install 2.2.0"
RUN /bin/bash -l -c "rvm use 2.2.0 --default"
RUN /bin/bash -l -c "gem install bundler --no-ri --no-rdoc"

WORKDIR /app
ADD . /app

RUN /bin/bash -l -c "bundle install"

EXPOSE 5000 
EXPOSE 28015 
EXPOSE 8080

CMD /bin/bash -l -c "/app/start_service.sh"

