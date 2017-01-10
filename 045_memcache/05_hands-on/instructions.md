# Import These Packages
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"

# Instructions
1. Store a uuid in a cookie
  1. Value:uuid
1. Store a value in memcache
  1. key:uuid
  1. value:your name 
1. Retrieve the uuid from the cookie
1. Retrieve the uuid & value from memcache