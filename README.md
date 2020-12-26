# image-build

root@student0:~/a# docker image build -f Dockerfile -t multi-stage-build .

root@student0:~/a# docker image build -f Dockerfile2 -t onlyrun-build .


root@student0:~/a# docker image ls multi-stage-build
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
multi-stage-build   latest              7aa5932e6908        25 minutes ago      12MB
root@student0:~/a# docker image ls onlyrun-build
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
onlyrun-build       latest              4dac2f68bdc5        33 seconds ago      299MB
