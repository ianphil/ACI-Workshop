from python:3.7.2-stretch

RUN mkdir scripts
WORKDIR /scripts

COPY . .

RUN pip install -r requirements.txt

CMD [ "python", "receive.py" ]