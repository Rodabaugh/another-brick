# About
another-brick is a very quick webapp that I created to learn more about HTMX and templ. It is a basic post wall/bulletin board. There is no authentication, and the wall is shared by all users. It was not designed as a production product, but as a fun toy/experiment. I am publicly hosting the project for a few days, and You can check out the project here: https://wall.erik.coffee/
 
# This Project Involved
- Storing posts in a PostgreSQL database.
- Getting a list of posts from the database and rendering them on the page using templ.
- Getting a post from the user using HTMX, passing it to the backend, and storing it in the database.
- Having HTMX automatically update the list of posts after writing to the database.
- Dynamically passing the ID of a post to the DELETE /api/posts endpoint so the backend can delete the post.
- Updating the posts list after the delete has been performed.

The new post and delete post endpoints check to see if the response should be in JSON and will reply in JSON if it is. This allows the use of a custom client over the API. If the Accept is not set to "application/json", the backend will reply with HTML so HTMX can render it on the page.

# Hosting
I hosted this project on a VPS, and used an nginx reverse proxy for public access. This allows me to use certbot to enable SSL with Let's Encrypt (not that it matters on this project.) nginx also enables me to see how much traffic the page gets using the nginx logs. If you would like to host this yourself, here are some basic steps to get the application up and running.

## Get the repo

1. Clone the repo `git clone https://github.com/Rodabaugh/another-brick/`
2. Navigate to the program dir `cd another-brick`

## Configuration

Environment variables are used for configuration. Create a .env file in the root of the project dir. You need to specify your DB_URL, PLATFORM. DB_URL is the url for your database. Platform can either be "prod" or "dev". Your `.env` file should look something like the one below. Please be sure to create your own JWT_SECRET and use your own DB_URL.
```
PLATFORM=prod
DB_URL="postgres://postgresUser:postgresPass@localhost:5432/another_brick?sslmode=disable"
```

A port may also be specified using ```PORT=1234```. If a port is not specified, it will default to 8080. You can also set a SITE_TITLE and SUB_TITLE.

## Setting up the database

Goose is used to manage the database migrations. Install goose with `go install github.com/pressly/goose/v3/cmd/goose@latest`

Navigate to the sql/schema dir `cd sql/schema`

Setup the database using goose `goose postgres <connection_string> up` e.g `goose postgres postgres://postgresUser:postgresPass@localhost:5432/another_brick?sslmode=disable up`

## Compile and run the application

Once your .env has been configured, and your database is setup, it is time to build and run the application.

Build the application with `make build`

Run the backend application with `./another-brick`

Once the application is running, you can setup your server to run the application as a service.

# API Endpoints
Be sure to set your Accept header to "application/json", so the backend responds with JSON and not HTML.

## POST /api/posts
Request body:
```json
{
    "content": "Post text"
}
```

Response body:
```json
{
	"id": "83e2613c-e923-4ad0-9a27-c4bd528dc849",
	"created_at": "2025-05-25T18:17:52.092405Z",
	"updated_at": "2025-05-25T18:17:52.092405Z",
	"content": "Post text"
}
```

## DELETE /api/posts/{post_id}

No content is required for this request. An empty post object will be returned, along with a 200.

Response body:
```json
{
	"id": "00000000-0000-0000-0000-000000000000",
	"created_at": "0001-01-01T00:00:00Z",
	"updated_at": "0001-01-01T00:00:00Z",
	"content": ""
}
```