FROM jekyll/jekyll:stable

RUN apk add --no-cache --virtual build-dependencies build-base
RUN apk add --no-cache libxml2-dev libxslt-dev
RUN apk add --no-cache ruby-dev curl-dev zlib-dev yaml-dev
RUN gem install nokogiri
RUN gem install minima
RUN gem install jekyll-import
RUN gem install jekyll-minifier
RUN gem install jekyll-tagging
RUN gem install jekyll-paginate
RUN gem install webrick
EXPOSE 4000
ENV TZ Asia/Tokyo

ENTRYPOINT ["jekyll", "serve", "--config", "/srv/jekyll/source/_config.yml,/srv/jekyll/source/_config_dev.yml", "--watch", "--destination", "/srv/jekyll/staging", "--source", "/srv/jekyll/source", "--future"]