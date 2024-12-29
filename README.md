# Novel Nudge Recommendation Service

Implementation of Novel Nudge's recommendation service for recommending books. This project leverages Collaborative and content filtering to find books a user will like

## Features

- Content Filtering

  - Generates a user embedding representing a users book interests
  - Uses cosine distance for closeness
  - Performs a similarity search on book embeddings that match

- Collaborative Filtering

  - Looks for similar users using cosine distance and vectors
  - Looks at users with a small distance and recommends based on shared books

- Re-ranking
  - Uses both forms of filtering to generate final list
  - Generate a final score using both forms of filtering

## Tests

**TODO ¯\\\_(ツ)\_/¯**

## Deployment

This service along with Novel Nudge embeddings are deployed in Cloud Functions _(gcp)_
