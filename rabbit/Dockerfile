FROM rabbitmq:3.10.6-management

copy ./definitions.json ./rabbitmq.conf /etc/rabbitmq/

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["rabbitmq-server"]
