# Bit Blossom
> Learning and Implementing Bloom Filters

- When we are comparing two strings, it's an O(m + n) to compare them. There are some good data structures around strings such as Tries, Suffix Trees, etc.
    - It's not easy to optimize Tries for any problem, they have good theoretical complexity, in practical applications, it may not be the best.
- In practice, hashing is the best approach we can have. We can simply use a Hash Map, why not?
    - What could be possible is, we don't have the bucket range to store all the strings, the range of the hash could be 0 to some number N. N is not very large, what we can store is multiple hashes.
    - Let's say there is a hash function, the output range of the hash function need not to high, so there are chances for collision. In such scenarios, we can use multiple hash functions to reduce collisions.
        - If we have three hash functions and two strings have the same values for three different hash functions, there is a very low probability that two different strings will give the same output for three different hash functions.
- This is concept of __Bloom Filters__.
    - Facebook also uses this for "one hit wonder".
    - Let's say we search for something which we don't search that often, the cache and CDN stores the commonly searched details, so that it won't have to keep hitting the servers for the response.
        - The problem with this approach is that, when we store something like this in a cache, which we may never search for again, it's going to space for something which might be more often searched.
        - If we use something like LRU, something more useful might be removed from the cache.
        - Because of this, the cache is poor, it would have to hit the server more, leading slower browser experience.
    - What Facebook does is, they won't store it the cache on justing hitting the entity once, they would expect you to search for it multiple times before storing it into the cache.
