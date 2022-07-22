## How to run
* ### Clone
      git clone git@github.com:asssswv/music-shop-v2.git
* ### Build
      sudo docker-compose up --build music-shop   
* ### But we get this error
      Attaching to music-shop-v2-music-shop-1
      music-shop-v2-music-shop-1  | {"level":"fatal","msg":"error init tables: pq: password authentication failed for user \"postgres\"","time":"2022-07-22T05:49:40Z"}
      music-shop-v2-music-shop-1 exited with code 1
* ###  _After build, our container in which postgres is launched is running, and the second one is not. We need to go to the container in which the postgres is spinning and manually set the postgres user to change the password to pass_
    * ### Get container id
          docker ps
    * ### Go to container
          docker exec -it <id> /bin/bash 
    * ### Run psql
          psql -U postgres
    * ### Change password of user postgres
          \password
          Enter new password for user "postgres": pass
          Enter it again: pass
* ### And finally run the application
      sudo docker-compose up music-shop