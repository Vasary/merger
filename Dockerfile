FROM apline:latest

RUN mkdir /app

COPY ./dist/app /app/merger
COPY .env /app/.env

RUN chmod +x /app/app
RUN export $(cat /app/.env | xargs)

WORKDIR /app

CMD /app/merger