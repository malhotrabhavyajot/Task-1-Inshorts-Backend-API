# Task-1-Inshorts-Backend-API

The task is to develop a basic version of a news app like InShorts. You are only required to develop the API for the system. 
Below are the details.

You are required to Design and Develop an HTTP JSON API capable of the following operations,
# Create an article
### Should be a POST request.
### Use JSON request body.
### URL should be ‘/articles’.

# Get an article using id
Should be a GET request.
Id should be in the url parameter.
URL should be ‘/articles/<id here>’.

# List all articles
Should be a GET request.
URL should be ‘/articles’.

# Search for an Article (search in title, subtitle, content)
Should be a GET request.
Search term should be in the query parameter with key ‘q’.
URL should be ‘/articles/search?q=<search term here>’.
