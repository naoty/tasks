FROM node
WORKDIR /tasks
ADD package.json yarn.lock /tasks/
RUN yarn install
ADD . /tasks/
CMD ["npx", "nuxt", "start"]
