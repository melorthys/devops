FROM node:14.19.0-alpine3.15 AS deps
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install

FROM node:14.19.0-alpine3.15 AS build
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN yarn build
RUN rm -rf node_modules
RUN yarn install --production

FROM node:14.19.0-alpine3.15
ENV NODE_ENV production
RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001
WORKDIR /app
COPY --from=build --chown=nextjs:nodejs /app/package.json /app/yarn.lock ./
COPY --from=build --chown=nextjs:nodejs /app/node_modules ./node_modules
COPY --from=build --chown=nextjs:nodejs /app/public ./public
COPY --from=build --chown=nextjs:nodejs /app/.next ./.next
COPY --from=build --chown=nextjs:nodejs /app/.env ./.env
COPY --from=build --chown=nextjs:nodejs /app/next.config.js  ./
USER nextjs
EXPOSE 3000
CMD [ "yarn", "start" ]
