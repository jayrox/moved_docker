# moved_docker
move files when downloads have been completed

# install

`git clone github.com/jayrox/moved_docker`

`cd moved_docker`

`docker build --no-cache -t jayrox/moved .`

# run

`docker run --rm -v /y/Media/SyncTest/Movies:/mnt/src -v /y/Media/PreQueue/Movies:/mnt/dst jayrox/moved`
