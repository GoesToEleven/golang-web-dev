Don't run this code

Just making updates - a several step process.

IMPORTANT:
Make sure you update your import statements to import packages from the correct location!

In this step:

MongoDB represents JSON documents in binary-encoded format called BSON behind the scenes. BSON extends the JSON model to provide additional data types and to be efficient for encoding and decoding within different languages.

We will update our user model to change the type of our Id field to be a bson.ObjectId

Add this to models/user.go

```
type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	Id     bson.ObjectId `json:"id" bson:"_id"`
}

```