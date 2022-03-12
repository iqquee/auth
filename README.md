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

### Import `"github.com/hisyntax/auth/auth"`
### Declear the user signup route
    Example code 
    {
        r.POST("signup", auth.SignUp)
    }

    Using the auth package, access the SignUp method with `auth.SignUp`


### Declear the user signin route
    Example code 
    {
        r.POST("signin", auth.SignIn)
    }

    Using the auth package, access the SignIn method with `auth.SignIn`

### If you need a users authorization for certain parts of the application like create post and so on, this package also has another method called Authorization
#### Import `"github.com/hisyntax/auth/middleware"`
    Example code 
    {
        r.POST("").Use(middleware.Authentication)
    }

    Using the auth package, access the Authentication method with `middleware.Authentication`

### Run you application and you have your endpoints running and working
