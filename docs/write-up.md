# Notes 

## Initial Thoughts

- Two platforms I think personally has a very good recommendation are the 
  giants of Spotify and Netflix, when building my own recommendation system 
  we have to start there
- In a similar way to dating websites match people together? How do 
  they define the expectation that someone will match someone else?
- How do we assign numerical quantitative to something as qualitative as 
  their personality
- How do we express the Identity the essence of a person in numerical values?

#### The Cold Start problem
- How do we generate recommendations without any user data fresh?
  - We could ask a user to select some books they have ready before
  - We can ask a user to select Genres they enjoy
  - Do we match to defaults for their demographics such as age/location?

#### Considerations
- How do we converge the algorithm as soon as possible?
  - Do we recommend outliers early at the expense of providing bad 
    recommendations?
  - Do we over-correct ratings early on? How do we ensure they don't get 
    stuck in a state which they don't like?
  
## Algorithm

- On dating apps both users must like each-other for the match to be 
  considered successful. I want to implement this into my book 
  recommendation service. For a person to match to book both the book and 
  the person must match. 

#### Features

The features we use to recommend must be the same for a user as it is for 
the book to allow for any meaningful comparison.
 

## References 
- https://help.netflix.com/en/node/100639
- https://www.music-tomorrow.com/blog/how-spotify-recommendation-system-works-a-complete-guide-2022