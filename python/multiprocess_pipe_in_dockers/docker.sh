
docker run -it --rm --name client -v `pwd`/client.py:/tmp/client.py --volumes-from server:ro -w /tmp/server python:3.6 bash -c 'python3.6 /tmp/client.py'

docker run -it --rm --name server -v `pwd`/server.py:/tmp/server.py -v /tmp/server           -w /tmp/server python:3.6 bash -c 'python3.6 /tmp/server.py'

