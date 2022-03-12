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
#### Import `"github.com/hisyntax/auth/auth"`
    Example code 
    {
        r.POST("signup", auth.SignUp)
    }

    Using the auth package, access the SignUp method with `auth.SignUp`


### Declear the user signin route
#### Import `"github.com/hisyntax/auth/auth"`
    Example code 
    {
        r.POST("signin", auth.SignIn)
    }

    Using the auth package, access the SignIn method with `auth.SignIn`

### You can also get information of a specific user
#### Import "github.com/hisyntax/auth/user"
    Example code 
    {
        r.GET("user", user.GetPublicUser)
    }

    Using the auth package, access the GetPublicUser method with `user.GetPublicUser`

### You can also get information of a all users
#### Import "github.com/hisyntax/auth/user"
    Example code 
    {
        r.GET("users", user.GetPublicUsers)
    }

    Using the auth package, access the GetPublicUsers method with `user.GetPublicUsers`


### If you need a users authorization for certain parts of the application like create post and so on, this package also has another method called Authorization
#### Import `"github.com/hisyntax/auth/middleware"`
    Example code 
    {
        r.POST("").Use(middleware.Authentication)
    }

    Using the auth package, access the Authentication method with `middleware.Authentication`

### Run you application and you have your endpoints running and working
### Currently this package uses a predefined User Model which you can edit to fit into you needs but I would be working on that so it does not use a predefined user model anymore (It would instead use the user Model you defined). Contributions to this package are welcomed as it is open source.