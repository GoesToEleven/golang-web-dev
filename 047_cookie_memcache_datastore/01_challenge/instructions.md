# Instructions
1. Store a uuid in a cookie
  1. Value:uuid
1. Store a value in memcache
  1. key:uuid
  1. value:your name 
1. Store a value in datastore
  1. stringID:uuid
  1. value:some struct
1. Retrieve the uuid from the cookie
1. Retrieve the value from memcache
  1. if it's not in memcache
    1. Retrieve the value from memcache
    1. Store value in memcache
