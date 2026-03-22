FROM ubuntu:latest

COPY *.deb /app/

RUN apt update && \
    apt install -y dpkg && \
    dpkg -i /app/*.deb || apt install -f -y

CMD ["latin_checker", "3", "1 2 3 2 3 1 3 1 2"]