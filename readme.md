# Social Network Prototype

## Description

This project implements a simple social network. Features include:

## Create User Profile

- **Method:** POST
- **URL:** `localhost:8080/api/profile`
- **Description:** Allows users to create and manage their profiles. Users can add details like their name, bio, and profile picture.

### Example request
```bash
curl --request POST \
  --url http://localhost:8080/api/profile \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Artem",
"avatar":"https://letsenhance.io/static/8f5e523ee6b2479e26ecc91b9c25261e/1015f/MainAfter.jpg",
	"email": "artem@gmail.com",
	"password": "pass",
	"confirmPassword": "pass"
}'
```

### Example response
```json
{
  "ID": "669e77d236bb9d246788d489",
  "Email": "c@gmail.com",
  "Name": "Artem",
  "Avatar": "https://letsenhance.io/static/8f5e523ee6b2479e26ecc91b9c25261e/1015f/MainAfter.jpg",
  "Password": "$2a$10$d15J/63//vQHBR9zAMC.ae4qlD5SPfjtJLf3NqxPlzQcBMJe0ZjYm",
  "Posts": {},
  "LikedPosts": {},
  "Notifications": []
}
```
*Note*: password was left on purpose. Just to let you know how it is stored in the database.


## Retrieve User Profile

- **Method:** GET
- **URL:** `localhost:8080/api/profile`
- **Description:** Provides the functionality to view user profiles. You can access details of a user's profile by their unique identifier.

### Example request
```bash
curl --request GET \
  --url http://localhost:8080/api/profile \
  --user artem@gmail.com:pass
```

### Example response
```json
{
  "ID": "669e77d236bb9d246788d489",
  "Email": "c@gmail.com",
  "Name": "Artem",
  "Avatar": "https://letsenhance.io/static/8f5e523ee6b2479e26ecc91b9c25261e/1015f/MainAfter.jpg",
  "Password": "$2a$10$d15J/63//vQHBR9zAMC.ae4qlD5SPfjtJLf3NqxPlzQcBMJe0ZjYm",
  "Posts": {},
  "LikedPosts": {},
  "Notifications": []
}
```


## Create Post

- **Method:** POST
- **URL:** `localhost:8080/api/posts`
- **Description:** Enables users to create posts. Posts can include text, images, or other media.

### Example request
```bash
curl --request POST \
  --url http://localhost:8080/api/posts \
  --user artem@gmail.com:pass \
  --header 'Content-Type: application/json' \
  --data '{
	"content": "https://img.freepik.com/free-photo/painting-mountain-lake-with-mountain-background_188544-9126.jpg"
}'
```

## Retrieve Posts

- **Method:** GET
- **URL:** `localhost:8080/api/posts`
- **Description:** Allows users to view posts. Users can fetch posts by their unique identifier or view a list of posts.

### Example request
```bash
curl --request GET \
  --url http://localhost:8080/api/posts \
  --user artem@gmail.com:pass
```

### Example response
```json
[
  {
    "ID": "669e77f836bb9d246788d48a",
    "Content": "https://img.freepik.com/free-photo/painting-mountain-lake-with-mountain-background_188544-9126.jpg",
    "Author": "669e77d236bb9d246788d489",
    "LikesCount": 0
  }
]
```

## Like Post

- **Method:** POST
- **URL:** `localhost:8080/api/posts/:id/like`
- **Description:** Provides the functionality for users to like posts. This action increases the like count of a post.

### Example request
```bash
curl --request POST \
  --url http://localhost:8080/api/posts/669e77f836bb9d246788d48a/like \
  --user artem@gmail.com:pass
```

*Note*: to unlike post just send the same request again.

## See Liked Posts

- **Method:** GET
- **URL:** `localhost:8080/api/posts/liked`
- **Description:** Allows users to view posts they have liked.

### Example request
```bash
curl --request GET \
  --url http://localhost:8080/api/posts/liked \
  --user artem@gmail.com:pass
```

### Example response
```json
[
  {
    "ID": "669e77f836bb9d246788d48a",
    "Content": "https://img.freepik.com/free-photo/painting-mountain-lake-with-mountain-background_188544-9126.jpg",
    "Author": "669e77d236bb9d246788d489",
    "LikesCount": 1
  }
]
```

## Notifications

- **Method:** GET
- **URL:** `localhost:8080/api/posts/liked`
- **Description:** Users receive notifications when their posts are liked. Notifications help users stay updated about interactions with their content.

### Example request
```bash
curl --request GET \
  --url http://localhost:8080/api/notifications \
  --user artem@gmail.com:pass
```

### Example response
```json
[
	{
		"Id": "669e780636bb9d246788d48b",
		"Type": "like",
		"PostId": "669e77f836bb9d246788d48a",
		"LikedBy": "669e5d0af5e28ac22ab3d811"
	}
]
```