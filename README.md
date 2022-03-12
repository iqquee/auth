# auth
Easy way to get an authetication system up and running in your application with Go

### Download package with...
    go get github.com/hisyntax/auth
### Create a .env file in your project root directory then copy the key value pairs below and assign appropraite values
    DATABASE_NAME=<database name>
    USER_COL=<collection name>
    PORT=<port number>
    SECRET_KEY=<secret key>
    MongoDB_URI=<mondodb uri>

### Note that I am using the Gin framework

### Declear the user signup route
### The user model for signup
    type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	User_id       string             `json:"user_id"`
	First_Name    string             `json:"first_name"`
	Last_Name     string             `json:"last_name"`
	Email         string             `json:"email"`
	Phone_Number  int                `json:"phone_number"`
	Password      string             `json:"password"`
	Token         string             `json:"token"`
	Refresh_Token string             `json:"refresh_token"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
}
#### Import `"github.com/hisyntax/auth/auth"`
    Example code 
    {
        r.POST("signup", auth.SignUp)
    }

    Using the auth package, access the SignUp method with `auth.SignUp`


### Declear the user signin route
### The user model for signup
    type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
#### Import `"github.com/hisyntax/auth/auth"`
    Example code 
    {
        r.POST("signin", auth.SignIn)
    }

    Using the auth package, access the SignIn method with `auth.SignIn`

### You can also get information of a specific user

### This is user Model which is returned
    type PublicUser struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name   string             `json:"first_name"`
	Last_Name    string             `json:"last_name"`
	Email        string             `json:"email"`
	Phone_Number string             `json:"phone_number"`
}
###  You have to pass the user email from the header 
    email is the Key
    user@something.com is the value
#### Import "github.com/hisyntax/auth/user"
    Example code 
    {
        r.GET("user", user.GetPublicUser)
    }

    Using the auth package, access the GetPublicUser method with `user.GetPublicUser`

### You can also get information of a all users
### This is user Model which is returned
    type PublicUser struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name   string             `json:"first_name"`
	Last_Name    string             `json:"last_name"`
	Email        string             `json:"email"`
	Phone_Number string             `json:"phone_number"`
}
#### Import "github.com/hisyntax/auth/user"
    Example code 
    {
        r.GET("users", user.GetPublicUsers)
    }

    Using the auth package, access the GetPublicUsers method with `user.GetPublicUsers`


### If you need a users authorization for certain parts of the application like create post and so on, this package also has another method called Authorization
###  You have to pass the user token generated for the user from the header 
    token is the Key 
    generatedUserToken is the Value
    
#### Import `"github.com/hisyntax/auth/middleware"`
    Example code 
    {
        r.POST("").Use(middleware.Authentication)
    }

    Using the auth package, access the Authentication method with `middleware.Authentication`

### Run you application and you have your endpoints running and working
### Currently this package uses a predefined User Model which you can overwrite to fit into you needs but I would be working on that so it does not use a predefined user model anymore (It would instead use the user Model you defined). Contributions to this package are welcomed as it is open source.