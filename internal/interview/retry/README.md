In this interview, I was asked to write retryable operations.
Initially the operations accepted no parameters and returned only an
error if they failed, or nil if they succeeded.

As an extension, I was asked to have the operations accept a parameter.
I proposed that the parameter should determine which attempt number
should succeed, any prior attempts should fail.

Prior to the interview, I was provided the following orientation:
You will decide which language to use, and how to run and test your code.
The prompt will describe a function for you to implement. It will be a 
[higher-order function](https://en.wikipedia.org/wiki/Higher-order_function),
so may accept a function as an argument, return a new function as its result,
or both.
