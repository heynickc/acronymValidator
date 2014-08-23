[![Build Status](https://travis-ci.org/heynickc/acronymValidator.svg?branch=master)](https://travis-ci.org/heynickc/acronymValidator)

Pseudocode:

Set up a container to hold matched acronym characters
  Set the length of the container to the number of product names
  Set the maximum capacity of each bucket in the container
For the length of the acronym
  Add one to each bucket
    At the end of the container
    If there are more letters
      Start back at the beginning of the container
    Else break

The container's bucket capacities should resize as each one fills and needs more capacity
  Methods:
    Add item to bucket
      If the bucket has enough capacity
        Add item to bucket
        Reduce capacity by 1
      If bucket doesn't have enough capacity
        Try to resize to fit
        Else return error

    Resize to fit
      For each bucket starting at last going backwards until current bucket
        Check capacity at index
          If the last bucket's capacity > 1
            Reduce the last bucket's capacity by 1
            Add capacity of 1 to current bucket
          If the last bucket's capacity = 1
        Check the capacity of the bucket index - 1

  Method for setting upcapacities ([]bucket, int)
    Add one to each bucket capacity
    At the end of the bucket array
    If int > 0
      Start back at the beginning of the list
    Else
      return list

For each name in product names
  Add name to bucket
  For each letter in acronym
    Pop first letter from acronym
    If it's contained in the name
      If the name bucket has available capacity
        Add acronym letter to name bucket
      Else break and go to the next name
  If bucket is empty
    Return false
  Else break and go to the next name

