# API Requirements

## Endpoints

### api/posts

- GET /posts returns all posts
- GET /posts/[postID] returns a specific post
- GET /posts/announcement returns the current announcement
- GET /posts/sticky returns all stickied posts

- POST /posts/[postID] makes a new post

- UPDATE /posts/[postID] updates a post

- DELETE /posts/[postID] deletes a post

### api/tokensignin

This uses Google's API for sign in auth.

- POST /tokensignin attempts to sign in with an ID token

## Business Rules

- Any user can create a new post
    - These posts must be verified by an admin
- Any user can comment on a post
- Only admins can delete posts
- Only admins can post official updates
- Only admins can sticky a post
- Only admins can update the tags of a post
- All posts can be retrieved when logged out
