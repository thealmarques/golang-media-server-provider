# Go-Media Server

Go-Media Server is a storage service for media files (although it can be used for all files, at this point). Right now it only includes static files, since this project was built to help me in another project ([Netflix-React](https://github.com/thealmarques/netflix-react)).

# Set up

Store your files in
```
resources/media
```

Run the application using

```javascript
docker-compose up --build -d
```

# How to use

Access your files using:
```
http://locahost:9081/media/media_file
```

# To-do

[ ] Media files only

[ ] Add files during runtime

# Contribution

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
