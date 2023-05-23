FROM bayesimpact/react-base
ADD . /usr/app
WORKDIR /usr/app
RUN npm install
RUN npm run build
RUN npm install -g serve
ENTRYPOINT ["/usr/local/bin/serve","-s","build"]