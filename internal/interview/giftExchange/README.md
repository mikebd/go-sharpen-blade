In this interview, I was asked to randomly associate gift
givers and receivers.

Everyone must give and receive one gift.

Gifts must be given to someone else.

### Pathological Case

There is a trap in this problem.  A bad approach to this problem:
* Iterate through the list of people one at a time and randomly choose a receiver for each sender
* Don't allow a person to be a receiver of their own gift
* Track whether a each receiver has received a gift
* Example problem: 3 people, depending on the random results may result in:
  * person 1 gives to person 2
  * person 2 gives to person 1
  * person 3 has no one to give to