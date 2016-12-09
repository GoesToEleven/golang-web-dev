# Install Postgresql

1. [Postgresql Linux download page](https://www.postgresql.org/download/linux/ubuntu/)
  - follow the directions on this page

## When creating this training, the instructions were like this for me

1. Discover linux release & codename
  - Log into your linux box.
  - terminal: lsb_release -a
  
1. Create this file
  - terminal: nano /etc/apt/sources.list.d/pgdg.list
  - add this content to the file (though you might have a different release codename) :
  ```
  deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main
  ```
  - terminal: wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
  - terminal: sudo apt-get update
  
## Now install postgres

1. terminal: apt-get install postgresql-9.4
