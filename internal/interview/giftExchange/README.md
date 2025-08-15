In this interview, I was asked to randomly associate gift
givers and receivers.

Everyone must give and receive one gift.

Gifts must be given to someone else.

### Pathological Case

There is a trap in this problem.

Here is a bad approach:
* Iterate through the list of people one at a time and randomly choose a receiver for each sender
* For each iteration, continue choosing a random receiver if either of the following is true:
  * The receiver is the same as the sender
  * The receiver has already received a gift
* Example problem with 3 people, depending on the random results for sender 1 & 2, may result in:
  * person 1 gives to person 2
  * person 2 gives to person 1
  * person 3 has no one to give to