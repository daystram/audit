FROM node:lts-alpine as builder
WORKDIR /build
COPY audit-fe audit-fe
WORKDIR /build/audit-fe
RUN yarn install --frozen-lockfile
RUN yarn build --mode production
RUN cat .env.production

FROM nginx:stable-alpine as production
COPY --from=builder /build/audit-fe/dist /usr/share/nginx/html
COPY --from=builder /build/audit-fe/nginx.conf /etc/nginx/conf.d
RUN rm /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
